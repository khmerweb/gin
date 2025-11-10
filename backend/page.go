package backend

import (
	"gin/db"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func RegisterRoutesPage(router *gin.RouterGroup) {

	router.GET("/", func(c *gin.Context) {
		userName, _ := c.Get("userName")
		session := sessions.Default(c)
		successFlashes := session.Flashes("success")
		errorFlashes := session.Flashes("error")
		session.Save()

		count := db.CountPages()
		dashboard := Setup().Dashboard
		pageNumbers := make([]int, 0)
		pageCount := (count + dashboard - 1) / dashboard
		for i := 0; i < pageCount; i++ {
			pageNumbers = append(pageNumbers, i+1)
		}
		pages := db.GetPages(dashboard)
		c.HTML(200, "page", gin.H{
			"Title":           "ទំព័រ​ស្តាទិក",
			"UserName":        userName,
			"SuccessMessages": successFlashes,
			"ErrorMessages":   errorFlashes,
			"Route":           "ទំព័រ​ស្តាទិក",
			"ItemsCount":      count,
			"Items":           pages,
			"Type":            "page",
			"PageNumbers":     pageNumbers,
			"PageNumber":      1,
		})
	})

	router.POST("/", func(c *gin.Context) {
		db.CreatePage(c)
		c.Redirect(302, "/admin/page")
	})

	router.GET("/edit/:id", func(c *gin.Context) {
		userName, _ := c.Get("userName")
		session := sessions.Default(c)
		successFlashes := session.Flashes("success")
		errorFlashes := session.Flashes("error")
		session.Save()
		count := db.CountPages()
		page := db.GetPage(c.Param("id"))
		dashboard := Setup().Dashboard
		pageNumbers := make([]int, 0)
		pageCount := (count + dashboard - 1) / dashboard
		for i := 0; i < pageCount; i++ {
			pageNumbers = append(pageNumbers, i+1)
		}
		pageStr, _ := c.GetQuery("p")
		pageInt, _ := strconv.Atoi(pageStr)
		pages := db.PaginatePages(c, dashboard, pageInt)

		c.HTML(200, "page-edit", gin.H{
			"Title":           "កែប្រែទំព័រ​ស្តាទិក",
			"UserName":        userName,
			"Page":            page,
			"SuccessMessages": successFlashes,
			"ErrorMessages":   errorFlashes,
			"Route":           "ទំព័រ​ស្តាទិក",
			"ItemsCount":      count,
			"Items":           pages,
			"Type":            "page",
			"PageNumbers":     pageNumbers,
			"PageNumber":      pageInt,
		})
	})

	router.POST("/edit/:id", func(c *gin.Context) {
		db.UpdatePage(c)
		page, _ := c.GetQuery("p")
		c.Redirect(302, "/admin/page/edit/"+c.Param("id")+"?p="+page)
	})

	router.GET("/paginate/:page", func(c *gin.Context) {
		dashboard := Setup().Dashboard
		pages := db.PaginatePages(c, dashboard, 0)
		c.JSON(200, gin.H{"items": pages})
	})

}
