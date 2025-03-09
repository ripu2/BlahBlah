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
	userIDRaw, _ := ctx.Get("userId")
	userID64, _ := userIDRaw.(int64) // ✅ Proper type assertion
	channel.CreatedBy = userID64
	err := ctx.ShouldBindBodyWithJSON(&channel)
	utils.ErrorHandler(ctx, err, "Bad request", http.StatusBadRequest)
	err = services.CreateChanelService(&channel)
	utils.ErrorHandler(ctx, err, "Failed to create event", http.StatusInternalServerError)
	utils.HandleResponse(ctx, utils.GenerateMapForResponseType("data", "Channel Creation successful", channel), http.StatusOK)
}

func GetAllChannelsHandler(ctx *gin.Context) {
	channels, err := services.GetAllChannels()
	utils.ErrorHandler(ctx, err, "Failed to get events", http.StatusInternalServerError)
	utils.HandleResponse(ctx, utils.GenerateMapForResponseType("data", "Fetched Channels Successfully", channels), http.StatusOK)

}

func GetOwnChannelsHandler(ctx *gin.Context) {
	userIDRaw, _ := ctx.Get("userId")
	channels, err := services.GetChannelById(userIDRaw.(int64))
	utils.ErrorHandler(ctx, err, "Failed to get events", http.StatusInternalServerError)
	utils.HandleResponse(ctx, utils.GenerateMapForResponseType("data", "Fetched Channels Successfully", channels), http.StatusOK)

}
