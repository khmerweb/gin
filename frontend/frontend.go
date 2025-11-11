// frontend/frontend.go
package frontend

import (
	"gin/backend"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/", func(c *gin.Context) {
		siteTitle := backend.Setup().SiteTitle
		c.HTML(200, "home", gin.H{
			"Title":     "My Gin Website",
			"SiteTitle": siteTitle,
			"message":   "Welcome to the homepage!",
		})
	})
	router.GET("/post/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Get specific post"})
	})
}
