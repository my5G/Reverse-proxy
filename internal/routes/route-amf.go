package routes

import (
	"Reverse-proxy/internal/controllers"
	"github.com/gin-gonic/gin"
)

func CreateRouteAmf(route *gin.RouterGroup) {
	route.POST("/amf", controllers.CreateAmf)
	route.PUT("/amf/:name", controllers.UpdateAmf)
}
