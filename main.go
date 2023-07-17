package main

import (
	"log"
	"tonothan/sangaku-pipeline-server/configs"
	"tonothan/sangaku-pipeline-server/routes"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	port := ":" + configs.EnvPort()

	// Run routes
	routes.ManifestRoutes(router)

	router.Run(port)

	log.Fatal(autotls.Run(router, configs.EnvBaseURI()))
}
