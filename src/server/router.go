package server

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	_ "github.com/0B1t322/Magic-Circle/docs"
)

func NewRouter(c *Controllers) *gin.Engine {
	router := gin.New()

	baseRouter := router.Group("/api/magic-circle")

	baseRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := baseRouter.Group("/v1")
	{
		sector := v1.Group("/sector")
		{
			sector.POST("", c.Sector.Create)
			sector.PUT("/:id", c.Sector.Update)
			sector.GET("", c.Sector.GetAll)
		}

		profile := v1.Group("/profile")
		{
			profile.GET("", c.Profile.GetAll)
		}

		institute := v1.Group("/institute")
		{
			institute.GET("", c.Institute.GetAll)
		}

		direction := v1.Group("/direction")
		{
			direction.GET("", c.Direction.GetAll)
		}

		adjacenttable := v1.Group("/adjacenttable")
		{
			adjacenttable.POST("", c.AdjacentTable.Create)
		}
	}

	return router
}