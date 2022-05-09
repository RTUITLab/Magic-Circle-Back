package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/0B1t322/Magic-Circle/ent"
	"github.com/0B1t322/Magic-Circle/ent/admin"
	paylaod "github.com/0B1t322/Magic-Circle/models/jwtpayload"
	"github.com/0B1t322/Magic-Circle/models/role"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func newLogFields(method string, err error) log.Fields {
	return log.Fields{
		"controller": "direction",
		"method":     method,
		"err":        err,
	}
}

var identyKey = "userId"

type AuthController struct {
	Client           *ent.Client
	AccessSecret     string
	RefreshSecret    string
	SuperAdminLogin  string
	SuperAdminPasswd string

	jwtHandlers *jwt.GinJWTMiddleware
}

func New(
	client *ent.Client,
	accessSecret string,
	refreshSecret string,
	superAdminLogin string,
	superAdminPasswd string,
) *AuthController {
	c := &AuthController{
		Client:           client,
		AccessSecret:     accessSecret,
		RefreshSecret:    refreshSecret,
		SuperAdminLogin:  superAdminLogin,
		SuperAdminPasswd: superAdminPasswd,
	}

	c.init()
	return c
}

func (a *AuthController) init() {
	a.initSuperAdmin()
	a.initJwt()
}

type LoginReq struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (a *AuthController) initJwt() {
	var err error
	a.jwtHandlers, err = jwt.New(
		&jwt.GinJWTMiddleware{
			Realm:       "realm",
			Key:         []byte(a.AccessSecret),
			Timeout:     time.Hour,
			MaxRefresh:  7 * 24 * time.Hour, // week
			IdentityKey: identyKey,
			PayloadFunc: func(data interface{}) jwt.MapClaims {
				if v, ok := data.(*paylaod.Payload); ok {
					return jwt.MapClaims{
						identyKey:      v.ID,
						"role":         v.Role,
						"intstituteId": v.InstituteID,
					}
				}
				return jwt.MapClaims{}
			},
			Authenticator: func(c *gin.Context) (interface{}, error) {
				var req LoginReq
				{
					if err := c.ShouldBindJSON(&req); err != nil {
						return "", jwt.ErrMissingLoginValues
					}
				}

				super, err := a.Client.SuperAdmin.Query().Only(c)
				if err != nil {
					log.WithFields(newLogFields("Authenticator", err)).Error("Failed to get super admin")
					return "", err
				}

				if super.Login == req.Login && super.Password == req.Password {
					return &paylaod.Payload{
						ID:   super.ID,
						Role: role.SUPERADMIN,
					}, nil
				}

				admin, err := a.Client.Admin.Query().Where(
					admin.And(
						admin.Login(req.Login),
						admin.Password(req.Password),
					),
				).
				Only(c)
				if ent.IsNotFound(err) {
					return nil, jwt.ErrFailedAuthentication
				} else if err != nil {
					return nil, err
				}

				return &paylaod.Payload{
					ID:   admin.ID,
					Role: role.ADMIN,
					InstituteID: admin.InstituteID,
				}, nil

			},
			Authorizator: func(data interface{}, c *gin.Context) bool {
				log.Info("Auth")
				return true
			},
			TokenLookup:   "header: Authorization",
			TokenHeadName: "Bearer",
			TimeFunc:      time.Now,
		},
	)
	if err != nil {
		panic(err)
	}
}

func (a AuthController) initSuperAdmin() {
	ctx := context.Background()
	superAdmin, err := a.Client.SuperAdmin.Query().Only(ctx)
	if ent.IsNotFound(err) {
		a.createSuperAdmin(ctx)
		return
	} else if err != nil {
		panic(err)
	}

	superAdmin.Update().SetLogin(a.SuperAdminLogin).SetPassword(a.SuperAdminPasswd).ExecX(ctx)
}

func (a AuthController) createSuperAdmin(ctx context.Context) {
	a.Client.SuperAdmin.Create().
		SetLogin(a.SuperAdminLogin).
		SetPassword(a.SuperAdminPasswd).
		SaveX(ctx)
}

type LoginResp struct {
	Expite time.Time `json:"expire"`
	Token  string    `json:"token"`
}

// LoginHandler
//
// @Summary login admin or super admin
// 
// @Tags auth
//
// @Router /v1/auth/login [post]
//
// @Accept json
//
// @Produce json
//
// @Param body body auth.LoginReq true "body"
//
// @Success 200 {object} auth.LoginResp
func (a AuthController) LoginHandler(c *gin.Context) {
	a.jwtHandlers.LoginHandler(c)
}

// RefreshHandler
//
// @Summary login admin or super admin
//
// @Tags auth
// 
// @Router /v1/auth/refreshToken [get]
//
// @Produce json
//
// @Success 200 {object} auth.LoginResp
//
// @Security ApiKeyAuth
func (a AuthController) RefreshHandler(c *gin.Context) {
	a.jwtHandlers.RefreshHandler(c)
}

type CreateAdminReq struct {
	Login       string `json:"login" binding:"required"`
	Password    string `json:"password" binding:"required"`
	InstituteID int    `uri:"intstitute_id" swaggerignore:"true"`
}

type CreateAdminResp struct {
	ID          int    `json:"id"`
	Login       string `json:"login"`
	InstituteID int    `json:"intstituteId"`
}

// CreateAdmin
//
// @Summary create admin
//
// @Router /v1/auth/admin/{intstitute_id} [post]
//
// @Accept json
//
// @Produce json
// 
// @Tags auth
//
// @Param body body auth.CreateAdminReq true "body"
//
// @Param intstitute_id path string true "id of institute"
//
// @Success 201 {object} auth.CreateAdminResp
//
// @Security ApiKeyAuth
func (a AuthController) CreateAdmin(c *gin.Context) {
	var req CreateAdminReq
	{
		if err := c.ShouldBindJSON(&req); err != nil {
			log.Info(err)
			c.String(http.StatusBadRequest, "bad body")
			c.Abort()
			return
		}

		if err := c.ShouldBindUri(&req); err != nil {
			c.String(http.StatusBadRequest, "bad institute id")
			c.Abort()
			return
		}
	}

	inst, err := a.Client.Institute.Get(c, req.InstituteID)
	if ent.IsNotFound(err) {
		c.String(http.StatusNotFound, "Institute not found")
		c.Abort()
		return
	} else if err != nil {
		log.WithFields(newLogFields("CreateAdmin", err)).Error("Failed to create Admin")
		c.String(http.StatusInternalServerError, "Failed to create Admin")
		c.Abort()
		return
	}

	admin, err := a.Client.Admin.Create().
		SetInstitute(inst).
		SetLogin(req.Login).
		SetPassword(req.Password).
		Save(c)
	if ent.IsConstraintError(err) {
		c.String(http.StatusBadRequest, "Admin with this login exisst")
		c.Abort()
		return
	} else if err != nil {
		log.WithFields(newLogFields("CreateAdmin", err)).Error("Failed to create Admin")
		c.String(http.StatusInternalServerError, "Failed to create Admin")
		c.Abort()
		return
	}

	c.JSON(
		http.StatusCreated,
		CreateAdminResp{
			ID:          admin.ID,
			Login:       admin.Login,
			InstituteID: req.InstituteID,
		},
	)
}

func (a AuthController) IsSuperAdminMiddleware(c *gin.Context) {
	claims := jwt.ExtractClaims(c)

	if claims["role"].(string) != string(role.SUPERADMIN) {
		c.String(http.StatusForbidden, "You are not super admin")
		c.Abort()
		return
	}
}

func (a AuthController) IsAdmin(c *gin.Context) {
	claims := jwt.ExtractClaims(c)

	if claims["role"].(string) != string(role.ADMIN) {
		c.String(http.StatusForbidden, "You are not admin")
		c.Abort()
		return
	}
}

func (a AuthController) AuthMiddleare() gin.HandlerFunc {
	return a.jwtHandlers.MiddlewareFunc()
}
