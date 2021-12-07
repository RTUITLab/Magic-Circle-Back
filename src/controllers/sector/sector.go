package sector

import (
	"context"
	"errors"
	"net/http"

	"github.com/0B1t322/Magic-Circle/ent"
	"github.com/0B1t322/Magic-Circle/ent/adjacenttable"
	"github.com/0B1t322/Magic-Circle/ent/direction"
	"github.com/0B1t322/Magic-Circle/ent/institute"
	"github.com/0B1t322/Magic-Circle/ent/predicate"
	"github.com/0B1t322/Magic-Circle/ent/profile"
	"github.com/0B1t322/Magic-Circle/ent/sector"
	"github.com/0B1t322/Magic-Circle/ent/variant"
	. "github.com/0B1t322/Magic-Circle/models/sector"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var (
	ErrSectorExist    = errors.New("Sector with this coords exist")
	ErrSectorNotFound = errors.New("Sector not found")
)

type SectorController struct {
	Client *ent.Client
}

func New(client *ent.Client) *SectorController {
	return &SectorController{
		Client: client,
	}
}

func (s SectorController) create(ctx context.Context, coords, description string) (*ent.Sector, error) {
	created, err := s.Client.Sector.Create().
		SetCoords(coords).
		SetDescription(description).
		Save(ctx)

	if ent.IsConstraintError(err) { // Sector with this coords exist
		return nil, ErrSectorExist
	} else if err != nil {
		return nil, err
	}

	return created, nil
}

type CreateSectorReq struct {
	Coords      string `json:"coords"`
	Description string `json:"description"`
}

func newLogFields(method string, err error) log.Fields {
	return log.Fields{
		"controller": "sector",
		"method":     method,
		"err":        err,
	}
}

// CreateSector
//
// @Summary Create Sector
//
// @Description create sector according to giving coords
//
// @Description coords is unique string
//
// @Router /v1/sector [post]
//
// @Accept json
//
// @Produce json
//
// @Param body body sector.CreateSectorReq true "body"
//
// @Success 201 {object} sector.Sector
//
// @Failure 400 {string} srting
func (s SectorController) Create(c *gin.Context) {
	var req CreateSectorReq
	{
		if err := c.ShouldBindJSON(&req); err != nil {
			c.String(http.StatusBadRequest, "Unexpected body")
			c.Abort()
			return
		}
	}

	created, err := s.create(c, req.Coords, req.Description)
	if err == ErrSectorExist {
		c.String(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	} else if err != nil {
		log.WithFields(newLogFields("Create", err)).Error("Failed to create sector")
		c.String(http.StatusInternalServerError, "Failed to create sector")
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, NewSector(created))
}

type UpdateSectorReq struct {
	ID          int     `json:"-" uri:"id"`
	Coords      *string `json:"coords,omitempty"`
	Description *string `json:"description,omitempty"`
}

func (s SectorController) update(
	ctx context.Context, 
	req UpdateSectorReq,
) (*ent.Sector, error) {
	builder := s.Client.Sector.UpdateOneID(req.ID)

	if req.Coords != nil {
		builder.SetCoords(*req.Coords)
	}

	if req.Description != nil {
		builder.SetDescription(*req.Description)
	}

	updated, err := builder.Save(ctx)
	if ent.IsConstraintError(err) {
		return nil, ErrSectorExist
	} else if ent.IsNotFound(err) {
		return nil, ErrSectorNotFound
	} else if err != nil {
		return nil, err
	}

	return updated, nil
}

// UpdateSector
//
// @Summary Update Sector
//
// @Description update sector
//
// @Router /v1/sector/{id} [put]
//
// @Param id path string true "id of sector"
// 
// @Accept json
//
// @Produce json
//
// @Param body body sector.UpdateSectorReq true "body"
//
// @Success 200 {object} sector.Sector
//
// @Failure 400 {string} srting
func (s SectorController) Update(c *gin.Context) {
	var req UpdateSectorReq
	{
		if err := c.ShouldBindJSON(&req); err != nil {
			c.String(http.StatusBadRequest, "Unexpected body")
			c.Abort()
			return
		}

		if err := c.ShouldBindUri(&req); err != nil {
			c.String(http.StatusBadRequest, "Unexpected body")
			c.Abort()
			return
		}
	}

	updated, err := s.update(c, req)
	if err == ErrSectorExist {
		c.String(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	} else if err == ErrSectorNotFound {
		c.String(http.StatusNotFound, err.Error())
		c.Abort()
		return
	} else if err != nil {
		log.WithFields(newLogFields("Update", err)).Error("Failed to update sector")
		c.String(http.StatusInternalServerError, "Failed to update")
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, NewSector(updated))
}

func (s SectorController) getAll(
	ctx 	context.Context,
	req		GetAllSectorsReq,
	) ([]*ent.Sector, error) {
	builder := s.Client.Sector.Query()
	
	var preds []predicate.Variant
	{
		if req.DirectionName != "" {
			preds = append(preds, variant.HasDirectionWith(direction.Name(req.DirectionName)))
		}

		if req.InstituteName != "" {
			preds = append(preds, variant.HasInsituteWith(institute.Name(req.InstituteName)))
		}

		if req.ProfileName != "" {
			preds = append(preds, variant.HasProfileWith(profile.Name(req.ProfileName)))
		}
	}

	if len(preds) > 0 {
		builder.Where(
			sector.HasAdjacentTablesWith(
				adjacenttable.HasVariantWith(
					preds...
				),
			),
		)
	}

	return builder.All(ctx)
}

type GetAllSectorsReq struct {
	InstituteName		string		`json:"-" query:"institute"`
	DirectionName		string		`json:"-" query:"direction"`
	ProfileName			string		`json:"-" query:"profile"`
}

type GetAllSectorsResp struct {
	Sectors []Sector		`json:"sectors"`
}


// GetSectors
//
// @Summary Get Sectors
//
// @Description return all sectors
//
// @Router /v1/sector [get]
// 
// @Param institute query string false "institute name"
// 
// @Param direction query string false "direction name"
// 
// @Param profile query string false "profile name"
//
// @Produce json
//
// @Success 200 {object} sector.GetAllSectorsResp
//
// @Failure 500 {string} srting
func (s SectorController) GetAll(c *gin.Context) {
	var req GetAllSectorsReq
	{
		req.InstituteName = c.Query("institute")
		req.DirectionName = c.Query("direction")
		req.ProfileName = c.Query("profile")
	}

	log.Infof("%+v", req)

	get, err := s.getAll(c, req)
	if err != nil {
		log.WithFields(newLogFields("Update", err)).Error("Failed to get sectors")
		c.String(http.StatusInternalServerError, "Failed to get sectors")
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, GetAllSectorsResp{Sectors: NewSectors(get)})
}