package main

import (
	"Reverse-proxy/internal/models"
	"Reverse-proxy/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	// init management
	mgmt := models.InitMgmt()

	// init routes
	router := gin.Default()
	rotasV1 := router.Group("/api/v1")
	routes.CreateRouteAmf(rotasV1, mgmt)

	// init sctp server to handle connections

	// init http server in
	router.Run(":8080")
}
