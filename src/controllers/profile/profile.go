package profile

import (
	"context"
	"errors"
	"net/http"

	"github.com/0B1t322/Magic-Circle/ent"
	"github.com/0B1t322/Magic-Circle/ent/adjacenttable"
	"github.com/0B1t322/Magic-Circle/ent/profile"
	. "github.com/0B1t322/Magic-Circle/models/profile"
	"github.com/0B1t322/Magic-Circle/models/role"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var (
	ProfileNotFound = errors.New("Profile not found")
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
// @Tags profile
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
// @Security ApiKeyAuth
// 
// @Tags profile
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

	claims := jwt.ExtractClaims(c)
	if claims["role"].(string) == string(role.ADMIN) {
		getProf, err := p.Client.Profile.Query().
			WithDirection().
			Where(
				profile.ID(req.ID),
			).
			Only(c)
		if ent.IsNotFound(err) {
			// Pass
		} else if err != nil {
			log.WithFields(newLogFields("DeleteByID", err)).Error()
			c.Status(http.StatusInternalServerError)
			c.Abort()
			return
		}

		if float64(getProf.Edges.Direction.InstituteID) != claims["intstituteId"].(float64) {
			c.String(http.StatusForbidden, "You are not superadmin or admin of this institute")
			c.Abort()
			return
		}
	}

	if _, err := p.Client.AdjacentTable.Delete().Where(adjacenttable.HasProfileWith(profile.ID(req.ID))).Exec(c); ent.IsNotFound(err) {
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

type UpdateProfileReq struct {
	ID   int    `json:"-" uri:"id" swaggerignore:"true"`
	Name string `json:"name"`
}

type UpdateProfileResp struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}


// UpdateProfile
// 
// @Summary Update profile
// 
// @Description update profile
// 
// @Router /v1/profile/{id} [put]
// 
// @Security ApiKeyAuth
// 
// @Tags profile
// 
// @Accept json
// 
// @Produce json
// 
// @Param id path int true "id of profile"
// 
// @Param body body profile.UpdateProfileReq true "body"
// 
// @Success 200 {object} profile.UpdateProfileResp
// 
// @Failure 400
// @Failure 404
// @Failure 500
func (p ProfileController) UpdateProfile(c *gin.Context) {
	var req UpdateProfileReq
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
	claims := jwt.ExtractClaims(c)
	if claims["role"].(string) == string(role.ADMIN) {
		getProf, err := p.Client.Profile.Query().
			WithDirection().
			Where(
				profile.ID(req.ID),
			).
			Only(c)
		if ent.IsNotFound(err) {
			// Pass
		} else if err != nil {
			log.WithFields(newLogFields("UpdateProfile", err)).Error()
			c.Status(http.StatusInternalServerError)
			c.Abort()
			return
		}

		if float64(getProf.Edges.Direction.InstituteID) != claims["intstituteId"].(float64) {
			c.String(http.StatusForbidden, "You are not superadmin or admin of this institute")
			c.Abort()
			return
		}
	}

	updated, err := p.Client.Profile.UpdateOneID(req.ID).SetName(req.Name).Save(c)
	if ent.IsNotFound(err) {
		c.String(http.StatusNotFound, ProfileNotFound.Error())
		c.Abort()
		return
	} else if err != nil {
		log.WithFields(newLogFields("Update", err)).Error("Failed to update Profile")
		c.String(http.StatusInternalServerError, "Failed to update Profile")
		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		UpdateProfileResp{
			ID: updated.ID,
			Name: updated.Name,
		},
	)
}