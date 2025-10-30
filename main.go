// package main
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init() {
	gin.SetMode(gin.ReleaseMode)
	var router = gin.Default()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/**/*.html")
	// Define your Gin routes here
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", gin.H{
			"title":   "My Gin Website",
			"message": "Welcome to the homepage!",
		})
	})
}

func Handler(w http.ResponseWriter, r *http.Request) {
	var router = gin.Default()
	Init()
	router.ServeHTTP(w, r)
}

func main() {
	var router = gin.Default()
	Init()
	router.Run(":8000")
}
