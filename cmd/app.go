package main

import (
	"Reverse-proxy/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// sctp.InitServer("127.0.0.1", 38412)

	//sctp.InitConn("127.0.0.2", 38412, "127.0.0.1", 38412)

	router := gin.Default()

	// define as rotas padrões
	rotasV1 := router.Group("/api/v1")

	routes.CreateRouteAmf(rotasV1)

	// roda o servidor HTTP na porta 8080.
	router.Run(":8080")
}
