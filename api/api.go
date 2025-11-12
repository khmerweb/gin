package api

import (
	"gin/backend"
	"gin/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StringArrayRequest struct {
	Thumbs []string `json:"thumbs"` // The 'json:"strings"' tag maps to the JSON field name
}

func RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Get specific post"})
	})

	router.POST("/playlist/:category", func(c *gin.Context) {
		var reqBody StringArrayRequest
		if err := c.ShouldBindJSON(&reqBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		limit := backend.Setup().Playlist
		playlist := db.GetPlaylist(c, limit, reqBody.Thumbs)
		c.JSON(200, gin.H{"playlist": playlist})
	})
}
