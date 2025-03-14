package handlers

import (
	"net/http"
	"limitify/config"
	"limitify/models"

	"github.com/gin-gonic/gin"
)

const rateLimitCacheKey = "global_rate_limit"
func SetGlobalRateLimit(c *gin.Context){
	var rateLimit models.RateLimit

	if err := c.ShouldBindJSON(&rateLimit); err != nil {
		c.JSON(http.StatusBadRequest , gin.H{
			"error" : "Invalid Request",
		})
		return
	}
	config.DB.Model(&models.RateLimit{}).Where("id = ?" , 1).Updates(rateLimit)

	config.RedisClient.Del(config.Ctx , rateLimitCacheKey)

	c.JSON(http.StatusOK , gin.H{"message" : "global rate limit update successful" })

}


func GetGlobalRateLimit(c *gin.Context){
	var rateLimit models.RateLimit
	config.DB.First(&rateLimit)
	c.JSON(http.StatusOK , rateLimit)
}

