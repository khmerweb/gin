// backend/backend.go
// go mod tidy
package backend

import (
	"gin/backend/post"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userName := session.Get("userName")

		if userName == nil {
			c.Redirect(302, "/login")
			return
		}

		c.Set("userName", userName)
		c.Set("userId", session.Get("userId"))
		c.Next()
	}
}

func RegisterRoutes(router *gin.RouterGroup) {
	router.Use(AuthRequired())

	router.GET("/", func(c *gin.Context) {
		userName, _ := c.Get("userName")
		session := sessions.Default(c)
		successFlashes := session.Flashes("success")
		errorFlashes := session.Flashes("error")
		session.Save()

		c.HTML(200, "admin", gin.H{
			"title":           "Admin Page",
			"userName":        userName,
			"SuccessMessages": successFlashes,
			"ErrorMessages":   errorFlashes,
		})

	})

	router.GET("/logout", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Clear()
		session.Options(sessions.Options{Path: "/", MaxAge: -1})
		session.Save()
		c.Redirect(302, "/login")
	})

	postRoute := router.Group("/post")
	postGroup := postRoute.Group("/")
	post.RegisterRoutes(postGroup)

}
