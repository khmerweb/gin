//package main

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func Init() {
	gin.SetMode(gin.ReleaseMode)
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
	Init()
	router.ServeHTTP(w, r)
}

func main() {
	Init()
	router.Run(":8000")
}
