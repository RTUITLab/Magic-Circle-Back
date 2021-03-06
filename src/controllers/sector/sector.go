package sector

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/0B1t322/Magic-Circle/ent"
	"github.com/0B1t322/Magic-Circle/ent/adjacenttable"
	"github.com/0B1t322/Magic-Circle/ent/direction"
	"github.com/0B1t322/Magic-Circle/ent/institute"
	"github.com/0B1t322/Magic-Circle/ent/predicate"
	"github.com/0B1t322/Magic-Circle/ent/profile"
	"github.com/0B1t322/Magic-Circle/ent/sector"
	"github.com/0B1t322/Magic-Circle/models/role"
	. "github.com/0B1t322/Magic-Circle/models/sector"
	"github.com/0B1t322/Magic-Circle/pkg/queue"
	jwt "github.com/appleboy/gin-jwt/v2"
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

func (s SectorController) create(
	ctx context.Context, coords,
	description string,
	// additionalDescription string,
) (*ent.Sector, error) {
	created, err := s.Client.Sector.Create().
		SetCoords(coords).
		SetDescription(description).
		// SetAdditionalDescription(additionalDescription).
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
	// AdditionalDescription string `json:"additionalDescription"`
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
// @Security ApiKeyAuth
//
// @Tags sector
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
	// AdditionalDescription *string `json:"additionalDescription,omitempty"`
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

	// if req.AdditionalDescription != nil {
	// 	builder.SetAdditionalDescription(*req.AdditionalDescription)
	// }

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
// @Security ApiKeyAuth
//
// @Tags sector
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

func buildInstPredicate(instId int) predicate.AdjacentTable {
	return adjacenttable.HasProfileWith(
		profile.HasDirectionWith(
			direction.HasInstituteWith(
				institute.ID(instId),
			),
		),
	)
}

func buildDirPredicate(dirId int) predicate.AdjacentTable {
	return adjacenttable.HasProfileWith(
		profile.HasDirectionWith(
			direction.ID(dirId),
		),
	)
}

func buildProfPredicate(profId int) predicate.AdjacentTable {
	return adjacenttable.HasProfileWith(
		profile.ID(profId),
	)
}

func buildPredicate(inst, prof, dir int) (pred predicate.AdjacentTable) {
	var preds []predicate.AdjacentTable
	{
		if inst != -1 {
			preds = append(preds, buildInstPredicate(inst))
		}

		if prof != -1 {
			preds = append(preds, buildProfPredicate(prof))
		}

		if dir != -1 {
			preds = append(preds, buildDirPredicate(dir))
		}
	}

	return adjacenttable.And(
		preds...,
	)
}

func (s SectorController) getAll(
	ctx context.Context,
	req GetAllSectorsReq,
) ([]*ent.Sector, error) {
	builder := s.Client.Sector.Query()

	var preds []predicate.AdjacentTable
	{
		castFunc := queue.StringQueueToIntQueue(queue.StringQueueToIntOpts{IfNotIntElemIgnore: true})
		var (
			insts, _ = castFunc(queue.NewStringQueue(strings.Split(req.InstituteName, " ")))
			profs, _ = castFunc(queue.NewStringQueue(strings.Split(req.ProfileName, " ")))
			dirs, _  = castFunc(queue.NewStringQueue(strings.Split(req.DirectionName, " ")))
		)

		for {
			inst, errInst := insts.Get()
			prof, errProf := profs.Get()
			dirs, errDirs := dirs.Get()

			if errInst == queue.QueueIsEmpty &&
				errProf == queue.QueueIsEmpty &&
				errDirs == queue.QueueIsEmpty {
				break
			}

			preds = append(
				preds,
				buildPredicate(inst, prof, dirs),
			)
		}

	}

	if len(preds) > 0 {
		builder.Where(
			sector.HasAdjacentTablesWith(
				adjacenttable.Or(
					preds...,
				),
			),
		)
	}

	return builder.
		WithAdjacentTables(
			func(atq *ent.AdjacentTableQuery) {
				atq.WithProfile(
					func(pq *ent.ProfileQuery) {
						pq.WithDirection(
							func(dq *ent.DirectionQuery) {
								dq.WithInstitute()
							},
						)
					},
				)
			},
		).
		All(ctx)
}

type GetAllSectorsReq struct {
	InstituteName string `json:"-" query:"institute"`
	DirectionName string `json:"-" query:"direction"`
	ProfileName   string `json:"-" query:"profile"`
}

type GetAllSectorsResp struct {
	Sectors []Sector `json:"sectors"`
}

// GetSectors
//
// @Summary Get Sectors
//
// @Description return all sectors
//
// @Description quey params can make a logical predicates for example
// @Description request: "/sectors?instutute=1+2&profile=1" equal "WHERE (institute_id=1 and profile_id=1) or institute_id=2"
//
// @Tags sector
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

	get, err := s.getAll(c, req)
	if err != nil {
		log.WithFields(newLogFields("GetAll", err)).Error("Failed to get sectors")
		c.String(http.StatusInternalServerError, "Failed to get sectors")
		c.Abort()
		return
	}
	
	c.JSON(http.StatusOK, GetAllSectorsResp{Sectors: NewSectors(get)})
}

type GetAllSectorsIdsResp struct {
	Sectors []CompactSector `json:"sectors"`
}

// GetAllIds
//
// @Summary Get Sectors ids and coords
//
// @Description return all sectors
//
// @Description quey params can make a logical predicates for example
// @Description request: "/sectors?instutute=1+2&profile=1" equal "WHERE (institute_id=1 and profile_id=1) or institute_id=2"
//
// @Tags sector
//
// @Router /v1/sectorIds [get]
//
// @Param institute query string false "institute name"
//
// @Param direction query string false "direction name"
//
// @Param profile query string false "profile name"
//
// @Produce json
//
// @Success 200 {object} sector.GetAllSectorsIdsResp
//
// @Failure 500 {string} srting
func (s SectorController) GetAllIds(c *gin.Context) {
	var req GetAllSectorsReq
	{
		req.InstituteName = c.Query("institute")
		req.DirectionName = c.Query("direction")
		req.ProfileName = c.Query("profile")
	}

	get, err := s.getAll(c, req)
	if err != nil {
		log.WithFields(newLogFields("GetAll", err)).Error("Failed to get sectors")
		c.String(http.StatusInternalServerError, "Failed to get sectors")
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, GetAllSectorsIdsResp{Sectors: NewCompactsSector(get)})
}

type GetSectorReq struct {
	ID int `json:"-" uri:"id"`
}

// GetSector
//
// @Summary get sector
//
// @Description return sector
//
// @Tags sector
//
// @Router /v1/sector/{id} [get]
//
// @Param id path string false "sector id"
//
// @Produce json
//
// @Success 200 {object} sector.Sector
//
// @Failure 500 {string} srting
func (s SectorController) GetSector(c *gin.Context) {
	var req GetSectorReq
	{
		if err := c.ShouldBindUri(&req); err != nil {
			c.String(http.StatusBadRequest, "Unexpected id")
			c.Abort()
			return
		}
	}

	get, err := s.Client.Sector.Query().
					Where(
						sector.ID(req.ID),
					).WithAdjacentTables(
						func(atq *ent.AdjacentTableQuery) {
							atq.WithProfile(
								func(pq *ent.ProfileQuery) {
									pq.WithDirection(
										func(dq *ent.DirectionQuery) {
											dq.WithInstitute()
										},
									)
								},
							)
						},
					).Only(c)
	if ent.IsNotFound(err) {
		c.String(http.StatusNotFound, "Sector not found")
		c.Abort()
		return
	} else if err != nil {
		log.WithFields(newLogFields("GetSector", err)).Error("Failed to get sector")
		c.String(http.StatusInternalServerError, "Failed to get sector")
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, NewSector(get))
}

type DeleteSectorReq struct {
	ID int `json:"-" uri:"id"`
}

func (s SectorController) deleteSector(ctx context.Context, id int) error {
	// delete all adjecent tables that relate with sector
	_, err := s.Client.AdjacentTable.Delete().
		Where(
			adjacenttable.HasSectorWith(
				sector.ID(id),
			),
		).Exec(ctx)

	if ent.IsNotFound(err) {
		// Pass
	} else if err != nil {
		return err
	}

	// Delete sector
	if err := s.Client.Sector.DeleteOneID(id).Exec(ctx); ent.IsNotFound(err) {
		return ErrSectorNotFound
	} else if err != nil {
		return err
	}

	return nil
}

// DeleteSector
//
// @Summary Delete Sector
//
// @Description delete sector and all adjacenttables that relate with this sector
//
// @Router /v1/sector/{id} [delete]
//
// @Security ApiKeyAuth
//
// @Tags sector
//
// @Param id path string true "id of sector"
//
// @Produce json
//
// @Success 200
//
// @Failure 500 {string} srting
//
// @Failure 404 {string} srting
func (s SectorController) DeleteSector(c *gin.Context) {
	var req DeleteSectorReq
	{
		if err := c.ShouldBindUri(&req); err != nil {
			c.String(http.StatusBadRequest, "Unexpected id")
			c.Abort()
			return
		}
	}

	if err := s.deleteSector(c, req.ID); err == ErrSectorNotFound {
		c.String(http.StatusNotFound, err.Error())
		c.Abort()
		return
	} else if err != nil {
		log.WithFields(newLogFields("Delete", err)).Error("Failed to delete sector")
		c.String(http.StatusInternalServerError, "Failed to delete")
		c.Abort()
		return
	}

	c.Status(http.StatusOK)
}

type CreateSectorsReq []CreateSectorReq

func (c CreateSectorsReq) getUnique(slice []CreateSectorReq) (sliceOfUnique CreateSectorsReq) {
	var twoSlices []CreateSectorReq
	{
		twoSlices = append(twoSlices, slice...)
		twoSlices = append(twoSlices, c...)
	}

	mapByCoords := map[string]CreateSectorReq{}
	{
		for _, req := range twoSlices {
			if _, find := mapByCoords[req.Coords]; !find {
				mapByCoords[req.Coords] = req
			} else if find {
				delete(mapByCoords, req.Coords)
			}
		}
	}

	for _, req := range mapByCoords {
		sliceOfUnique = append(sliceOfUnique, req)
	}

	return sliceOfUnique
}

func (s SectorController) deleteSectorsThatExist(ctx context.Context, reqs []CreateSectorReq) ([]CreateSectorReq, error) {
	coords := func(reqs []CreateSectorReq) (slice []string) {
		for _, req := range reqs {
			slice = append(slice, req.Coords)
		}
		return slice
	}(reqs)

	sectorThatExist, err := s.Client.Sector.Query().Where(
		sector.CoordsIn(coords...),
	).All(ctx)
	if ent.IsNotFound(err) {
		return reqs, nil
	} else if err != nil {
		return nil, err
	}

	findedSectors := func(finded []*ent.Sector) (slice []CreateSectorReq) {
		for _, find := range finded {
			slice = append(slice, CreateSectorReq{Coords: find.Coords, Description: find.Description})
		}
		return slice
	}(sectorThatExist)

	return CreateSectorsReq(reqs).getUnique(findedSectors), nil
}

func (s SectorController) createALot(ctx context.Context, reqs []CreateSectorReq) ([]*ent.Sector, error) {
	var builders []*ent.SectorCreate

	filteredReqs, err := s.deleteSectorsThatExist(ctx, reqs)
	if err != nil {
		return nil, err
	}

	for _, req := range filteredReqs {
		builders = append(
			builders,
			s.Client.Sector.Create().SetCoords(req.Coords).SetDescription(req.Description),
		)
	}

	created, err := s.Client.Sector.CreateBulk(builders...).Save(ctx)
	if ent.IsConstraintError(err) {
		return nil, ErrSectorExist
	} else if err != nil {
		return nil, err
	}

	return created, nil
}

// CreateSectors
//
// @Summary Create Sectors
//
// @Description create sectors
//
// @Router /v1/sectors [post]
//
// @Security ApiKeyAuth
//
// @Tags sector
//
// @Accept json
//
// @Produce json
//
// @Param body body []sector.CreateSectorReq true "body"
//
// @Success 201 {array} sector.Sector
//
// @Failure 500 {string} srting
//
// @Failure 400 {string} srting
func (s SectorController) CreateSectors(c *gin.Context) {
	var reqs []CreateSectorReq
	{
		if err := c.ShouldBindJSON(&reqs); err != nil {
			c.String(http.StatusBadRequest, "Unexpected body")
			c.Abort()
			return
		}
	}

	created, err := s.createALot(c, reqs)
	if err == ErrSectorExist {
		c.String(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	} else if err != nil {
		log.WithFields(newLogFields("CreateALot", err)).Error("Failed to create sectors")
		c.String(http.StatusInternalServerError, "Failed to create sectors")
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, NewSectors(created))
}

type UpdateAdditionalDescriptionReq struct {
	ID                    int     `json:"-" uri:"id"`
	AdditionalDescription *string `json:"additionalDescription,omitempty"`
}

type AddAdditionalDescriptionReq struct {
	SectorID              int    `uri:"id" swaggerignore:"true" json:"-"`
	ProfileID             int    `uri:"profile_id" swaggerignore:"true" json:"-"`
	AdditionalDescription string `json:"additionalDescription"`
}

// AddAdditionalDescription
//
// @Summary Update additional description
//
// @Router /v1/sector/{sector_id}/profile/{profile_id} [put]
//
// @Security ApiKeyAuth
//
// @Param sector_id path integer true "id of sector"
//
// @Param profile_id path integer true "id of propfile"
//
// @Tags sector
//
// @Accept json
//
// @Produce json
//
// @Param body body sector.AddAdditionalDescriptionReq true "body"
//
// @Success 200 {object} sector.AdditionalDescription
func (s SectorController) AddAdditionalDescription(c *gin.Context) {
	var req AddAdditionalDescriptionReq
	{
		if err := c.ShouldBindUri(&req); err != nil {
			c.String(http.StatusBadRequest, "Bad id's")
			c.Abort()
			return
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.String(http.StatusBadRequest, "Unexpected body")
			c.Abort()
			return
		}
	}

	adj, err := s.Client.AdjacentTable.
		Query().
		Where(
			adjacenttable.SectorID(req.SectorID),
			adjacenttable.ProfileID(req.ProfileID),
		).
		WithProfile(
			func(pq *ent.ProfileQuery) {
				pq.WithDirection(
					func(dq *ent.DirectionQuery) {
						dq.WithInstitute()
					},
				)
			},
		).
		Only(c)
	if ent.IsNotFound(err) {
		c.String(http.StatusNotFound, "Not found additional description with this sector id and profile id")
		c.Abort()
		return
	} else if err != nil {
		log.WithFields(newLogFields("AddAdditionalDescription", err)).Error("Failed to create sectors")
		c.String(http.StatusInternalServerError, "Failed to create sectors")
		c.Abort()
		return
	}

	claims := jwt.ExtractClaims(c)
	if claims["role"].(string) == string(role.ADMIN) {
		if float64(adj.Edges.Profile.Edges.Direction.InstituteID) != claims["intstituteId"].(float64) {
			c.String(http.StatusForbidden, "You are not admin of this institute")
			c.Abort()
			return
		}
	}

	upd, err := adj.Update().SetAdditionalDescription(req.AdditionalDescription).Save(c)
	if err != nil {
		log.WithFields(newLogFields("AddAdditionalDescription", err)).Error("Failed to create sectors")
		c.String(http.StatusInternalServerError, "Failed to create sectors")
		c.Abort()
		return
	}

	upd.Edges = adj.Edges

	c.JSON(http.StatusOK, NewAdditionalDescription(upd))
}
