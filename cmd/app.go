package main

import (
	"Reverse-proxy/config"
	"Reverse-proxy/internal/models"
	"Reverse-proxy/internal/routes"
	"Reverse-proxy/internal/sctp"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	var c config.Config

	cfg := c.GetConf()

	// init management
	mgmt := models.InitMgmt(cfg.Sctp.Ip, 3000, cfg.Sctp.Port)

	// init routes
	router := gin.Default()
	rotasV1 := router.Group("/api/v1")
	routes.CreateRouteAmf(rotasV1, mgmt)

	// init sctp server to handle connections
	sctp.InitServer(mgmt)

	// init http server
	addr := fmt.Sprintf("%s:%d", cfg.Http.Ip, cfg.Http.Port)
	router.Run(addr)
}
