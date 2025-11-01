// login/login.go
package login

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	//db.CreateRootUser()
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "login.html", gin.H{
			"title":   "Login Page",
			"message": "Welcome to the login page!",
		})
	})
	router.GET("/post/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Get specific post"})
	})
}
