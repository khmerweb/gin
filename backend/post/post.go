// backend/post/post.go
package post

import (
	"gin/db"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Post struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Categories string `json:"categories"`
	Thumb      string `json:"thumb"`
	Date       string `json:"date"`
	Videos     string `json:"videos"`
	Author     string `json:"author"`
	CreatedAt  string `json:"created_at"`
}

func RegisterRoutes(router *gin.RouterGroup) {
	mydb := db.Connect()
	router.GET("/", func(c *gin.Context) {

		c.Redirect(302, "/admin")
	})

	router.POST("/", func(c *gin.Context) {
		id := uuid.New()
		title := c.PostForm("title")
		content := c.PostForm("content")
		categories := c.PostForm("categories")
		thumb := c.PostForm("thumb")
		date := c.PostForm("date")
		videos := c.PostForm("videos")
		userId, _ := c.Get("userId")
		author := userId.(string)
		created_at := time.Now().Format("2006-01-02 15:04:05")

		sql := `INSERT INTO Post VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
		mydb.Exec(sql, id, title, content, categories, thumb, date, videos, author, created_at)
		defer mydb.Close()
		c.Redirect(302, "/admin")
	})
}
