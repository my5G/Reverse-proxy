package controllers

import (
	"Reverse-proxy/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateAmf(ctx *gin.Context) {
	var amf models.Amf

	err := ctx.ShouldBindJSON(&amf)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create amf in management in memory

	// produto criado com sucesso.
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": amf})

}

func UpdateAmf(ctx *gin.Context) {
	var amf models.Amf

	// id da reserva
	_ = ctx.Params.ByName("name")

	// status da reserva.
	if err := ctx.ShouldBindJSON(&amf); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// update amf in management

	// return reservations.
	ctx.JSON(http.StatusOK, amf)
}
