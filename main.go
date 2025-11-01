//package main

package handler

import (
	"gin/backend"
	"gin/frontend"
	"gin/login"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	gin.SetMode(gin.ReleaseMode)
	var router = gin.Default()
	router.Static("/static", "./public/static")
	router.LoadHTMLGlob("templates/**/*.html")
	// Define your Gin routes here
	front := router.Group("/")
	frontGroup := front.Group("/")
	frontend.RegisterRoutes(frontGroup)

	loginRoute := router.Group("/")
	loginGroup := loginRoute.Group("/login")
	login.RegisterRoutes(loginGroup)

	back := router.Group("/")
	backGroup := back.Group("/admin")
	backend.RegisterRoutes(backGroup)

	router.ServeHTTP(w, r)
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	var router = gin.Default()
	router.Static("/static", "./public/static")
	router.LoadHTMLGlob("templates/**/*.html")

	front := router.Group("/")
	frontGroup := front.Group("/")
	frontend.RegisterRoutes(frontGroup)

	loginRoute := router.Group("/")
	loginGroup := loginRoute.Group("/login")
	login.RegisterRoutes(loginGroup)

	back := router.Group("/")
	backGroup := back.Group("/admin")
	backend.RegisterRoutes(backGroup)

	router.Run(":8000")
}
