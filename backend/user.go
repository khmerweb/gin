package backend

import (
	"gin/db"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func RegisterRoutesUser(router *gin.RouterGroup) {
	router.GET("/", func(c *gin.Context) {
		userName, _ := c.Get("userName")
		session := sessions.Default(c)
		successFlashes := session.Flashes("success")
		errorFlashes := session.Flashes("error")
		session.Save()
		count := db.CountUsers()
		pageNumbers := make([]int, 0)
		dashboard := Setup().Dashboard
		pageCount := (count + dashboard - 1) / dashboard
		for i := 0; i < pageCount; i++ {
			pageNumbers = append(pageNumbers, i+1)
		}
		users := db.GetUsers(dashboard)
		c.HTML(200, "user", gin.H{
			"Title":           "ទំព័រ​អ្នក​និពន្ធ",
			"UserName":        userName,
			"SuccessMessages": successFlashes,
			"ErrorMessages":   errorFlashes,
			"Route":           "អ្នក​និពន្ធ",
			"ItemsCount":      count,
			"Items":           users,
			"Type":            "user",
			"PageNumbers":     pageNumbers,
			"PageNumber":      1,
		})
	})

	router.POST("/", func(c *gin.Context) {
		db.CreateUser(c)
		c.Redirect(302, "/admin/user")
	})

	router.GET("/paginate/:page", func(c *gin.Context) {
		dashboard := Setup().Dashboard
		users := db.PaginateUsers(c, dashboard, 0)
		c.JSON(200, gin.H{"items": users})
	})

	router.GET("/delete/:id", func(c *gin.Context) {
		postID := c.Param("id")
		db.DeleteUser(postID, c)
		c.Redirect(302, "/admin/user")
	})

	router.GET("/edit/:id", func(c *gin.Context) {
		userName, _ := c.Get("userName")
		session := sessions.Default(c)
		successFlashes := session.Flashes("success")
		errorFlashes := session.Flashes("error")
		session.Save()
		count := db.CountUsers()
		user := db.GetUser(c.Param("id"))

		options := []string{"Author", "Editor", "Admin"}
		selected := user.Role
		dashboard := Setup().Dashboard
		pageNumbers := make([]int, 0)
		pageCount := (count + dashboard - 1) / dashboard
		for i := 0; i < pageCount; i++ {
			pageNumbers = append(pageNumbers, i+1)
		}
		pageStr, _ := c.GetQuery("p")
		pageInt, _ := strconv.Atoi(pageStr)
		users := db.PaginateUsers(c, dashboard, pageInt)

		c.HTML(200, "user-edit", gin.H{
			"Title":           "កែប្រែអ្នក​​និពន្ធ",
			"UserName":        userName,
			"User":            user,
			"SuccessMessages": successFlashes,
			"ErrorMessages":   errorFlashes,
			"Route":           "អ្នក​និពន្ធ",
			"Options":         options,
			"Selected":        selected,
			"ItemsCount":      count,
			"Items":           users,
			"Type":            "user",
			"PageNumbers":     pageNumbers,
			"PageNumber":      pageInt,
		})
	})

	router.POST("/edit/:id", func(c *gin.Context) {
		db.UpdateUser(c)
		page, _ := c.GetQuery("p")
		c.Redirect(302, "/admin/user/edit/"+c.Param("id")+"?p="+page)
	})

}
