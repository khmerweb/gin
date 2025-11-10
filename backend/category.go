package backend

import (
	"gin/db"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func RegisterRoutesCategory(router *gin.RouterGroup) {

	router.GET("/", func(c *gin.Context) {
		userName, _ := c.Get("userName")
		session := sessions.Default(c)
		successFlashes := session.Flashes("success")
		errorFlashes := session.Flashes("error")
		session.Save()
		count := db.CountCategories()
		dashboard := Setup().Dashboard
		pageNumbers := make([]int, 0)
		pageCount := (count + dashboard - 1) / dashboard
		for i := 0; i < pageCount; i++ {
			pageNumbers = append(pageNumbers, i+1)
		}
		categories := db.GetCategories(dashboard)
		c.HTML(200, "category", gin.H{
			"Title":           "ទំព័រ​ជំពូក",
			"UserName":        userName,
			"SuccessMessages": successFlashes,
			"ErrorMessages":   errorFlashes,
			"Route":           "ជំពូក",
			"ItemsCount":      count,
			"Items":           categories,
			"Type":            "category",
			"PageNumbers":     pageNumbers,
			"PageNumber":      1,
		})
	})

	router.POST("/", func(c *gin.Context) {
		db.CreateCategory(c)
		c.Redirect(302, "/admin/category")
	})

	router.GET("/edit/:id", func(c *gin.Context) {
		userName, _ := c.Get("userName")
		session := sessions.Default(c)
		successFlashes := session.Flashes("success")
		errorFlashes := session.Flashes("error")
		session.Save()
		count := db.CountCategories()
		category := db.GetCategory(c.Param("id"))
		dashboard := Setup().Dashboard
		pageNumbers := make([]int, 0)
		pageCount := (count + dashboard - 1) / dashboard
		for i := 0; i < pageCount; i++ {
			pageNumbers = append(pageNumbers, i+1)
		}
		page, _ := c.GetQuery("p")
		pageInt, _ := strconv.Atoi(page)
		categories := db.PaginateCategories(c, dashboard, pageInt)

		c.HTML(200, "category-edit", gin.H{
			"Title":           "ទំព័រ​កែប្រែជំពូក",
			"UserName":        userName,
			"Category":        category,
			"SuccessMessages": successFlashes,
			"ErrorMessages":   errorFlashes,
			"Route":           "ជំពូក",
			"ItemsCount":      count,
			"Items":           categories,
			"Type":            "category",
			"PageNumbers":     pageNumbers,
			"PageNumber":      pageInt,
		})
	})

	router.POST("/edit/:id", func(c *gin.Context) {
		db.UpdateCategory(c)
		page, _ := c.GetQuery("p")
		c.Redirect(302, "/admin/category/edit/"+c.Param("id")+"?p="+page)
	})

	router.GET("/delete/:id", func(c *gin.Context) {
		db.DeleteCategory(c)
		c.Redirect(302, "/admin/category")
	})

	router.GET("/paginate/:page", func(c *gin.Context) {
		dashboard := Setup().Dashboard
		categories := db.PaginateCategories(c, dashboard, 0)
		c.JSON(200, gin.H{"items": categories})
	})

}
