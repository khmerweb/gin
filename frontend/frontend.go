// frontend/frontend.go
package frontend

import (
	"gin/backend"
	"gin/db"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/", func(c *gin.Context) {
		siteTitle := backend.Setup().SiteTitle
		limit := backend.Setup().Playlist
		playlists := db.GetPlaylists(limit)
		countPlaylists := db.CountPlaylists()
		frontend := backend.Setup().Frontend
		c.HTML(200, "home", gin.H{
			"Title":          "My Gin Website",
			"SiteTitle":      siteTitle,
			"Playlists":      playlists,
			"CountPlaylists": countPlaylists,
			"Frontend":       frontend,
		})
	})

	router.GET("/post/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Get specific post"})
	})
}
