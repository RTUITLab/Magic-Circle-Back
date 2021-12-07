package institute

import (
	"context"
	"net/http"

	"github.com/0B1t322/Magic-Circle/controllers/utils"
	"github.com/0B1t322/Magic-Circle/ent"
	"github.com/0B1t322/Magic-Circle/ent/institute"
	"github.com/0B1t322/Magic-Circle/ent/variant"
	. "github.com/0B1t322/Magic-Circle/models/institute"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)


func newLogFields(method string, err error) log.Fields {
	return log.Fields{
		"controller": "institute",
		"method": method,
		"err": err,
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
	get, err := i.Client.Institute.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	return get, nil
}

type GetInstitutesReq struct {

}

type GetInstitutesResp struct {
	Ins		[]Institute		`json:"institutes"`
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
	if err := utils.DeleteVariant(c, p.Client, variant.HasInsituteWith(institute.ID(req.ID))); ent.IsNotFound(err) {
		// Pass
	} else if err != nil {
		log.WithFields(newLogFields("Delete", err)).Error("Failed to delete Institute")
		c.String(http.StatusInternalServerError, "Failed to delete Institute")
		c.Abort()
		return
	}
	
	if err := p.Client.Direction.DeleteOneID(req.ID).Exec(c); ent.IsNotFound(err) {
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