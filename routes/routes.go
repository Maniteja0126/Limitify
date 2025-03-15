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
	admin.POST("/create-api-key", handlers.CreateApiKey)
	admin.GET("/list-api-keys", handlers.ListAPIKeys)

	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware(), middleware.RateLimiterMiddleware())
	protected.GET("/protected", handlers.ProtectedResource)

	api := r.Group("/api")
	api.Use(middleware.RateLimiterMiddleware())
	api.Use(middleware.AuthMiddleware())
	api.Use(middleware.APIGateway())
	api.Any("/*path")

}
