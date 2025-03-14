package routes

import (
	"limitify/handlers"
	"limitify/middleware"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(r *gin.Engine) {
	r.POST("/signup", handlers.Signup)
	r.POST("/login", handlers.Login)

	admin := r.Group("/admin")
	admin.Use(middleware.AuthMiddleware())
	admin.POST("/set-rate-limit", handlers.SetGlobalRateLimit)
	admin.GET("/get-rate-limit", handlers.GetGlobalRateLimit)

	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware(), middleware.RateLimiterMiddleware())
	protected.GET("/protected", handlers.ProtectedResource)
}
