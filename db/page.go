package db

import (
	"fmt"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Page struct {
	ID        string `json:"id"`
	Title     string `json:"title" form:"title" binding:"required"`
	Content   string `json:"content" form:"content"`
	Thumb     string `json:"thumb" form:"thumb" binding:"required"`
	Date      string `json:"date" form:"date" binding:"required"`
	Videos    string `json:"videos"`
	UpdatedAt string `json:"updated_at"`
}

func CreatePageSchema() {
	mydb := Connect()
	defer mydb.Close()
	sql := `CREATE TABLE IF NOT EXISTS Page (
		id TEXT PRIMARY KEY,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		thumb TEXT NOT NULL,
		date TEXT NOT NULL,
		updated_at TEXT DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TRIGGER IF NOT EXISTS update_category_timestamp_trigger
		AFTER UPDATE ON Category
		FOR EACH ROW
		BEGIN
			UPDATE Category
			SET updated_at = CURRENT_TIMESTAMP
			WHERE id = NEW.id;
		END;
	);
	`
	_, err := mydb.Exec(sql)
	if err != nil {
		fmt.Println("Error creating Page schema:", err)
	}
}

func CountPages() int {
	mydb := Connect()
	defer mydb.Close()
	var count int
	sql := `SELECT COUNT(*) FROM Page`
	err := mydb.QueryRow(sql).Scan(&count)
	if err != nil {
		fmt.Println("Error counting pages:", err)
		return 0
	}
	return count
}

func GetPages(limit int) []Page {
	mydb := Connect()
	defer mydb.Close()
	sql := `SELECT id, title, content , thumb, date FROM Page ORDER BY date DESC LIMIT ?`
	rows, err := mydb.Query(sql, limit)
	if err != nil {
		fmt.Println("Error retrieving pages:", err)
		return nil
	}
	defer rows.Close()

	var pages []Page
	for rows.Next() {
		var page Page
		err := rows.Scan(&page.ID, &page.Title, &page.Content, &page.Thumb, &page.Date)
		if err != nil {
			fmt.Println("Error scanning category:", err)
			continue
		}
		pages = append(pages, page)
	}
	return pages
}

func CreatePage(c *gin.Context) {
	userRole, _ := c.Get("userRole")
	session := sessions.Default(c)
	if userRole != "Admin" {
		session.AddFlash("អ្នក​គ្មាន​សិទ្ធិ​បង្កើតទំព័រ​ស្តាទិក​ទេ!", "error")
		session.Save()
		return
	}

	mydb := Connect()
	defer mydb.Close()
	var page Page
	if err := c.ShouldBind(&page); err != nil {
		session.AddFlash("ការបង្កើត​ជំពូក​មាន​បញ្ហា!: "+err.Error(), "error")
		session.Save()
		return
	}
	page.ID = uuid.New().String()
	sql := `INSERT INTO Page (id, title, content, thumb, date) VALUES (?, ?, ?, ?, ?)`
	_, err := mydb.Exec(sql, page.ID, page.Title, page.Content, page.Thumb, page.Date)
	if err != nil {
		session := sessions.Default(c)
		session.AddFlash("ការបង្កើត​ស្តាទិក​មាន​បញ្ហា!: "+err.Error(), "error")
		session.Save()
		return
	}

	session.AddFlash("ទំព័រ​ស្តាទិក​​ត្រូវ​បាន​បង្កើត​ដោយ​ជោគជ័យ!", "success")
	session.Save()
}

func GetPage(id string) Page {
	mydb := Connect()
	page := &Page{}
	mysql := `SELECT * FROM Page WHERE id = ?`
	row := mydb.QueryRow(mysql, id)
	err := row.Scan(&page.ID, &page.Title, &page.Content, &page.Thumb, &page.Date, &page.UpdatedAt)
	if err != nil {
		fmt.Println("Error scanning row:", err)
		return Page{}
	}
	defer mydb.Close()
	return *page
}

func PaginatePages(c *gin.Context, limit int, query int) []Page {
	mydb := Connect()
	defer mydb.Close()
	var page int
	if query != 0 {
		page = query
	} else {
		pageStr := c.Param("page")
		page, _ = strconv.Atoi(pageStr)
	}
	offset := (page - 1) * limit
	sql := `SELECT id, title, content, thumb, date FROM Page ORDER BY date DESC LIMIT ? OFFSET ?`
	rows, err := mydb.Query(sql, limit, offset)
	if err != nil {
		fmt.Println("Error retrieving pages:", err)
		return nil
	}
	defer rows.Close()

	var pages []Page
	for rows.Next() {
		var page Page
		err := rows.Scan(&page.ID, &page.Title, &page.Content, &page.Thumb, &page.Date)
		if err != nil {
			fmt.Println("Error scanning page:", err)
			continue
		}
		pages = append(pages, page)
	}
	return pages
}

func UpdatePage(c *gin.Context) {
	userRole, _ := c.Get("userRole")
	session := sessions.Default(c)
	if userRole != "Admin" {
		session.AddFlash("អ្នក​គ្មាន​សិទ្ធិកែប្រែ​ទំព័រ​ស្តាទិក​ទេ!", "error")
		session.Save()
		return
	}

	mydb := Connect()
	defer mydb.Close()
	var page Page
	if err := c.ShouldBind(&page); err != nil {
		session.AddFlash("ការកែប្រែ​ទំព័រ​ស្តាទិក​​មាន​បញ្ហា!: "+err.Error(), "error")
		session.Save()
		return
	}
	page.ID = c.Param("id")
	sql := `UPDATE Page SET title = ?, content = ?, thumb = ?, date = ? WHERE id = ?`
	_, err := mydb.Exec(sql, page.Title, page.Content, page.Thumb, page.Date, page.ID)
	if err != nil {
		session.AddFlash("ការកែប្រែ​ទំព័រ​ស្តាទិក​​មាន​បញ្ហា!: "+err.Error(), "error")
		session.Save()
		return
	}
	session.AddFlash("ទំព័រ​ស្តាទិក​​ត្រូវ​បាន​កែប្រែ​ដោយ​ជោគជ័យ!", "success")
	session.Save()
}
