package profile

import (
	"context"
	"net/http"

	"github.com/0B1t322/Magic-Circle/controllers/utils"
	"github.com/0B1t322/Magic-Circle/ent"
	"github.com/0B1t322/Magic-Circle/ent/profile"
	"github.com/0B1t322/Magic-Circle/ent/variant"
	. "github.com/0B1t322/Magic-Circle/models/profile"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func newLogFields(method string, err error) log.Fields {
	return log.Fields{
		"controller": "profile",
		"method":     method,
		"err":        err,
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
	Profiles []Profile `json:"profiles"`
}

// GetAll
//
// @Summary Get all profiles
//
// @Description return all profiles
//
// @Router /v1/profile [get]
//
// @Produce json
//
// @Success 200 {object} profile.GetAllProfilesResp
//
// @Failure 500 {string} srting
func (p ProfileController) GetAll(c *gin.Context) {
	ps, err := p.getAll(c)
	if err != nil {
		log.WithFields(newLogFields("GetAll", err)).Error("Failed to get profiles")
		c.String(http.StatusInternalServerError, "Failed to get profiles")
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, GetAllProfilesResp{Profiles: ProfilesFromEnt(ps)})
}

type DeleteProfileByID struct {
	ID int `json:"-" uri:"id"`
}

// DeleteByID
//
// @Summary Delete profile by id
//
// @Description Delete profile by id
//
// @Router /v1/profile/{id} [delete]
//
// @Param id path int true "id of profile"
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
func (p ProfileController) DeleteByID(c *gin.Context) {
	var req DeleteProfileByID
	{
		if err := c.ShouldBindUri(&req); err != nil {
			log.WithFields(newLogFields("Delete", err)).Error("Failed to delete profile")
			c.String(http.StatusInternalServerError, "Failed to delete profile")
			c.Abort()
			return
		}
	}
	if err := utils.DeleteVariant(c, p.Client, variant.HasProfileWith(profile.ID(req.ID))); ent.IsNotFound(err) {
		// Pass
	} else if err != nil {
		log.WithFields(newLogFields("Delete", err)).Error("Failed to delete profile")
		c.String(http.StatusInternalServerError, "Failed to delete profile")
		c.Abort()
		return
	}

	if err := p.Client.Profile.DeleteOneID(req.ID).Exec(c); ent.IsNotFound(err) {
		c.String(http.StatusNotFound, "Failed to delete profile")
		c.Abort()
		return
	} else if err != nil {
		log.WithFields(newLogFields("Delete", err)).Error("Failed to delete profile")
		c.String(http.StatusInternalServerError, "Failed to delete profile")
		c.Abort()
		return
	}

	c.Status(http.StatusOK)
}
