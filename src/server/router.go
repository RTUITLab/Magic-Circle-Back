package server

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	_ "github.com/0B1t322/Magic-Circle/docs"
)

func NewRouter(c *Controllers) *gin.Engine {
	router := gin.New()
	router.SetTrustedProxies(nil)
	baseRouter := router.Group("/api/magic-circle")

	baseRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := baseRouter.Group("/v1")
	{
		sector := v1.Group("/sector")
		{
			sector.
				GET("", c.Sector.GetAll)

			sector.
				POST(
				"", 
					c.Auth.AuthMiddleare(),
					c.Auth.IsSuperAdminMiddleware,
					c.Sector.Create,
				)

			sector.
				PUT(
					"/:id", 
					c.Auth.AuthMiddleare(),
					c.Auth.IsSuperAdminMiddleware,
					c.Sector.Update,
				)

			sector.
				DELETE(
					"/:id", 
					c.Auth.AuthMiddleare(),
					c.Auth.IsSuperAdminMiddleware,
					c.Sector.DeleteSector,
				)

			sector.
				PUT(
					"/:id/profile/:profile_id",
					c.Auth.AuthMiddleare(),
					c.Sector.AddAdditionalDescription,
				)
		}

		sectors := v1.Group("/sectors")
		{
			sectors.
				POST(
					"",
					c.Auth.AuthMiddleare(),
					c.Auth.IsSuperAdminMiddleware,
					c.Sector.CreateSectors,
				)
		}

		profile := v1.Group("/profile")
		{
			profile.GET("", c.Profile.GetAll)

			profile.
				DELETE(
					"/:id",
					c.Auth.AuthMiddleare(),
					c.Profile.DeleteByID,
				)

			profile.
				PUT(
					"/:id",
					c.Auth.AuthMiddleare(),
					c.Profile.UpdateProfile,
				)
		}

		institute := v1.Group("/institute")
		{
			institute.GET("", c.Institute.GetAll)
			institute.
				DELETE(
					"/:id", 
					c.Auth.AuthMiddleare(),
					c.Institute.DeleteByID,
				)

			institute.
				PUT(
					"/:id", 
					c.Auth.AuthMiddleare(),
					c.Institute.UpdateInstitute,
				)
		}

		direction := v1.Group("/direction")
		{
			direction.GET("", c.Direction.GetAll)

			direction.
				DELETE(
					"/:id", 
					c.Auth.AuthMiddleare(),
					c.Direction.DeleteByID,
				)

			direction.
				PUT(
					"/:id",
					c.Auth.AuthMiddleare(),
					c.Direction.UpdateDirection,
				)
		}


		auth := v1.Group("/auth")
		{
			auth.
				POST("/login", c.Auth.LoginHandler)

			auth.
				GET("/refreshToken", c.Auth.RefreshHandler)
			auth.
				POST(
					"/admin/:intstitute_id", 
					c.Auth.AuthMiddleare(),
					c.Auth.IsSuperAdminMiddleware,
					c.Auth.CreateAdmin,
				)

			auth.
				GET(
					"/admin",
					c.Auth.AuthMiddleare(),
					c.Auth.IsSuperAdminMiddleware,
					c.Auth.GetAdmins,
				)
			
			auth.
				PUT(
					"/admin/:id",
					c.Auth.AuthMiddleare(),
					c.Auth.ChangeAdminPassword,
				)
			
			auth.
				POST(
					"/superadmin",
					c.Auth.AuthMiddleare(),
					c.Auth.IsSuperAdminMiddleware,
					c.Auth.CreateSuperAdmin,
				)

			auth.
				GET(
					"/superadmin",
					c.Auth.AuthMiddleare(),
					c.Auth.IsSuperAdminMiddleware,
					c.Auth.GetSuperAdmins,
				)
			
			auth.
				DELETE(
					"/admin/:id",
					c.Auth.AuthMiddleare(),
					c.Auth.IsSuperAdminMiddleware,
					c.Auth.DeleteAdmin,
				)
			
			auth.
				DELETE(
					"/superadmin/:id",
					c.Auth.AuthMiddleare(),
					c.Auth.IsSuperAdminMiddleware,
					c.Auth.DeleteSuperAdmin,
				)
		}
		

		v1.
			POST(
				"/", 
				c.Auth.AuthMiddleare(),
				c.Root.CreateInstDirProf,
			)

		v1.
			DELETE(
				"/",
				c.Auth.AuthMiddleare(),
				c.Root.DeleteRelate,
			)
	}

	return router
}