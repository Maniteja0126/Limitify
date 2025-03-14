package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"limitify/config"
	"limitify/models"
	"limitify/utils"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const rateLimitCacheKey = "global_rate_limit"


func RateLimiterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context){
		authHeader := c.GetHeader("Authorization")
		if authHeader == ""{
			c.JSON(http.StatusUnauthorized , gin.H{"error" : "Missing token"})
			c.Abort()
			return
		}

		tokenString := strings.Split(authHeader , "Bearer ")[1]
		token , err := utils.VerifyJWT(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized , gin.H{"error" : "Invalid token"})
			c.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		email := claims["email"].(string)

		var rateLimit models.RateLimit
		rateLimitJSON , err := config.RedisClient.Get(config.Ctx , rateLimitCacheKey).Result()
		if err == nil {
			json.Unmarshal([]byte(rateLimitJSON) , &rateLimit)	
		}else{
			config.DB.First(&rateLimit)
			data , _ := json.Marshal(rateLimit)
			config.RedisClient.Set(config.Ctx , rateLimitCacheKey , data , 10*time.Minute)
		}
		

		limit := rateLimit.Requests
		window := time.Duration(rateLimit.TimeWindow) * time.Second


		
		key := fmt.Sprintf("rate_limit:%s" , email)
		count , _ := config.RedisClient.Get(config.Ctx , key ).Int()

		if count >= limit {
			logRequest(email , c.FullPath() , 429 , "Too many requests")
			c.JSON(http.StatusTooManyRequests , gin.H{"error" : "Too many requests"})
			c.Abort()
			return
		}

		pipe := config.RedisClient.TxPipeline()
		pipe.Incr(config.Ctx , key)
		pipe.Expire(config.Ctx , key , window)
		_ , _ = pipe.Exec(config.Ctx)

		logRequest(email , c.FullPath() , 200 , "Request successful")

		c.Next()

	}
}


func logRequest(email , endpoint string , status int  , message string) {
	log := models.RequestLog{
		Email : email,
		Endpoint : endpoint,
		StatusCode:  status,
		Message:  message,
	}
	config.DB.Create(&log)
}
