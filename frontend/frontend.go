// frontend/frontend.go
package frontend

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "home", gin.H{
			"title":   "My Gin Website",
			"message": "Welcome to the homepage!",
		})
	})
	router.GET("/post/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Get specific post"})
	})
}
