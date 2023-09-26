package routes

import (
	"github.com/gin-gonic/gin"

	"git.tashilcar.com/core/tashilcar/presentation/rest/controllers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(router *gin.Engine) {
	// This is what I needed to add for it to work, this is "docs" in the root of my application
	// generated with swag init
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.POST("/url/create", controllers.CreateAPI)
	router.GET("/urls", controllers.GetAPIs)
	router.PUT("/url/:id", controllers.EnableCheckHealth)
	router.DELETE("/url/:id", controllers.DeleteAPI)
}
