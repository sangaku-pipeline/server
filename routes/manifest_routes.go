package routes

import (
	"tonothan/sangaku-pipeline-server/controllers"

	"github.com/gin-gonic/gin"
)

func ManifestRoutes(router *gin.Engine) {
	router.GET("/", controllers.Ping())
	router.GET("/manifest/:manifestId", controllers.GenerateManifestByID())
	router.GET("/manifest-data/:manifestId", controllers.GenerateManifestByID())
	router.POST("/manifest", controllers.CreateManifestMetadata())
}
