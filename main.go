package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Gin router with default middleware (Logger and Recovery)
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/**/*.html")

	router.GET("/", func(c *gin.Context) {
		// Render the "index.html" template and pass data
		c.HTML(http.StatusOK, "home.html", gin.H{
			"title":   "My Gin Website",
			"message": "Welcome to the homepage!",
		})
	})

	router.Run(":8000")
}
