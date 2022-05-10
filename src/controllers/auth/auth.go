package auth

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/0B1t322/Magic-Circle/ent"
	"github.com/0B1t322/Magic-Circle/ent/admin"
	"github.com/0B1t322/Magic-Circle/ent/superadmin"
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

var (
	ErrLoginExist = errors.New("Login exist")
)

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
					ID:          admin.ID,
					Role:        role.ADMIN,
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
	superAdmin, err := a.Client.SuperAdmin.Query().
		Where(superadmin.Login(a.SuperAdminLogin)).
		Only(ctx)
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

	if err := a.loginIsExist(c, req.Login); err == ErrLoginExist {
		c.String(http.StatusBadRequest, "Admin with this login exisst")
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

type GetAdminResp struct {
	ID    int    `json:"id"`
	Login string `json:"login"`
}

type GetAdminsResp struct {
	Admins []GetAdminResp `json:"admins"`
}

// GetAdmins
//
// @Summary get admins
//
// @Router /v1/auth/admin [get]
//
// @Produce json
//
// @Tags auth
//
// @Success 200 {object} auth.GetAdminsResp
//
// @Security ApiKeyAuth
func (a AuthController) GetAdmins(c *gin.Context) {
	admins, err := a.Client.Admin.Query().All(c)
	if err != nil {
		log.WithFields(newLogFields("GetAdmins", err)).Error("Failed to get Admins")
		c.String(http.StatusInternalServerError, "Failed to get Admins")
		c.Abort()
		return
	}

	var adminsResp GetAdminsResp
	for _, admin := range admins {
		adminsResp.Admins = append(adminsResp.Admins, GetAdminResp{ID: admin.ID, Login: admin.Login})
	}

	c.JSON(
		http.StatusOK,
		adminsResp,
	)
}

type ChangePasswordReq struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
	AdminID     int    `json:"-" uri:"id" swaggerignore:"true"`
}

// ChangeAdminPassword
//
// @Summary change admin password
//
// @Description superadmin can pass only new password
//
// @Description admin change only their password and should pass their old password
//
// @Router /v1/auth/admin/{id} [put]
//
// @Param id path integer true "admin id"
//
// @Param body body auth.ChangePasswordReq true "body"
//
// @Accept json
//
// @Produce json
//
// @Tags auth
//
// @Success 200
//
// @Security ApiKeyAuth
func (a AuthController) ChangeAdminPassword(c *gin.Context) {
	var req ChangePasswordReq
	{
		if err := c.ShouldBindUri(&req); err != nil {
			c.String(http.StatusBadRequest, "Unexpected id")
			c.Abort()
			return
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.String(http.StatusBadRequest, "Unexpected body")
			c.Abort()
			return
		}
	}

	admin, err := a.Client.Admin.Get(c, req.AdminID)
	if ent.IsNotFound(err) {
		c.String(http.StatusNotFound, "admin not found")
		c.Abort()
		return
	} else if err != nil {
		log.WithFields(newLogFields("ChangeAdminPassword", err)).Error("Failed to Change Admin Password")
		c.String(http.StatusInternalServerError, "Failed to Change Admin Password")
		c.Abort()
		return
	}

	claims := jwt.ExtractClaims(c)

	if claims["userId"].(float64) != float64(admin.ID) {
		c.Status(http.StatusForbidden)
		c.Abort()
		return
	}

	if claims["role"].(string) == string(role.ADMIN) {
		if req.OldPassword != admin.Password {
			c.String(http.StatusForbidden, "Old Password not eq password")
			c.Abort()
			return
		}
	}

	if err := admin.Update().SetPassword(req.NewPassword).Exec(c); err != nil {
		c.String(http.StatusInternalServerError, "Failed to Change Admin Password")
		c.Abort()
		return
	}

	c.Status(http.StatusOK)
}

type CreateSuperAdminResp struct {
	ID    int    `json:"id"`
	Login string `json:"login"`
}


// CreateSuperAdmin
//
// @Summary create super admin
//
// @Router /v1/auth/superadmin [post]
//
// @Param body body auth.CreateAdminReq true "body"
//
// @Accept json
//
// @Produce json
//
// @Tags auth
//
// @Success 200 {object} auth.CreateAdminResp 
//
// @Security ApiKeyAuth
func (a AuthController) CreateSuperAdmin(c *gin.Context) {
	var req CreateAdminReq
	{
		if err := c.ShouldBindJSON(&req); err != nil {
			log.Info(err)
			c.String(http.StatusBadRequest, "bad body")
			c.Abort()
			return
		}
	}

	if err := a.loginIsExist(c, req.Login); err == ErrLoginExist {
		c.String(http.StatusBadRequest, "SuperAdmin with this login exisst")
		c.Abort()
		return
	} else if err != nil {
		log.WithFields(newLogFields("CreateSuperAdmin", err)).Error("Failed to create SuperAdmin")
		c.String(http.StatusInternalServerError, "Failed to create SuperAdmin")
		c.Abort()
		return
	}

	superAdmin, err := a.Client.SuperAdmin.Create().
		SetLogin(req.Login).
		SetPassword(req.Password).
		Save(c)
	if ent.IsConstraintError(err) {
		c.String(http.StatusBadRequest, "SuperAdmin with this login exisst")
		c.Abort()
		return
	} else if err != nil {
		log.WithFields(newLogFields("CreateSuperAdmin", err)).Error("Failed to create SuperAdmin")
		c.String(http.StatusInternalServerError, "Failed to create SuperAdmin")
		c.Abort()
		return
	}

	c.JSON(
		http.StatusCreated,
		CreateAdminResp{
			ID: superAdmin.ID,
			Login: superAdmin.Login,
		},
	)
}

// GetSuperAdmins
//
// @Summary get admins
//
// @Router /v1/auth/superadmin [get]
//
// @Produce json
//
// @Tags auth
//
// @Success 200 {object} auth.GetAdminsResp
//
// @Security ApiKeyAuth
func (a AuthController) GetSuperAdmins(c *gin.Context) {
	admins, err := a.Client.SuperAdmin.Query().All(c)
	if err != nil {
		log.WithFields(newLogFields("GetAdmins", err)).Error("Failed to get Admins")
		c.String(http.StatusInternalServerError, "Failed to get Admins")
		c.Abort()
		return
	}

	var adminsResp GetAdminsResp
	for _, admin := range admins {
		adminsResp.Admins = append(adminsResp.Admins, GetAdminResp{ID: admin.ID, Login: admin.Login})
	}

	c.JSON(
		http.StatusOK,
		adminsResp,
	)	
}

type DeleteAdminReq struct {
	ID	int `uri:"id"`
}

// DeleteAdmin
//
// @Summary delete admin
//
// @Router /v1/auth/admin/{id} [delete]
//
// @Produce json
//
// @Tags auth
//
// @Success 200
//
// @Security ApiKeyAuth
func (a AuthController) DeleteAdmin(c *gin.Context) {
	var req DeleteAdminReq
	{
		if err := c.ShouldBindUri(&req); err != nil {
			c.String(http.StatusBadRequest, "Unexpected id")
			c.Abort()
			return
		}
	}

	err := a.Client.Admin.DeleteOneID(req.ID).Exec(c)
	if ent.IsNotFound(err) {
		c.String(http.StatusNotFound, "admin not found")
		c.Abort()
		return		
	} else if err != nil {
		c.String(http.StatusInternalServerError, "failed to delete admin")
		c.Abort()
		return
	}

	c.Status(http.StatusOK)
}

// DeleteSuperAdmin
//
// @Summary delete super admin
//
// @Router /v1/auth/superadmin/{id} [delete]
//
// @Produce json
//
// @Tags auth
//
// @Success 200
//
// @Security ApiKeyAuth
func (a AuthController) DeleteSuperAdmin(c *gin.Context) {
	var req DeleteAdminReq
	{
		if err := c.ShouldBindUri(&req); err != nil {
			c.String(http.StatusBadRequest, "Unexpected id")
			c.Abort()
			return
		}
	}

	claims := jwt.ExtractClaims(c)

	if claims["userId"].(float64) == float64(req.ID) {
		c.String(http.StatusBadRequest, "You can't delete yourself")
		c.Abort()
		return
	}

	err := a.Client.SuperAdmin.DeleteOneID(req.ID).Exec(c)
	if ent.IsNotFound(err) {
		c.String(http.StatusNotFound, "admin not found")
		c.Abort()
		return		
	} else if err != nil {
		c.String(http.StatusInternalServerError, "failed to delete admin")
		c.Abort()
		return
	}

	c.Status(http.StatusOK)
}

func (a AuthController) loginIsExist(ctx context.Context, login string) error {
	_, err := a.Client.Admin.Query().Where(
		admin.Login(login),
	).Only(ctx)
	if ent.IsNotFound(err) {
		// Pass
	} else if err != nil {
		return err
	} else if err == nil {
		return ErrLoginExist
	}

	_, err = a.Client.SuperAdmin.Query().Where(
		superadmin.Login(login),
	).Only(ctx)
	if ent.IsNotFound(err) {
		// Pass
	} else if err != nil {
		return err
	} else if err == nil {
		return ErrLoginExist
	}

	return nil
}