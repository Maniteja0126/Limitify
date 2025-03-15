package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ProtectedResource(c *gin.Context){
	userID , exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized , gin.H{"error" : "User not found"})
		return
	}
	c.JSON(http.StatusOK , gin.H{
		"message" : "Access granted",
		"user" : userID,
	})

}
