package controllers

import (
	"Reverse-proxy/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateAmf(mgmt *models.Management) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var amf models.Amf

		err := ctx.ShouldBindJSON(&amf)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// create amf in management in memory
		mgmt.CreateAmf(amf)

		// produto criado com sucesso.
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": amf})
	}

}

func UpdateAmfState(mgmt *models.Management) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var amf models.Amf

		// id of amf
		name := ctx.Params.ByName("name")

		// status of amf
		if err := ctx.ShouldBindJSON(&amf); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		// update amf in management
		result, err := mgmt.UpdateAmfState(amf, name)
		if err {
			// return amf
			ctx.JSON(http.StatusOK, result)
		} else {
			// return failure.
			ctx.JSON(http.StatusNotFound, result)
		}
	}
}
