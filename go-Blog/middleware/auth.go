package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"tugas4/services"
	"tugas4/utils"

	"github.com/gin-gonic/gin"
)

func Authenticate(jwtService services.JWTService, role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := utils.BuildErrorResponse("No token found", http.StatusUnauthorized)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		if !strings.Contains(authHeader, "Bearer ") {
			response := utils.BuildErrorResponse("No token found", http.StatusUnauthorized)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		authHeader = strings.Replace(authHeader, "Bearer ", "", -1)
		token, err := jwtService.ValidateToken(authHeader)
		if err != nil {
			response := utils.BuildErrorResponse("Invalid token", http.StatusUnauthorized)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		if !token.Valid {
			response := utils.BuildErrorResponse("Invalid token", http.StatusUnauthorized)
			c.AbortWithStatusJSON(http.StatusForbidden, response)
			return
		}

		userRole, err := jwtService.GetRoleByToken(string(authHeader))
		fmt.Println("ROLE", userRole)
		if err != nil || (userRole != "admin" && userRole != role) {
			response := utils.BuildErrorResponse("Failed to process request", http.StatusUnauthorized)
			c.AbortWithStatusJSON(http.StatusForbidden, response)
			return
		}

		// get userID from token
		userID, err := jwtService.GetUserIDByToken(authHeader)
		if err != nil {
			response := utils.BuildErrorResponse("Failed to process request", http.StatusUnauthorized)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		fmt.Println("ROLE", userRole)
		c.Set("userID", userID)
		c.Set("token", authHeader)
		c.Next()
	}
}
