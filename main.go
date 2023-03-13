package main

import (
	"crud-api/models"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env") // loading variables and their values from .env file
	models.ConnectDataBase()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// getting the port from .env file
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	// logger will output logs about any successful or unsuccessful calls
	router.Use(gin.Logger())
	// recovery is to handle any panic with 500  (not necesary)
	router.Use(gin.Recovery())

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-1"})
	})

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-2"})
	})

	router.Run(":" + port)
}
