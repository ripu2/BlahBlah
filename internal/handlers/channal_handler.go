package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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
	services.DeleteValueFromCache("com_channels")
}
func GetAllChannelsHandler(ctx *gin.Context) {
	cachedData, err := services.GetValueFromCache("com_channels")
	fmt.Println("Chached", cachedData, err)
	if err == nil && cachedData != "" {
		var channels []models.Channel // 👈 yaha apne model ka type use kar
		_ = json.Unmarshal([]byte(cachedData), &channels)

		utils.HandleResponse(ctx, utils.GenerateMapForResponseType("data: Fetched Channels from Cache", "Fetched Channels from Cache", channels), http.StatusOK)
		return
	}

	channels, err := services.GetAllChannelsService()
	if err != nil {
		utils.ErrorHandler(ctx, err, "Failed to get channels", http.StatusInternalServerError)
		return
	}

	jsonData, _ := json.Marshal(channels)
	_ = services.SetValueInCache("com_channels", string(jsonData))

	utils.HandleResponse(ctx, utils.GenerateMapForResponseType("data: Fetched Channels Successfully", "Fetched Channels Successfully", channels), http.StatusOK)
}

func GetOwnChannelsHandler(ctx *gin.Context) {
	userIDRaw, _ := ctx.Get("userId")
	channels, err := services.GetChannelByIdService(userIDRaw.(int64))
	utils.ErrorHandler(ctx, err, "Failed to get events", http.StatusInternalServerError)
	utils.HandleResponse(ctx, utils.GenerateMapForResponseType("data: Fetched Channels Successfully", "Fetched Channels Successfully", channels), http.StatusOK)

}

func AddUserToCHannelsHandler(ctx *gin.Context) {
	var user models.ChannelUser
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	utils.ErrorHandler(ctx, err, "Bad request", http.StatusBadRequest)
	userIDRaw, _ := ctx.Get("userId")
	user.UserID = userIDRaw.(int64)
	user.Role = "member" // Default role for new users in a channel. Can be customized as per requirements.
	user.ChannelID = id
	err = services.InsertUserInChannelService(&user)
	utils.ErrorHandler(ctx, err, "Failed to create event", http.StatusInternalServerError)
	utils.HandleResponse(ctx, utils.GenerateMapForResponseType("data", "Users  Added to channel successful", user), http.StatusOK)

}
