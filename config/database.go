package config

import (
	"fmt"
	"log"
	"os"
	"limitify/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var DB *gorm.DB


func ConnectDB(){
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
	os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
	os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	
	var err error
	DB , err = gorm.Open(postgres.Open(dsn) , &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database : " , err)
	}

	fmt.Println("Connected to DB ")
}


func MigrateDatabase(){
	DB.AutoMigrate(&models.User{} , &models.RateLimit{} , &models.RequestLog{})

	fmt.Println("Database migration completed")

	var count int64
	DB.Model(&models.RateLimit{}).Count(&count)
	if count == 0 {
		DB.Create(&models.RateLimit{Requests:  100 , TimeWindow:  60})
	}
}

