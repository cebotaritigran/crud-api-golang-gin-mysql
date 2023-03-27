package main

import (
	"crud-api/models"
	routes "crud-api/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	cors "github.com/rs/cors/wrapper/gin"
	"github.com/unrolled/secure"
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

	/**/
	//SECURE PACKAGE
	secureMiddleware := secure.New(secure.Options{
		SSLRedirect:          false, // set it to false because postman doesn't allow https on localhost
		ForceSTSHeader:       true,
		STSSeconds:           31536000,
		STSIncludeSubdomains: true,
		STSPreload:           true,
		FrameDeny:            true,
	})
	secureFunc := func() gin.HandlerFunc {
		return func(c *gin.Context) {
			err := secureMiddleware.Process(c.Writer, c.Request)

			// If there was an error, do not continue.
			if err != nil {
				c.Abort()
				return
			}

			// Avoid header rewrite if response is a redirection.
			if status := c.Writer.Status(); status > 300 && status < 399 {
				c.Abort()
			}
		}
	}()
	/**/

	router := gin.New()

	/**/
	//SECURE PACKAGE
	router.Use(secureFunc)
	/**/

	// logger will output logs about any successful or unsuccessful callls
	router.Use(gin.Logger())
	// recovery is to handle any panic with 500  (not necesarry)
	router.Use(gin.Recovery())

	router.Use(cors.Default()) // TO BE REMOVED

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-1"})
	})

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-2"})
	})

	router.Run(":" + port)
}
