package db

import (
	"fmt"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email" form:"email" binding:"required"`
	Title    string `json:"title" form:"title" binding:"required"`
	Password string `json:"password" form:"password"`
	Role     string `json:"role" form:"role" binding:"required"`
	Thumb    string `json:"thumb" form:"thumb" binding:"required"`
	Content  string `json:"content" form:"content"`
	Date     string `json:"date" form:"date" binding:"required"`
	Videos   string `json:"videos" form:"videos"`
}

func CreateUserSchema() {
	mydb := Connect()
	defer mydb.Close()
	sql := `CREATE TABLE IF NOT EXISTS User (
        id TEXT PRIMARY KEY,
        title TEXT NOT NULL,
        email TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL,
		role TEXT NOT NULL,
		thumb TEXT,
		content TEXT,
		date TEXT NOT NULL
    );
	`

	_, err := mydb.Exec(sql)
	if err != nil {
		panic(err)
	}
}

func CreateRootUser() {
	mydb := Connect()
	sql := `INSERT INTO User VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	id := uuid.New()
	title := "Sokhavuth"
	email := "sokhavuth@khmerweb.app"
	passwordbyte := []byte("*********")
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordbyte, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	password := string(hashedPassword)
	role := "Admin"
	thumb := "/static/images/no-image.png"
	content := ""
	date := "2006-01-02T15:04:05"
	mydb.Exec(sql, id, title, email, password, role, thumb, content, date)
	defer mydb.Close()
}

func CountUsers() int {
	mydb := Connect()
	defer mydb.Close()
	var count int
	sql := `SELECT COUNT(*) FROM User`
	row := mydb.QueryRow(sql)
	row.Scan(&count)
	return count
}

func GetUsers(limit int) []User {
	mydb := Connect()
	user := &User{}
	mysql := `SELECT * FROM User ORDER BY date DESC LIMIT ?`
	rows, err := mydb.Query(mysql, limit)
	if err != nil {
		fmt.Println("Error querying database:", err)
		return nil
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Title, &user.Email, &user.Password, &user.Role, &user.Thumb, &user.Content, &user.Date)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		users = append(users, *user)
	}
	return users
}

func CreateUser(c *gin.Context) {
	userRole, _ := c.Get("userRole")
	session := sessions.Default(c)
	if userRole != "Admin" {
		session.AddFlash("អ្នក​គ្មាន​សិទ្ធិ​បង្កើតអ្នក​និពន្ធ​ទេ!", "error")
		session.Save()
		return
	}

	mydb := Connect()
	defer mydb.Close()

	user := User{}
	if err := c.ShouldBind(&user); err != nil {
		session.AddFlash("មាន​បញ្ហា​ក្នុង​ការបង្កើត​អ្នក​និពន្អ!", "error")
		session.Save()
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user.ID = uuid.New().String()
	passwordbyte := []byte(user.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordbyte, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	password := string(hashedPassword)

	sql := `INSERT INTO User VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	_, err1 := mydb.Exec(sql, user.ID, user.Title, user.Email, password, user.Role, user.Thumb, user.Content, user.Date)
	if err1 != nil {
		session.AddFlash("មាន​បញ្ហា​ក្នុង​ការបង្កើត​អ្នក​និពន្ធ!", "error")
		session.Save()
		c.JSON(500, gin.H{"error": err1.Error()})
		return
	}
	defer mydb.Close()
	session.AddFlash("អ្នក​និពន្ធ​ត្រូវ​បាន​បង្កើត​ឡើងដោយ​ជោគជ័យ!", "success")
	session.Save()
}

func PaginateUsers(c *gin.Context, dashboard int, query int) []User {
	mydb := Connect()
	defer mydb.Close()
	var page int
	if query != 0 {
		page = query
	} else {
		pageStr := c.Param("page")
		page, _ = strconv.Atoi(pageStr)
	}

	limit := dashboard
	offset := (page - 1) * dashboard
	user := &User{}
	mysql := `SELECT * FROM User ORDER BY date DESC LIMIT ? OFFSET ?`
	rows, err := mydb.Query(mysql, limit, offset)
	if err != nil {
		fmt.Println("Error querying database:", err)
		return nil
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Title, &user.Email, &user.Password, &user.Role, &user.Thumb, &user.Content, &user.Date)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		users = append(users, *user)
	}
	return users
}

func DeleteUser(id string, c *gin.Context) {
	session := sessions.Default(c)
	userRole, _ := c.Get("userRole")
	if userRole != "Admin" {
		session.AddFlash("អ្នក​មិន​មាន​សិទ្ធិ​លុប​អ្នក​និពន្ធ​ទេ!", "error")
		session.Save()
		return
	}

	mydb := Connect()
	defer mydb.Close()
	mysql := `DELETE FROM User WHERE id = ?`
	_, err := mydb.Exec(mysql, id)
	if err != nil {
		session.AddFlash("មាន​បញ្ហា​ក្នុង​ការលុបអ្នក​និពន្ធ!", "error")
		session.Save()
		fmt.Println("Error deleting user:", err)
		return
	}
	session.AddFlash("​អ្នក​និពន្ធ​ត្រូវ​​បានលុប​​ដោយ​ជោគជ័យ!", "success")
	session.Save()
}

func GetUser(id string) User {
	mydb := Connect()
	user := &User{}
	mysql := `SELECT * FROM User WHERE id = ?`
	row := mydb.QueryRow(mysql, id)
	err := row.Scan(&user.ID, &user.Title, &user.Email, &user.Password, &user.Role, &user.Thumb, &user.Content, &user.Date)
	if err != nil {
		fmt.Println("Error scanning row:", err)
		return User{}
	}
	defer mydb.Close()
	return *user
}

func UpdateUser(c *gin.Context) {
	userRole, _ := c.Get("userRole")
	currentUserID, _ := c.Get("userId")
	session := sessions.Default(c)
	mydb := Connect()
	defer mydb.Close()

	var user User
	user.ID = c.Param("id")
	oldUser := GetUser(user.ID)

	if err := c.ShouldBind(&user); err != nil {
		session.AddFlash("ការកែប្រែ​អ្នក​និពន្ធ​​មាន​បញ្ហា!: "+err.Error(), "error")
		session.Save()
		return
	}

	var newPassword string
	if user.Password == oldUser.Password {
		newPassword = oldUser.Password
	} else {
		passwordbyte := []byte(user.Password)
		hashedPassword, err := bcrypt.GenerateFromPassword(passwordbyte, bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}
		password := string(hashedPassword)
		newPassword = password
	}

	if userRole != "Admin" {
		if currentUserID == user.ID {
			sql := `UPDATE User SET title = ?, email = ?, password = ?, thumb = ?, content = ?, date = ? WHERE id = ?`
			_, err := mydb.Exec(sql, user.Title, user.Email, newPassword, user.Thumb, user.Content, user.Date, user.ID)
			if err != nil {
				session.AddFlash("ការកែប្រែ​អ្នក​និពន្ធ​​មាន​បញ្ហា!: "+err.Error(), "error")
				session.Save()
				return
			}
			session.AddFlash("​​អ្នក​និពន្ធត្រូវ​បាន​កែប្រែ​ដោយ​ជោគជ័យ!", "success")
			session.Save()
		} else {
			session.AddFlash("អ្នក​គ្មាន​សិទ្ធិកែប្រែ​អ្នក​និពន្ធផ្សេងបាន​​​​ឡើយ!", "error")
			session.Save()
		}
		return
	}

	sql := `UPDATE User SET title = ?, email = ?, password = ?, role = ?, thumb = ?, content = ?, date = ? WHERE id = ?`
	_, err := mydb.Exec(sql, user.Title, user.Email, newPassword, user.Role, user.Thumb, user.Content, user.Date, user.ID)
	if err != nil {
		session.AddFlash("ការកែប្រែ​អ្នក​និពន្ធ​​មាន​បញ្ហា!: "+err.Error(), "error")
		session.Save()
		return
	}
	session.AddFlash("​​អ្នក​និពន្ធត្រូវ​បាន​កែប្រែ​ដោយ​ជោគជ័យ!", "success")
	session.Save()
}
