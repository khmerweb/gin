// backend/post/post.go
package post

import (
	"gin/db"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Post struct {
	ID         string `json:"id"`
	Title      string `json:"title" form:"title" binding:"required"`
	Content    string `json:"content" form:"content"`
	Categories string `json:"categories" form:"categories" binding:"required"`
	Thumb      string `json:"thumb" form:"thumb" binding:"required"`
	Date       string `json:"date" form:"date" binding:"required"`
	Videos     string `json:"videos" form:"videos"`
	Author     string `json:"author"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

func RegisterRoutes(router *gin.RouterGroup) {
	//db.CreateSchema()
	mydb := db.Connect()
	router.GET("/", func(c *gin.Context) {

		c.Redirect(302, "/admin")
	})

	router.POST("/", func(c *gin.Context) {
		session := sessions.Default(c)
		var post Post
		if err := c.ShouldBind(&post); err != nil {
			session.AddFlash("Unable to create post!", "error")
			session.Save()
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		id := uuid.New()
		title := c.PostForm("title")
		content := c.PostForm("content")
		categories := c.PostForm("categories")
		thumb := c.PostForm("thumb")
		date := c.PostForm("date")
		videos := c.PostForm("videos")
		userId, _ := c.Get("userId")
		author := userId.(string)
		created_at := time.Now().Format("2006-01-02T15:04:05")
		updated_at := time.Now().Format("2006-01-02T15:04:05")

		sql := `INSERT INTO Post VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
		mydb.Exec(sql, id, title, content, categories, thumb, date, videos, author, created_at, updated_at)
		defer mydb.Close()
		session.AddFlash("Post created successfully!", "success")
		session.Save()
		c.Redirect(302, "/admin")
	})
}
