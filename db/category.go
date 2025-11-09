package db

import (
	"fmt"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Category struct {
	ID        string `json:"id"`
	Title     string `json:"title" form:"title" binding:"required"`
	Thumb     string `json:"thumb" form:"thumb" binding:"required"`
	Date      string `json:"date" form:"date" binding:"required"`
	Videos    string `json:"videos"`
	UpdatedAt string `json:"updated_at"`
}

func CreateCategorySchema() {
	mydb := Connect()
	defer mydb.Close()
	sql := `CREATE TABLE IF NOT EXISTS Category (
		id TEXT PRIMARY KEY,
		title TEXT NOT NULL,
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
		fmt.Println("Error creating Category schema:", err)
	}
}

func CountCategories() int {
	mydb := Connect()
	defer mydb.Close()
	var count int
	sql := `SELECT COUNT(*) FROM Category`
	err := mydb.QueryRow(sql).Scan(&count)
	if err != nil {
		fmt.Println("Error counting categories:", err)
		return 0
	}
	return count
}

func CreateCategory(c *gin.Context) {
	userRole, _ := c.Get("userRole")
	session := sessions.Default(c)
	if userRole != "Admin" {
		session.AddFlash("អ្នក​គ្មាន​សិទ្ធិ​បង្កើតជំពូក​ទេ!", "error")
		session.Save()
		return
	}

	mydb := Connect()
	defer mydb.Close()
	var category Category

	if err := c.ShouldBind(&category); err != nil {
		session.AddFlash("ការបង្កើត​ជំពូក​មាន​បញ្ហា!: "+err.Error(), "error")
		session.Save()
		return
	}
	category.ID = uuid.New().String()
	sql := `INSERT INTO Category (id, title, thumb, date) VALUES (?, ?, ?, ?)`
	_, err := mydb.Exec(sql, category.ID, category.Title, category.Thumb, category.Date)
	if err != nil {
		session := sessions.Default(c)
		session.AddFlash("ការបង្កើត​ជំពូក​មាន​បញ្ហា!: "+err.Error(), "error")
		session.Save()
		return
	}

	session.AddFlash("ជំពូក​ត្រូវ​បាន​បង្កើត​ដោយ​ជោគជ័យ!", "success")
	session.Save()

}

func GetCategories(limit int) []Category {
	mydb := Connect()
	defer mydb.Close()
	sql := `SELECT id, title, thumb, date FROM Category ORDER BY date DESC LIMIT ?`
	rows, err := mydb.Query(sql, limit)
	if err != nil {
		fmt.Println("Error retrieving categories:", err)
		return nil
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var category Category
		err := rows.Scan(&category.ID, &category.Title, &category.Thumb, &category.Date)
		if err != nil {
			fmt.Println("Error scanning category:", err)
			continue
		}
		categories = append(categories, category)
	}
	return categories
}

func GetCategory(id string) Category {
	mydb := Connect()
	category := &Category{}
	mysql := `SELECT * FROM Category WHERE id = ?`
	row := mydb.QueryRow(mysql, id)
	err := row.Scan(&category.ID, &category.Title, &category.Thumb, &category.Date, &category.UpdatedAt)
	if err != nil {
		fmt.Println("Error scanning row:", err)
		return Category{}
	}
	defer mydb.Close()
	return *category
}

func PaginateCategories(c *gin.Context, limit int, query int) []Category {
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
	sql := `SELECT id, title, thumb, date FROM Category ORDER BY date DESC LIMIT ? OFFSET ?`
	rows, err := mydb.Query(sql, limit, offset)
	if err != nil {
		fmt.Println("Error retrieving categories:", err)
		return nil
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var category Category
		err := rows.Scan(&category.ID, &category.Title, &category.Thumb, &category.Date)
		if err != nil {
			fmt.Println("Error scanning category:", err)
			continue
		}
		categories = append(categories, category)
	}
	return categories
}

func UpdateCategory(c *gin.Context) {
	userRole, _ := c.Get("userRole")
	session := sessions.Default(c)
	if userRole != "Admin" {
		session.AddFlash("អ្នក​គ្មាន​សិទ្ធិកែប្រែ​ជំពូក​ទេ!", "error")
		session.Save()
		return
	}

	mydb := Connect()
	defer mydb.Close()
	var category Category
	if err := c.ShouldBind(&category); err != nil {
		session.AddFlash("ការកែប្រែ​ជំពូក​មាន​បញ្ហា!: "+err.Error(), "error")
		session.Save()
		return
	}
	category.ID = c.Param("id")
	sql := `UPDATE Category SET title = ?, thumb = ?, date = ? WHERE id = ?`
	_, err := mydb.Exec(sql, category.Title, category.Thumb, category.Date, category.ID)
	if err != nil {
		session.AddFlash("ការកែប្រែ​ជំពូក​មាន​បញ្ហា!: "+err.Error(), "error")
		session.Save()
		return
	}
	session.AddFlash("ជំពូក​ត្រូវ​បាន​កែប្រែ​ដោយ​ជោគជ័យ!", "success")
	session.Save()
}

func DeleteCategory(c *gin.Context) {
	userRole, _ := c.Get("userRole")
	session := sessions.Default(c)
	if userRole != "Admin" {
		session.AddFlash("អ្នក​គ្មាន​សិទ្ធិ​លុប​ជំពូក​ទេ!", "error")
		session.Save()
		return
	}

	mydb := Connect()
	defer mydb.Close()
	id := c.Param("id")

	if userRole == "Admin" {
		mysql := `DELETE FROM Category WHERE id = ?`
		_, err := mydb.Exec(mysql, id)
		if err != nil {
			session.AddFlash("មាន​បញ្ហា​ក្នុង​ការលុបជំពូក!", "error")
			session.Save()
			fmt.Println("Error deleting post:", err)
			return
		}
		session.AddFlash("ជំពូក​ត្រូវ​បាន​លុប​ដោយជោគជ័យ!", "success")
		session.Save()
	}
}
