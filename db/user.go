package db

import "github.com/google/uuid"

func CreateRootUser() {
	mydb := Connect()
	sql := `INSERT INTO User VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	id := uuid.New()
	title := "Sokhavuth"
	email := "sokhavuth@khmerweb.app"
	password := "Tin2025"
	role := "Admin"
	thumb := ""
	content := ""
	date := ""
	_, err := mydb.Exec(sql, id, title, email, password, role, thumb, content, date)
	if err != nil {
		panic(err)
	}
	defer mydb.Close()
}
