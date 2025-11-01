// backend/backend.go
package backend

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "admin", gin.H{
			"title":   "Welcome to the admin page!",
			"message": "Welcome to the admin page!",
		})
	})
	router.GET("/post/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Get specific post"})
	})
}
