package middleware

import (
	"limitify/config"
	"limitify/models"
	"limitify/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Token"})
			c.Abort()
			return
		}
		tokenParts := strings.Split(authHeader, "Bearer ")
		if len(tokenParts) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token  from token parts"})
			c.Abort()
			return
		}

		tokenString := tokenParts[1]

		token, err := utils.VerifyJWT(tokenString)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token from the token string"})
			c.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		email, ok := claims["email"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token  from the claims "})
			c.Abort()
			return
		}
		var user models.User
		if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		c.Set("userId", user.ID)
		c.Next()
	}
}
