// login/login.go
package login

import (
	"fmt"
	"gin/db"
	"net/http"

	"github.com/gin-contrib/sessions"
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
		session := sessions.Default(c)
		//session.Clear()
		//session.Options(sessions.Options{Path: "/", MaxAge: -1})
		//session.Save()
		value := session.Get("userId")

		var message string
		if value != nil {
			userId, _ := value.(string)
			if userId == "noUser" {
				message = "Email ឬ ​ពាក្យ​សំងាត់​មិន​ត្រូវ​ទេ!"
			} else {
				c.Redirect(http.StatusFound, "/admin")
			}
		}
		c.HTML(200, "login", gin.H{
			"title":   "Login Page",
			"message": message,
		})
	})
	router.POST("/", func(c *gin.Context) {
		session := sessions.Default(c)
		email := c.PostForm("email")
		password := c.PostForm("password")

		mysql := `SELECT id, title, email, password, role FROM User WHERE email = ?`
		row := mydb.QueryRow(mysql, email)
		defer mydb.Close()

		u := &User{}
		err := row.Scan(&u.Id, &u.Title, &u.Email, &u.Password, &u.Role)
		if err != nil {
			session.Set("userId", "noUser")
			session.Save()
			c.Redirect(http.StatusFound, "/login")
		} else {
			err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
			if err != nil {
				if err == bcrypt.ErrMismatchedHashAndPassword {
					session.Set("userId", "noUser")
					session.Save()
					c.Redirect(http.StatusFound, "/login")
				} else {
					fmt.Println("Error comparing passwords:", err)
					c.Redirect(http.StatusFound, "/login")
				}
				return
			}

			session.Set("userId", u.Id)
			session.Set("userName", u.Title)
			session.Set("userRole", u.Role)
			session.Save()
			c.Redirect(http.StatusFound, "/admin")
		}
	})
}
