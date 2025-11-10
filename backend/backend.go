// backend/backend.go
// go mod tidy
package backend

import (
	"gin/settings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userName := session.Get("userName")

		if userName == nil {
			c.Redirect(302, "/login")
			return
		}

		c.Set("userName", userName)
		c.Set("userId", session.Get("userId"))
		c.Set("userRole", session.Get("userRole"))
		c.Next()
	}
}

func RegisterRoutes(router *gin.RouterGroup) {
	router.Use(AuthRequired())

	router.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/admin/post")
	})

	router.GET("/logout", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Clear()
		session.Options(sessions.Options{Path: "/", MaxAge: -1})
		session.Save()
		c.Redirect(302, "/login")
	})

	postRoute := router.Group("/post")
	postGroup := postRoute.Group("/")
	RegisterRoutesPost(postGroup)

	categoryRoute := router.Group("/category")
	categoryGroup := categoryRoute.Group("/")
	RegisterRoutesCategory(categoryGroup)

	pageRoute := router.Group("/page")
	pageGroup := pageRoute.Group("/")
	RegisterRoutesPage(pageGroup)

	uploadRoute := router.Group("/upload")
	uploadGroup := uploadRoute.Group("/")
	RegisterRoutesUpload(uploadGroup)

	userRoute := router.Group("/user")
	userGroup := userRoute.Group("/")
	RegisterRoutesUser(userGroup)

	settingRoute := router.Group("/setting")
	settingGroup := settingRoute.Group("/")
	settings.RegisterRoutesSetting(settingGroup)

}
