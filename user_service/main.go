package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


func main() {
	router := gin.Default()

	// Configure CORS
	config := cors.Config{
        AllowOrigins:     []string{"http://localhost:5173", "http://152.59.146.250"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},     // Allow specific headers
        ExposeHeaders:    []string{"Content-Length"},                       // Expose specific headers
        AllowCredentials: true,                                            // Allow credentials
        MaxAge:           12 * time.Hour,                                   // Max age for preflight requests
    }

    // Apply the CORS middleware
    router.Use(cors.New(config))

	// Get all users
	router.GET("/users", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"data": []string{"user data"},
		})
	})	

	router.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"data": []string{"user data"},
		})
	})

	router.Run("0.0.0.0:8080")
}