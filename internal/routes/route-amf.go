package routes

import (
	"Reverse-proxy/internal/controllers"
	"Reverse-proxy/internal/models"
	"github.com/gin-gonic/gin"
)

func CreateRouteAmf(route *gin.RouterGroup, mgmt *models.Management) {
	route.POST("/amf", controllers.CreateAmf(mgmt))
	route.PUT("/amf/state/:name", controllers.UpdateAmfState(mgmt))
}
