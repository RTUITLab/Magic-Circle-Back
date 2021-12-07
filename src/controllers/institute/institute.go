package institute

import (
	"context"
	"net/http"

	"github.com/0B1t322/Magic-Circle/ent"
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