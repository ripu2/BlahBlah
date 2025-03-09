package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ripu2/blahblah/internal/models"
	"github.com/ripu2/blahblah/internal/services"
	"github.com/ripu2/blahblah/internal/utils"
)

func CreateUserHandler(ctx *gin.Context) {

	var user *models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		utils.ErrorHandler(ctx, err, "Invalid user data", http.StatusBadRequest)
		return
	}

	token, err := services.CreateUserService(user)
	if err != nil {
		utils.ErrorHandler(ctx, err, "Failed to create user", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"id":           user.ID,
		"user_name":    user.UserName,
		"first_name":   user.FirstName,
		"last_name":    user.LastName,
		"phone_number": user.PhoneNumber,
		"dob":          user.DOB,
		"created_at":   user.CreatedAt,
		"updated_at":   user.UpdatedAt,
		"token":        token, // Add empty token
	}
	utils.HandleResponse(ctx, utils.GenerateMapForResponseType("data", "User created successfully", response), http.StatusCreated)
}

func LoginUserHandler(ctx *gin.Context) {
	var loginRequest models.LoginRequest
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		utils.ErrorHandler(ctx, err, "Invalid login request", http.StatusBadRequest)
		return
	}

	token, err := services.LoginUserService(&loginRequest)
	if err != nil {
		utils.ErrorHandler(ctx, err, "Failed to login", http.StatusUnauthorized)
		return
	}
	ctx.SetCookie("auth_token", token, 3600*24, "/", "", false, false)
	utils.HandleResponse(ctx, utils.GenerateMapForResponseType("data", "Login successful", map[string]string{"token": token}), http.StatusOK)
}
