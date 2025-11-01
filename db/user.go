package db

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func CreateRootUser() {
	mydb := Connect()
	sql := `INSERT INTO User VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	id := uuid.New()
	title := "Sokhavuth"
	email := "sokhavuth@khmerweb.app"
	passwordbyte := []byte("***********")
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordbyte, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	password := string(hashedPassword)
	role := "Admin"
	thumb := ""
	content := ""
	date := ""
	mydb.Exec(sql, id, title, email, password, role, thumb, content, date)
	defer mydb.Close()
}
