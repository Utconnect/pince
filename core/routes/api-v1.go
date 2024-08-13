package routes

import (
	"github.com/gin-gonic/gin"
	"pince/core/controllers"
	"pince/core/repositories"
	"pince/database"
	"pince/database/drivers"
)

// swagger:route GET /teachers/:id
func NewApiRoutes(router *gin.Engine) {
	databaseConnection := database.Connection{
		Driver: &drivers.PostgresDriver{},
	}
	databaseConnection.Connect()

	apiRouterGroup := router.Group("/api")
	apiV1RouterGroup := apiRouterGroup.Group("/v1")
	{
		fileController := controllers.FileController{
			Repository: repositories.FileRepository{
				Connection: databaseConnection,
			},
		}

		apiV1RouterGroup.POST("/files/upload", fileController.Create)
		apiV1RouterGroup.GET("/files/:id/data", fileController.ReadData)
	}
}
