package direction

import (
	"context"
	"errors"
	"net/http"

	"github.com/0B1t322/Magic-Circle/ent"
	"github.com/0B1t322/Magic-Circle/ent/direction"
	"github.com/0B1t322/Magic-Circle/ent/profile"
	. "github.com/0B1t322/Magic-Circle/models/direction"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var (
	DirectionNotFound = errors.New("Direction not found")
)

func newLogFields(method string, err error) log.Fields {
	return log.Fields{
		"controller": "direction",
		"method": method,
		"err": err,
	}
}

type DirectionController struct {
	Client *ent.Client
}

func New(client *ent.Client) *DirectionController {
	return &DirectionController{
		Client: client,
	}
}

func (d DirectionController) getAll(ctx context.Context) ([]*ent.Direction, error) {
	return d.Client.Direction.Query().WithInstitute().WithProfile().All(ctx)
}

type GetDirectionsReq struct {

}

type GetDirectionsResp struct {
	Dirs	[]Direction		`json:"directions"`
}

// GetAll
// 
// @Summary Get all directions
// 
// @Description return all directions
// 
// @Router /v1/direction [get]
// 
// @Produce json
// 
// @Success 200 {object} direction.GetDirectionsResp
// 
// @Failure 500 {string} srting
func (d DirectionController) GetAll(c *gin.Context) {
	ds, err := d.getAll(c)
	if err != nil {
		log.WithFields(newLogFields("GetAll", err)).Error("Failed to get institutes")
		c.String(http.StatusInternalServerError, "Failed to get institutes")
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, GetDirectionsResp{Dirs: DirectionsFromEnt(ds)})
}


type DeleteDirectionByID struct {
	ID int `json:"-" uri:"id"`
}

// DeleteByID
//
// @Summary Delete Direction by id
//
// @Description Delete Direction by id
//
// @Router /v1/direction/{id} [delete]
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
func (p DirectionController) DeleteByID(c *gin.Context) {
	var req DeleteDirectionByID
	{
		if err := c.ShouldBindUri(&req); err != nil {
			log.WithFields(newLogFields("Delete", err)).Error("Failed to delete Direction")
			c.String(http.StatusInternalServerError, "Failed to delete Direction")
			c.Abort()
			return
		}
	}

	// Если у направления есть профиль то нельзя удалить
	if _, err := p.Client.Profile.Query().Where(
		profile.HasDirectionWith(
			direction.ID(req.ID),
		),
	).All(c); ent.IsNotFound(err) {
		// Pass
		// Можно удалять мы не нашли профилей
	} else if !ent.IsNotFound(err) {
		c.String(http.StatusBadRequest, "Direction has profiles")
		c.Abort()
		return
	} else if err != nil {
		log.WithFields(newLogFields("Delete", err)).Error("Failed to delete Direction")
		c.String(http.StatusInternalServerError, "Failed to delete Direction")
		c.Abort()
		return
	}
	
	if err := p.Client.Direction.DeleteOneID(req.ID).Exec(c); ent.IsNotFound(err) {
		c.String(http.StatusNotFound, "Failed to delete Direction")
		c.Abort()
		return
	} else if err != nil {
		log.WithFields(newLogFields("Delete", err)).Error("Failed to delete Direction")
		c.String(http.StatusInternalServerError, "Failed to delete Direction")
		c.Abort()
		return
	}

	c.Status(http.StatusOK)
}

type UpdateDirectionReq struct {
	ID   int    `json:"-" uri:"id" swaggerignore:"true"`
	Name string `json:"name"`
}

type UpdateDirectionResp struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}


// UpdateDirection
// 
// @Summary Update dirction
// 
// @Description update direction
// 
// @Router /v1/direction/{id} [put]
// 
// @Accept json
// 
// @Produce json
// 
// @Param id path int true "id of direction"
// 
// @Param body body direction.UpdateDirectionReq true "body"
// 
// @Success 200 {object} direction.UpdateDirectionResp
// 
// @Failure 400
// @Failure 404
// @Failure 500
func (d DirectionController) UpdateDirection(c *gin.Context) {
	var req UpdateDirectionReq
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

	updated, err := d.Client.Direction.UpdateOneID(req.ID).SetName(req.Name).Save(c)
	if ent.IsNotFound(err) {
		c.String(http.StatusNotFound, DirectionNotFound.Error())
		c.Abort()
		return
	} else if err != nil {
		log.WithFields(newLogFields("Update", err)).Error("Failed to update Direction")
		c.String(http.StatusInternalServerError, "Failed to update Direction")
		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		UpdateDirectionResp{
			ID: updated.ID,
			Name: updated.Name,
		},
	)
}