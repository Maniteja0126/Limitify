package main

import (
	"github.com/gin-gonic/gin"
	"limitify/config"
	"limitify/routes"
	"github.com/joho/godotenv"
	"log"
	"os"
	"fmt"
)


func main(){

	err := godotenv.Load()
	if err != nil {
		log.Println("Warning : No .env found")
	}
	config.PrintConfig()


	config.ConnectDB()
	config.MigrateDatabase()
	config.InitRedis()

	r := gin.Default()

	routes.SetUpRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fmt.Println("Server running on port " , port)
	r.Run(":" + port)

}
