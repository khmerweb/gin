// backend/post/post.go
package post

import (
	"gin/db"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {

	router.GET("/", func(c *gin.Context) {
		userName, _ := c.Get("userName")
		session := sessions.Default(c)
		successFlashes := session.Flashes("success")
		errorFlashes := session.Flashes("error")
		session.Save()
		count := db.CountPosts()
		c.HTML(200, "admin", gin.H{
			"title":           "Post Page",
			"userName":        userName,
			"SuccessMessages": successFlashes,
			"ErrorMessages":   errorFlashes,
			"PostCount":       count,
			"Items":           "ការផ្សាយ",
		})
	})

	router.POST("/", func(c *gin.Context) {
		db.CreatePost(c)
		c.Redirect(302, "/admin/post")
	})
}
