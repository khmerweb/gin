// login/login.go
package login

import (
	"gin/db"
	"log"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id    string
	Title string
	Email string
	Role  string
}

func RegisterRoutes(router *gin.RouterGroup) {
	//db.CreateRootUser()
	mydb := db.Connect()
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "login.html", gin.H{
			"title":   "Login Page",
			"message": "Welcome to the login page!",
		})
	})
	router.POST("/", func(c *gin.Context) {
		email := c.PostForm("email")
		//password := c.PostForm("password")

		sql := `SELECT id, title, email, role FROM User WHERE email = ?`
		row := mydb.QueryRow(sql, email)

		u := &User{}
		err := row.Scan(&u.Id, &u.Title, &u.Email, &u.Role)
		if err != nil {
			defer mydb.Close()
		}

		log.Println(u.Title)
	})
}
