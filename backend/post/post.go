// backend/post/post.go
package post

import (
	"fmt"
	"gin/db"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
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

	router.GET("/", func(c *gin.Context) {
		userName, _ := c.Get("userName")
		session := sessions.Default(c)
		successFlashes := session.Flashes("success")
		errorFlashes := session.Flashes("error")
		session.Save()
		count := db.CountPosts()
		pageNumbers := make([]int, 0)
		pageCount := (count + 20 - 1) / 20
		for i := 0; i < pageCount; i++ {
			pageNumbers = append(pageNumbers, i+1)
		}
		println("Page Numbers:", pageNumbers)
		mydb := db.Connect()
		post := &Post{}
		mysql := `SELECT * FROM Post ORDER BY date DESC LIMIT ?`
		rows, err := mydb.Query(mysql, 20)
		if err != nil {
			fmt.Println("Error querying database:", err)
			return
		}
		defer rows.Close()
		var posts []Post
		for rows.Next() {
			err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Categories, &post.Thumb, &post.Date, &post.Videos, &post.Author, &post.CreatedAt, &post.UpdatedAt)
			if err != nil {
				fmt.Println("Error scanning row:", err)
				continue
			}
			posts = append(posts, *post)
		}
		c.HTML(200, "admin", gin.H{
			"Title":           "ទំព័រ​ការផ្សាយ",
			"UserName":        userName,
			"SuccessMessages": successFlashes,
			"ErrorMessages":   errorFlashes,
			"Route":           "ការផ្សាយ",
			"PostCount":       count,
			"Items":           posts,
			"Type":            "post",
			"Value":           "post",
			"PageNumbers":     pageNumbers,
			"PageNumber":      1,
		})
	})

	router.POST("/", func(c *gin.Context) {
		db.CreatePost(c)
		c.Redirect(302, "/admin/post")
	})
}
