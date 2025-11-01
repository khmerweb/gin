// login/login.go
package login

import (
	"fmt"
	"gin/db"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       string
	Title    string
	Email    string
	Password string
	Role     string
}

func RegisterRoutes(router *gin.RouterGroup) {
	//db.CreateRootUser()
	mydb := db.Connect()
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "login", gin.H{
			"title":   "Login Page",
			"message": "Welcome to the login page!",
		})
	})
	router.POST("/", func(c *gin.Context) {
		email := c.PostForm("email")
		password := c.PostForm("password")

		sql := `SELECT id, title, email, password, role FROM User WHERE email = ?`
		row := mydb.QueryRow(sql, email)
		defer mydb.Close()

		u := &User{}
		err := row.Scan(&u.Id, &u.Title, &u.Email, &u.Password, &u.Role)
		if err != nil {
			return
		} else {
			err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
			if err != nil {
				if err == bcrypt.ErrMismatchedHashAndPassword {
					fmt.Println("Invalid password.")
				} else {
					fmt.Println("Error comparing passwords:", err)
				}
				return
			}

			c.Redirect(http.StatusFound, "/admin/")
		}
	})
}
