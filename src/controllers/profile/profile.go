package profile

import (
	"context"
	"net/http"

	"github.com/0B1t322/Magic-Circle/ent"
	. "github.com/0B1t322/Magic-Circle/models/profile"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func newLogFields(method string, err error) log.Fields {
	return log.Fields{
		"controller": "profile",
		"method": method,
		"err": err,
	}
}

type ProfileController struct {
	Client *ent.Client
}

func New(client *ent.Client) *ProfileController {
	return &ProfileController{
		Client: client,
	}
}

func (p ProfileController) getAll(ctx context.Context) ([]*ent.Profile, error) {
	return p.Client.Profile.Query().All(ctx)
}

type GetAllProfilesReq struct {

}

type GetAllProfilesResp struct {
	Profiles []Profile	`json:"profiles"`
}

// GetAll
// 
// @Summary Get all profiles
// 
// @Description return all profiles
// 
// @Router /v1/profiles [get]
// 
// @Produce json
// 
// @Success 200 {object} profile.GetAllProfilesResp
// 
// @Failure 500 {string} srting
func (p ProfileController) GetAll(c *gin.Context) {
	ps, err := p.getAll(c)
	if err != nil {
		log.WithFields(newLogFields("GetAll", err)).Error("Failed to get institutes")
		c.String(http.StatusInternalServerError, "Failed to get institutes")
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, GetAllProfilesResp{Profiles: ProfilesFromEnt(ps)})
}