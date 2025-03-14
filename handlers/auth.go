package handlers

import (
	"net/http"
	"limitify/config"
	"limitify/models"
	"limitify/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request"})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user data"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"error": "User registered successfully "})

}

func Login(c *gin.Context) {
	var body models.User
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request"})
		return
	}

	var user models.User

	if err := config.DB.Where("email = ?", body.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Email or Password"})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Password"})
		return
	}

	token, _ := utils.GenerateJWT(user.Email, user.Plan)
	c.JSON(http.StatusOK, gin.H{"token": token})

}
