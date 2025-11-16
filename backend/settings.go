package backend

import (
	"gin/db"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Setup() db.Setting {
	var settings db.Setting
	settings = db.Setting{
		Description: "description",
		Dashboard:   10,
		Frontend:    20,
		Categories:  20,
		Playlist:    20,
		Thumb:       "/static/images/no-image.png",
		Date:        "2006-01-02T15:04:05",
	}

	mySettings := db.GetSettings(settings.Dashboard)
	if len(mySettings) > 0 {
		settings = mySettings[0]
		return settings
	}

	return settings
}

func RegisterRoutesSetting(router *gin.RouterGroup) {
	router.GET("/", func(c *gin.Context) {
		userName, _ := c.Get("userName")
		session := sessions.Default(c)
		successFlashes := session.Flashes("success")
		errorFlashes := session.Flashes("error")
		session.Save()
		count := db.CountSettings()
		pageNumbers := make([]int, 0)
		dashboard := Setup().Dashboard
		pageCount := (count + dashboard - 1) / dashboard
		for i := 0; i < pageCount; i++ {
			pageNumbers = append(pageNumbers, i+1)
		}
		settings := db.GetSettings(dashboard)
		c.HTML(200, "setting", gin.H{
			"Title":           "ទំព័រ​ Setting",
			"UserName":        userName,
			"SuccessMessages": successFlashes,
			"ErrorMessages":   errorFlashes,
			"Route":           "setting ",
			"ItemsCount":      count,
			"Items":           settings,
			"Type":            "setting",
			"PageNumbers":     pageNumbers,
			"PageNumber":      1,
		})
	})

	router.POST("/", func(c *gin.Context) {
		db.CreateSetting(c)
		c.Redirect(302, "/admin/setting")
	})

	router.GET("/paginate/:page", func(c *gin.Context) {
		dashboard := Setup().Dashboard
		settings := db.PaginateSettings(c, dashboard, 0)
		c.JSON(200, gin.H{"items": settings})
	})

	router.GET("/delete/:id", func(c *gin.Context) {
		postID := c.Param("id")
		db.DeleteSetting(postID, c)
		c.Redirect(302, "/admin/setting")
	})

	router.GET("/edit/:id", func(c *gin.Context) {
		userName, _ := c.Get("userName")
		session := sessions.Default(c)
		successFlashes := session.Flashes("success")
		errorFlashes := session.Flashes("error")
		session.Save()
		count := db.CountSettings()
		setting := db.GetSetting(c.Param("id"))
		dashboard := Setup().Dashboard
		pageNumbers := make([]int, 0)
		pageCount := (count + dashboard - 1) / dashboard
		for i := 0; i < pageCount; i++ {
			pageNumbers = append(pageNumbers, i+1)
		}
		pageStr, _ := c.GetQuery("p")
		pageInt, _ := strconv.Atoi(pageStr)
		settings := db.PaginateSettings(c, dashboard, pageInt)

		c.HTML(200, "setting-edit", gin.H{
			"Title":           "កែប្រែ setting",
			"UserName":        userName,
			"Setting":         setting,
			"SuccessMessages": successFlashes,
			"ErrorMessages":   errorFlashes,
			"Route":           "setting",
			"ItemsCount":      count,
			"Items":           settings,
			"Type":            "setting",
			"PageNumbers":     pageNumbers,
			"PageNumber":      pageInt,
		})
	})

	router.POST("/edit/:id", func(c *gin.Context) {
		db.UpdateSetting(c)
		page, _ := c.GetQuery("p")
		c.Redirect(302, "/admin/setting/edit/"+c.Param("id")+"?p="+page)
	})

}
