package utils

import (
	"github.com/gin-gonic/gin"
)

func HandleResponse(ctx *gin.Context, responseMap map[string]interface{}, statusCode int) {
	responseBody := gin.H{}
	for key, value := range responseMap {
		responseBody[key] = value
	}
	ctx.JSON(statusCode, responseBody)
}

func ErrorHandler(ctx *gin.Context, err error, errorMessage string, statusCode int) {
	if err != nil {
		HandleResponse(ctx, GenerateMapForResponseType("Error", errorMessage, err.Error()), statusCode)
		panic(err)
	}
}
