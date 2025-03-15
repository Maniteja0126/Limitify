package handlers

import(
	"crypto/rand"
	"encoding/hex"
	"limitify/config"
	"limitify/models"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
)

func GenerateAPIKey() string {
	bytes := make([]byte , 16)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func CreateApiKey(c *gin.Context){
	var request struct {
		BackendUrl string `json:"backend_url"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&request); err != nil{
		c.JSON(http.StatusBadRequest  , gin.H{"error" : "Invalid Request"})
		return
	}

	userID , exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized , gin.H{"error" : "Unauthorized User not exists"  })
		return
	}


	apiKey := GenerateAPIKey()
	newKey := models.APIKey{
		UserId: userID.(uint),
		ApiKey:  apiKey,
		BackendUrl: request.BackendUrl,
		Description: request.Description,
	}

	if err := config.DB.Create(&newKey).Error; err != nil {
		c.JSON(http.StatusInternalServerError , gin.H{"error" : "Failed to create API key"})
		return
	}

	config.RedisClient.Set(config.Ctx , "apikey:"+apiKey , request.BackendUrl , time.Hour * 24 )
	c.JSON(http.StatusOK , gin.H{"api_key" : apiKey})
}



func ListAPIKeys(c *gin.Context){
	userID , exists := c.Get("userId")

	if !exists {
		c.JSON(http.StatusUnauthorized , gin.H{"error" : "Unauthorized"})
		return
	}

	var apiKeys []models.APIKey
	config.DB.Where("user_id = ?" , userID).Find(&apiKeys)
	c.JSON(http.StatusOK , apiKeys)
}
