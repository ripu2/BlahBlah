package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ripu2/blahblah/internal/utils"
)

// Middleware to check authentication
func CheckForAuthentication(ctx *gin.Context) {
	// Get token from header
	// token := ctx.Request.Header.Get("Authorization")
	token, _ := ctx.Cookie("auth_token")
	if token == "" {
		utils.ErrorHandler(ctx, errors.New("Unauthenticated"), "Unauthenticated", http.StatusUnauthorized)
		ctx.Abort() // 🔥 Ensure request is stopped
		return
	}

	// Remove "Bearer " prefix if present
	token = strings.TrimPrefix(token, "Bearer ")
	// Verify token
	id, err := utils.VerifyToken(token)
	if err != nil {
		utils.ErrorHandler(ctx, errors.New(err.Error()), "Unauthenticated", http.StatusUnauthorized)
		ctx.Abort() // 🔥 Ensure request is stopped on failure
		return
	}

	// Set userId in context
	ctx.Set("userId", id)
	ctx.Next()
}
