package api

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Get specific post"})
	})
}
