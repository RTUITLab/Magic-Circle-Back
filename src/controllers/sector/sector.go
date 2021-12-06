package sector

import (
	"context"
	"errors"
	"net/http"

	"github.com/0B1t322/Magic-Circle/ent"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var (
	ErrSectorExist = errors.New("Sector with this coords exist")
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

func (s SectorController) create(ctx context.Context, coords string) (*ent.Sector, error) {
	created, err := s.Client.Sector.Create().
		SetCoords(coords).
		Save(ctx)
	
	if ent.IsConstraintError(err) { // Sector with this coords exist
		return nil, ErrSectorExist
	} else if err != nil {
		return nil, err
	}

	return created, nil
}

type CreateSectorReq struct {
	Coords	string	`json:"coords"`
}

type Sector struct {
	ID		int		`json:"int"`
	Coords	string	`json:"coords"`
}

func NewSector(s *ent.Sector) Sector {
	return Sector{
		ID: s.ID,
		Coords: s.Coords,
	}
}

func newLogFields(method string, err error) log.Fields {
	return log.Fields{
		"controller": "sector",
		"method": method,
		"err": err,
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

	created, err := s.create(c, req.Coords)
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