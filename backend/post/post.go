// backend/post/post.go
package post

import (
	"gin/db"
	"gin/settings"
	"strconv"

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
		pageCount := (count + 10 - 1) / 10
		for i := 0; i < pageCount; i++ {
			pageNumbers = append(pageNumbers, i+1)
		}
		dashboard := settings.Setup().Dashboard
		posts := db.GetPosts(dashboard)
		c.HTML(200, "admin", gin.H{
			"Title":           "ទំព័រ​ការផ្សាយ",
			"UserName":        userName,
			"SuccessMessages": successFlashes,
			"ErrorMessages":   errorFlashes,
			"Route":           "ការផ្សាយ",
			"PostCount":       count,
			"Items":           posts,
			"Type":            "post",
			"PageNumbers":     pageNumbers,
			"PageNumber":      1,
		})
	})

	router.POST("/", func(c *gin.Context) {
		db.CreatePost(c)
		c.Redirect(302, "/admin/post")
	})

	router.GET("/delete/:id", func(c *gin.Context) {
		postID := c.Param("id")
		db.DeletePost(postID, c)
		c.Redirect(302, "/admin/post")
	})

	router.GET("/edit/:id", func(c *gin.Context) {
		userName, _ := c.Get("userName")
		session := sessions.Default(c)
		successFlashes := session.Flashes("success")
		errorFlashes := session.Flashes("error")
		session.Save()
		count := db.CountPosts()
		post := db.GetPost(c.Param("id"))
		pageNumbers := make([]int, 0)
		pageCount := (count + 10 - 1) / 10
		for i := 0; i < pageCount; i++ {
			pageNumbers = append(pageNumbers, i+1)
		}
		dashboard := settings.Setup().Dashboard
		page, _ := c.GetQuery("p")
		pageInt, _ := strconv.Atoi(page)
		posts := db.PaginatePosts(c, dashboard, pageInt)

		c.HTML(200, "admin-edit", gin.H{
			"Title":           "ទំព័រ​កែប្រែការផ្សាយ",
			"UserName":        userName,
			"Post":            post,
			"SuccessMessages": successFlashes,
			"ErrorMessages":   errorFlashes,
			"Route":           "ការផ្សាយ",
			"PostCount":       count,
			"Items":           posts,
			"Type":            "post",
			"PageNumbers":     pageNumbers,
			"PageNumber":      pageInt,
		})
	})

	router.POST("/edit/:id", func(c *gin.Context) {
		db.UpdatePost(c)
		page, _ := c.GetQuery("p")
		c.Redirect(302, "/admin/post/edit/"+c.Param("id")+"?p="+page)
	})

	router.GET("/paginate/:page", func(c *gin.Context) {
		dashboard := settings.Setup().Dashboard
		posts := db.PaginatePosts(c, dashboard, 0)
		c.JSON(200, gin.H{"items": posts})
	})
}
