package routes

import (
	"tonothan/sangaku-pipeline-server/controllers"

	"github.com/gin-gonic/gin"
)

func ManifestRoutes(router *gin.Engine) {
	router.GET("/", controllers.Ping())
	router.GET("/manifest/:manifestId", controllers.GetManifestMetadata())
	router.GET("/manifest-test", controllers.GetManifestTest())
	router.POST("/manifest", controllers.CreateManifestMetadata())
}
