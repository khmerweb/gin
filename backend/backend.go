// backend/backend.go
package backend

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "admin", gin.H{
			"title":   "Admin Page",
			"message": "Welcome to the admin page!",
		})
	})
	router.GET("/logout", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Clear()
		session.Options(sessions.Options{Path: "/", MaxAge: -1})
		session.Save()
		c.Redirect(302, "/login")
	})
}
