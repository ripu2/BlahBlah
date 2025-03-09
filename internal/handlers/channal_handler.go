package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ripu2/blahblah/internal/models"
	"github.com/ripu2/blahblah/internal/services"
	"github.com/ripu2/blahblah/internal/utils"
)

func CreateChanelHandler(ctx *gin.Context) {
	var channel models.Channel
	id := ctx.GetInt("userId")
	err := ctx.ShouldBindBodyWithJSON(&channel)
	utils.ErrorHandler(ctx, err, "Bad request", http.StatusBadRequest)
	err = services.CreateChanelService(&channel, id)
	utils.ErrorHandler(ctx, err, "Failed to create event", http.StatusInternalServerError)
}
