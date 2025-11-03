//package main

package handler

import (
	"gin/api"
	"gin/backend"
	"gin/frontend"
	"gin/login"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func createMyRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("home", "templates/layouts/base.html", "templates/pages/home.html", "templates/partials/footer.html")
	r.AddFromFiles("admin", "templates/layouts/baseAdmin.html", "templates/pages/admin.html", "templates/partials/headerAdmin.html", "templates/partials/footer.html")
	r.AddFromFiles("login", "templates/pages/login.html")
	return r
}

func Handler(w http.ResponseWriter, r *http.Request) {
	gin.SetMode(gin.ReleaseMode)
	var router = gin.Default()
	router.Static("/static", "./public/static")
	router.HTMLRender = createMyRender()
	store := cookie.NewStore([]byte(os.Getenv("SECRET_KEY")))
	store.Options(sessions.Options{MaxAge: 0, Path: "/"})
	router.Use(sessions.Sessions("mysession", store))
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://khmertube.vercel.app/"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	front := router.Group("/")
	frontGroup := front.Group("/")
	frontend.RegisterRoutes(frontGroup)

	loginRoute := router.Group("/")
	loginGroup := loginRoute.Group("/login")
	login.RegisterRoutes(loginGroup)

	back := router.Group("/admin")
	backGroup := back.Group("/")
	backend.RegisterRoutes(backGroup)

	apiRoute := router.Group("/api")
	apiGroup := apiRoute.Group("/")
	api.RegisterRoutes(apiGroup)

	router.ServeHTTP(w, r)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	gin.SetMode(gin.ReleaseMode)
	var router = gin.Default()
	router.Static("/static", "./public/static")
	store := cookie.NewStore([]byte(os.Getenv("SECRET_KEY")))
	store.Options(sessions.Options{MaxAge: 0, Path: "/"})
	router.Use(sessions.Sessions("mysession", store))
	router.HTMLRender = createMyRender()
	router.Use(cors.Default())

	front := router.Group("/")
	frontGroup := front.Group("/")
	frontend.RegisterRoutes(frontGroup)

	loginRoute := router.Group("/")
	loginGroup := loginRoute.Group("/login")
	login.RegisterRoutes(loginGroup)

	back := router.Group("/admin")
	backGroup := back.Group("/")
	backend.RegisterRoutes(backGroup)

	apiRoute := router.Group("/api")
	apiGroup := apiRoute.Group("/")
	api.RegisterRoutes(apiGroup)

	router.Run(":8000")
}
