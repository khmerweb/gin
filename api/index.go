package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/**/*.html")
	// Define your Gin routes here//
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", gin.H{
			"title":   "My Gin Website",
			"message": "Welcome to the homepage!",
		})
	})

	// Add more routes as needed
	// router.POST("/api/data", yourHandlerFunction)

	router.ServeHTTP(w, r)
}
