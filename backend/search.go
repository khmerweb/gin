package backend

import (
	"gin/db"

	"github.com/gin-gonic/gin"
)

func RegisterRoutesSearch(router *gin.RouterGroup) {
	router.POST("/", func(c *gin.Context) {
		userName, _ := c.Get("userName")
		q := c.PostForm("q")
		searchPosts := db.SearchPosts(q, 10)
		count := db.CountPosts()
		pageNumbers := make([]int, 0)
		dashboard := Setup().Dashboard
		pageCount := (count + dashboard - 1) / dashboard
		for i := 0; i < pageCount; i++ {
			pageNumbers = append(pageNumbers, i+1)
		}
		posts := db.GetPosts(dashboard)
		c.HTML(200, "search", gin.H{
			"Title":       "ទំព័រ​ការផ្សាយ",
			"UserName":    userName,
			"SearchPosts": searchPosts,
			"Route":       "ការផ្សាយ",
			"ItemsCount":  count,
			"Items":       posts,
			"Type":        "post",
			"PageNumbers": pageNumbers,
			"PageNumber":  1,
		})
	})

}
