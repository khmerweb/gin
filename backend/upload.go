package backend

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutesUpload(router *gin.RouterGroup) {
	router.GET("/", func(c *gin.Context) {
		userName, _ := c.Get("userName")

		count := 1
		pageNumbers := make([]int, 0)
		dashboard := Setup().Dashboard
		pageCount := (count + dashboard - 1) / dashboard
		for i := 0; i < pageCount; i++ {
			pageNumbers = append(pageNumbers, i+1)
		}
		c.HTML(200, "upload", gin.H{
			"Title":       "ទំព័រ​ Upload",
			"UserName":    userName,
			"PageNumbers": pageNumbers,
			"PageNumber":  1,
		})
	})
}
