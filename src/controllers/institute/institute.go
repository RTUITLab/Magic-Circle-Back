package institute

import (
	"context"
	"errors"
	"net/http"

	"github.com/0B1t322/Magic-Circle/ent"
	"github.com/0B1t322/Magic-Circle/ent/institute"
	. "github.com/0B1t322/Magic-Circle/models/institute"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var (
	InstituteNotFound = errors.New("Institute not found")
	InstituteExist    = errors.New("Institute with this name exist")
)

func newLogFields(method string, err error) log.Fields {
	return log.Fields{
		"controller": "institute",
		"method":     method,
		"err":        err,
	}
}

type InstituteController struct {
	Client *ent.Client
}

func New(client *ent.Client) *InstituteController {
	return &InstituteController{
		Client: client,
	}
}

func (i InstituteController) getAll(ctx context.Context) ([]*ent.Institute, error) {
	get, err := i.Client.Institute.Query().
		WithDirections(
			func(dq *ent.DirectionQuery) {
				dq.WithProfile()
			},
		).
		All(ctx)
	if err != nil {
		return nil, err
	}

	return get, nil
}

type GetInstitutesReq struct {
}

type GetInstitutesResp struct {
	Ins []Institute `json:"institutes"`
}

// GetAll
//
// @Summary Get all institutes
//
// @Description return all institutes
//
// @Router /v1/institute [get]
//
// @Produce json
//
// @Success 200 {object} institute.GetInstitutesResp
//
// @Failure 500 {string} srting
func (i InstituteController) GetAll(c *gin.Context) {
	is, err := i.getAll(c)
	if err != nil {
		log.WithFields(newLogFields("GetAll", err)).Error("Failed to get institutes")
		c.String(http.StatusInternalServerError, "Failed to get institutes")
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, GetInstitutesResp{Ins: InstitutesFromEnt(is)})
}

type DeleteInstituteByID struct {
	ID int `json:"-" uri:"id"`
}

// DeleteByID
//
// @Summary Delete Institute by id
//
// @Description Delete Institute by id
//
// @Router /v1/institute/{id} [delete]
//
// @Param id path int true "id of institute"
//
// @Produce json
//
// @Success 200
//
// @Failure 404 {string} string
//
// @Failure 400 {string} string
//
// @Failure 500 {string} srting
func (p InstituteController) DeleteByID(c *gin.Context) {
	var req DeleteInstituteByID
	{
		if err := c.ShouldBindUri(&req); err != nil {
			log.WithFields(newLogFields("Delete", err)).Error("Failed to delete Institute")
			c.String(http.StatusInternalServerError, "Failed to delete Institute")
			c.Abort()
			return
		}
	}
	// Если у направления есть профиль то нельзя удалить
	if _, err := p.Client.Institute.Query().Where(
		institute.Or(
			institute.HasDirections(),
		),
	).All(c); ent.IsNotFound(err) {
		// Pass
		// Можно удалять мы не нашли профилей
	} else if !ent.IsNotFound(err) {
		c.String(http.StatusBadRequest, "Institute has profiles or directions")
		c.Abort()
		return
	} else if err != nil {
		log.WithFields(newLogFields("Delete", err)).Error("Failed to delete Institute")
		c.String(http.StatusInternalServerError, "Failed to delete Institute")
		c.Abort()
		return
	}

	if err := p.Client.Institute.DeleteOneID(req.ID).Exec(c); ent.IsNotFound(err) {
		c.String(http.StatusNotFound, "Failed to delete Institute")
		c.Abort()
		return
	} else if err != nil {
		log.WithFields(newLogFields("Delete", err)).Error("Failed to delete Institute")
		c.String(http.StatusInternalServerError, "Failed to delete Institute")
		c.Abort()
		return
	}

	c.Status(http.StatusOK)
}

type UpdateInstituteReq struct {
	ID   int    `json:"-" uri:"id" swaggerignore:"true"`
	Name string `json:"name"`
}

type UpdateInstituteResp struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}


// UpdateInstitute
// 
// @Summary Update institute
// 
// @Description update institute
// 
// @Router /v1/institute/{id} [put]
// 
// @Accept json
// 
// @Produce json
// 
// @Param id path int true "id of institute"
// 
// @Param body body institute.UpdateInstituteReq true "body"
// 
// @Success 200 {object} institute.UpdateInstituteResp
// 
// @Failure 400
// @Failure 404
// @Failure 500
func (i InstituteController) UpdateInstitute(
	c *gin.Context,
) {
	var req UpdateInstituteReq
	{
		if err := c.ShouldBindUri(&req); err != nil {
			c.String(http.StatusBadRequest, "Unexpected id")
			c.Abort()
			return
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.String(http.StatusBadRequest, "Unexpected id")
			c.Abort()
			return
		}
	}

	updated, err := i.Client.Institute.UpdateOneID(req.ID).SetName(req.Name).Save(c)
	if ent.IsNotFound(err) {
		c.String(http.StatusNotFound, InstituteNotFound.Error())
		c.Abort()
		return
	} else if ent.IsConstraintError(err) {
		c.String(http.StatusBadRequest, InstituteExist.Error())
		c.Abort()
		return
	} else if err != nil {
		log.WithFields(newLogFields("Update", err)).Error("Failed to update Institute")
		c.String(http.StatusInternalServerError, "Failed to update Institute")
		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		UpdateInstituteResp{
			ID: updated.ID,
			Name: updated.Name,
		},
	)

}
