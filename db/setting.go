package db

import (
	"fmt"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Setting struct {
	ID    string `json:"id"`
	Title string `json:"title" form:"title" binding:"required"`
	//SiteTitle   string `json:"siteTitle" form:"siteTitle" binding:"required"`
	Description string `json:"description" form:"description"`
	Dashboard   int    `json:"dashboard" form:"dashboard" binding:"required"`
	Frontend    int    `json:"frontend" form:"frontend" binding:"required"`
	Categories  int    `json:"categories" form:"categories" binding:"required"`
	Playlist    int    `json:"playlist" form:"playlist" binding:"required"`
	Thumb       string `json:"thumb" form:"thumb" binding:"required"`
	Date        string `json:"date" form:"date" binding:"required"`
	Videos      string
}

func CreateSettingSchema() {
	mydb := Connect()
	defer mydb.Close()
	sql := `CREATE TABLE IF NOT EXISTS Setting (
		id TEXT PRIMARY KEY,
		title TEXT NOT NULL,
		siteTitle TEXT NOT NULL,
		description TEXT,
		dashboard INT NOT NULL,
		frontend INT NOT NULL,
		categories INT NOT NULL,
		playlist INT NOT NULL,
		thumb TEXT NOT NULL,
		date TEXT NOT NULL
	);
	`
	_, err := mydb.Exec(sql)
	if err != nil {
		fmt.Println("Error creating Setting schema:", err)
	}
}

func CountSettings() int {
	mydb := Connect()
	defer mydb.Close()
	var count int
	sql := `SELECT COUNT(*) FROM Setting`
	err := mydb.QueryRow(sql).Scan(&count)
	if err != nil {
		fmt.Println("Error counting settings:", err)
		return 0
	}
	return count
}

func GetSettings(limit int) []Setting {
	mydb := Connect()
	defer mydb.Close()
	sql := `SELECT id, title, description, dashboard, frontend, categories, playlist, thumb, date FROM Setting ORDER BY date DESC LIMIT ?`
	rows, err := mydb.Query(sql, limit)
	if err != nil {
		fmt.Println("Error retrieving categories:", err)
		return nil
	}
	defer rows.Close()

	var settings []Setting
	for rows.Next() {
		var setting Setting
		err := rows.Scan(&setting.ID, &setting.Title, &setting.Description, &setting.Dashboard, &setting.Frontend, &setting.Categories, &setting.Playlist, &setting.Thumb, &setting.Date)
		if err != nil {
			fmt.Println("Error scanning category:", err)
			continue
		}
		settings = append(settings, setting)
	}
	return settings
}

func CreateSetting(c *gin.Context) {
	userRole, _ := c.Get("userRole")
	session := sessions.Default(c)
	if userRole != "Admin" {
		session.AddFlash("អ្នក​គ្មាន​សិទ្ធិ​បង្កើត​​​ setting ទេ!", "error")
		session.Save()
		return
	}

	mydb := Connect()
	defer mydb.Close()
	var setting Setting

	if err := c.ShouldBind(&setting); err != nil {
		session.AddFlash("ការបង្កើត​ setting ​មាន​បញ្ហា!: "+err.Error(), "error")
		session.Save()
		return
	}
	setting.ID = uuid.New().String()
	sql := `INSERT INTO Setting VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := mydb.Exec(sql, setting.ID, setting.Title, setting.Description, setting.Dashboard, setting.Frontend, setting.Categories, setting.Playlist, setting.Thumb, setting.Date)
	if err != nil {
		session := sessions.Default(c)
		session.AddFlash("ការបង្កើត​ setting ​មាន​បញ្ហា!: "+err.Error(), "error")
		session.Save()
		return
	}

	session.AddFlash("setting ​ត្រូវ​បាន​បង្កើត​ដោយ​ជោគជ័យ!", "success")
	session.Save()
}

func PaginateSettings(c *gin.Context, dashboard int, query int) []Setting {
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
	setting := &Setting{}
	mysql := `SELECT * FROM Setting ORDER BY date DESC LIMIT ? OFFSET ?`
	rows, err := mydb.Query(mysql, limit, offset)
	if err != nil {
		fmt.Println("Error querying database:", err)
		return nil
	}
	defer rows.Close()
	var settings []Setting
	for rows.Next() {
		err := rows.Scan(&setting.ID, &setting.Title, &setting.Description, &setting.Dashboard, &setting.Frontend, &setting.Categories, &setting.Playlist, &setting.Thumb, &setting.Date)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		settings = append(settings, *setting)
	}
	return settings
}

func DeleteSetting(id string, c *gin.Context) {
	session := sessions.Default(c)
	userRole, _ := c.Get("userRole")
	if userRole != "Admin" {
		session.AddFlash("អ្នក​មិន​មាន​សិទ្ធិ​លុប​ setting ទេ!", "error")
		session.Save()
		return
	}

	mydb := Connect()
	defer mydb.Close()
	mysql := `DELETE FROM Setting WHERE id = ?`
	_, err := mydb.Exec(mysql, id)
	if err != nil {
		session.AddFlash("មាន​បញ្ហា​ក្នុង​ការលុប setting!", "error")
		session.Save()
		fmt.Println("Error deleting setting:", err)
		return
	}
	session.AddFlash("​setting ​ត្រូវ​​បានលុប​​ដោយ​ជោគជ័យ!", "success")
	session.Save()
}

func GetSetting(id string) Setting {
	mydb := Connect()
	setting := &Setting{}
	mysql := `SELECT * FROM Setting WHERE id = ?`
	row := mydb.QueryRow(mysql, id)
	err := row.Scan(&setting.ID, &setting.Title, &setting.Description, &setting.Dashboard, &setting.Frontend, &setting.Categories, &setting.Playlist, &setting.Thumb, &setting.Date)
	if err != nil {
		fmt.Println("Error scanning row:", err)
		return Setting{}
	}
	defer mydb.Close()
	return *setting
}

func UpdateSetting(c *gin.Context) {
	userRole, _ := c.Get("userRole")
	session := sessions.Default(c)
	if userRole != "Admin" {
		session.AddFlash("អ្នក​គ្មាន​សិទ្ធិកែប្រែ​ setting ​ទេ!", "error")
		session.Save()
		return
	}

	mydb := Connect()
	defer mydb.Close()
	var setting Setting
	if err := c.ShouldBind(&setting); err != nil {
		session.AddFlash("ការកែប្រែ​ setting មាន​បញ្ហា!: "+err.Error(), "error")
		session.Save()
		return
	}
	setting.ID = c.Param("id")
	sql := `UPDATE Setting SET title = ?, description = ?, dashboard = ?, frontend = ?, categories = ?, playlist = ?, thumb = ?, date = ? WHERE id = ?`
	_, err := mydb.Exec(sql, setting.Title, setting.Description, setting.Dashboard, setting.Frontend, setting.Categories, setting.Playlist, setting.Thumb, setting.Date, setting.ID)
	if err != nil {
		session.AddFlash("ការកែប្រែ​ setting មាន​បញ្ហា!: "+err.Error(), "error")
		session.Save()
		return
	}
	session.AddFlash("setting ​ត្រូវ​បាន​កែប្រែ​ដោយ​ជោគជ័យ!", "success")
	session.Save()
}
