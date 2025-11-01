//package main

package handler

import (
	"gin/backend"
	"gin/frontend"
	"gin/login"
	"log"
	"net/http"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func createMyRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("home", "templates/layouts/base.html", "templates/pages/home.html")
	r.AddFromFiles("admin", "templates/layouts/base.html", "templates/pages/admin.html")
	r.AddFromFiles("login", "templates/pages/login.html")
	return r
}

func Handler(w http.ResponseWriter, r *http.Request) {
	gin.SetMode(gin.ReleaseMode)
	var router = gin.Default()
	router.Static("/static", "./public/static")
	router.HTMLRender = createMyRender()
	//router.LoadHTMLGlob("templates/**/*.html")
	// Define your Gin routes here
	front := router.Group("/")
	frontGroup := front.Group("/")
	frontend.RegisterRoutes(frontGroup)

	loginRoute := router.Group("/")
	loginGroup := loginRoute.Group("/login")
	login.RegisterRoutes(loginGroup)

	back := router.Group("/admin")
	backGroup := back.Group("/")
	backend.RegisterRoutes(backGroup)

	router.ServeHTTP(w, r)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	gin.SetMode(gin.DebugMode)
	var router = gin.Default()
	router.Static("/static", "./public/static")
	//router.LoadHTMLGlob("templates/**/*.html")
	router.HTMLRender = createMyRender()

	front := router.Group("/")
	frontGroup := front.Group("/")
	frontend.RegisterRoutes(frontGroup)

	loginRoute := router.Group("/")
	loginGroup := loginRoute.Group("/login")
	login.RegisterRoutes(loginGroup)

	back := router.Group("/admin")
	backGroup := back.Group("/")
	backend.RegisterRoutes(backGroup)

	router.Run(":8000")
}
