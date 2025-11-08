package db

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Post struct {
	ID         string `json:"id"`
	Title      string `json:"title" form:"title" binding:"required"`
	Content    string `json:"content" form:"content"`
	Categories string `json:"categories" form:"categories" binding:"required"`
	Thumb      string `json:"thumb" form:"thumb" binding:"required"`
	Date       string `json:"date" form:"date" binding:"required"`
	Videos     string `json:"videos" form:"videos"`
	Author     string `json:"author"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

func CreatePostSchema() {
	mydb := Connect()
	defer mydb.Close()
	sql := `CREATE TABLE IF NOT EXISTS Post (
        id TEXT PRIMARY KEY,
        title TEXT NOT NULL,
        content TEXT,
		categories TEXT NOT NULL,
		thumb TEXT NOT NULL,
		date TEXT NOT NULL,
		videos TEXT,
		author TEXT NOT NULL,
		created_at TEXT DEFAULT CURRENT_TIMESTAMP,
        updated_at TEXT DEFAULT CURRENT_TIMESTAMP
    );

	CREATE TRIGGER IF NOT EXISTS delete_old_posts 
        AFTER INSERT ON Post
        BEGIN
            DELETE FROM Post 
            WHERE created_at < DATE('now', '-90 days') AND categories LIKE '%news%';
        END;

    CREATE TRIGGER IF NOT EXISTS update_timestamp_trigger
        AFTER UPDATE ON Post
        FOR EACH ROW
        BEGIN
            UPDATE Post
            SET updated_at = CURRENT_TIMESTAMP
    		WHERE id = NEW.id;
        END;
	`

	_, err := mydb.Exec(sql)
	if err != nil {
		panic(err)
	}
}

func CountPosts() int {
	mydb := Connect()
	defer mydb.Close()
	var count int
	sql := `SELECT COUNT(*) FROM Post`
	row := mydb.QueryRow(sql)
	row.Scan(&count)
	return count
}

func CreatePost(c *gin.Context) {
	mydb := Connect()
	defer mydb.Close()

	post := Post{}
	session := sessions.Default(c)
	if err := c.ShouldBind(&post); err != nil {
		session.AddFlash("មាន​បញ្ហា​ក្នុង​ការបង្កើត​ការផ្សាយ!", "error")
		session.Save()
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id := uuid.New()
	title := c.PostForm("title")
	content := c.PostForm("content")
	categories := c.PostForm("categories")
	thumb := c.PostForm("thumb")
	date := c.PostForm("date")
	videos := c.PostForm("videos")
	userId, _ := c.Get("userId")
	author := userId.(string)
	created_at := time.Now().Format("2006-01-02T15:04:05")
	updated_at := time.Now().Format("2006-01-02T15:04:05")

	sql := `INSERT INTO Post VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	mydb.Exec(sql, id, title, content, categories, thumb, date, videos, author, created_at, updated_at)
	defer mydb.Close()
	session.AddFlash("ការផ្សាយ​ត្រូវ​បាន​បង្កើត​ឡើងដោយ​ជោគជ័យ!", "success")
	session.Save()
}

func GetPosts(limit int) []Post {
	mydb := Connect()
	post := &Post{}
	mysql := `SELECT * FROM Post ORDER BY date DESC LIMIT ?`
	rows, err := mydb.Query(mysql, limit)
	if err != nil {
		fmt.Println("Error querying database:", err)
		return nil
	}
	defer rows.Close()
	var posts []Post
	for rows.Next() {
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Categories, &post.Thumb, &post.Date, &post.Videos, &post.Author, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		posts = append(posts, *post)
	}
	return posts
}

func DeletePost(id string, c *gin.Context) {
	mydb := Connect()
	session := sessions.Default(c)
	userRole, _ := c.Get("userRole")
	if userRole != "Admin" {
		post := GetPost(id)
		author, _ := c.Get("userId")
		if post.Author != author {
			session.AddFlash("អ្នក​មិន​មាន​សិទ្ធិ​លុប​ការ​ផ្សាយ​របស់​អ្នក​ដទៃទេ!", "error")
			session.Save()
			return
		} else {
			mysql := `DELETE FROM Post WHERE id = ? AND author = ?`
			_, err := mydb.Exec(mysql, id, author)
			if err != nil {
				session.AddFlash("មាន​បញ្ហា​ក្នុង​ការលុបការផ្សាយ!", "error")
				session.Save()
				fmt.Println("Error deleting post:", err)
				return
			}
			defer mydb.Close()
		}
	} else if userRole == "Admin" {
		mysql := `DELETE FROM Post WHERE id = ?`
		_, err := mydb.Exec(mysql, id)
		if err != nil {
			session.AddFlash("មាន​បញ្ហា​ក្នុង​ការលុបការផ្សាយ!", "error")
			session.Save()
			fmt.Println("Error deleting post:", err)
			return
		}
		defer mydb.Close()
	}
	session.AddFlash("​ការផ្សាយត្រូវ​​បានលុប​​ដោយ​ជោគជ័យ!", "success")
	session.Save()
}

func GetPost(id string) Post {
	mydb := Connect()
	post := &Post{}
	mysql := `SELECT * FROM Post WHERE id = ?`
	row := mydb.QueryRow(mysql, id)
	err := row.Scan(&post.ID, &post.Title, &post.Content, &post.Categories, &post.Thumb, &post.Date, &post.Videos, &post.Author, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		fmt.Println("Error scanning row:", err)
		return Post{}
	}
	defer mydb.Close()
	return *post
}

func UpdatePost(c *gin.Context) {
	mydb := Connect()
	defer mydb.Close()
	session := sessions.Default(c)

	title := c.PostForm("title")
	content := c.PostForm("content")
	categories := c.PostForm("categories")
	thumb := c.PostForm("thumb")
	date := c.PostForm("date")
	videos := c.PostForm("videos")
	postID := c.Param("id")
	userRole, _ := c.Get("userRole")
	mysql := `UPDATE Post SET title = ?, content = ?, categories = ?, thumb = ?, date = ?, videos = ? WHERE id = ?`

	if userRole != "Admin" {
		post := GetPost(postID)
		author, _ := c.Get("userId")
		if post.Author != author {
			session.AddFlash("អ្នក​មិន​មាន​សិទ្ធិ​កែប្រែ​ការ​ផ្សាយ​របស់​អ្នក​ដទៃទេ!", "error")
			session.Save()
			return
		}
		_, err := mydb.Exec(mysql, title, content, categories, thumb, date, videos, postID)
		if err != nil {
			session.AddFlash("មាន​បញ្ហា​ក្នុង​ការកែប្រែ​ការផ្សាយ!", "error")
			session.Save()
			fmt.Println("Error updating post:", err)
			return
		}
	} else if userRole == "Admin" {
		_, err := mydb.Exec(mysql, title, content, categories, thumb, date, videos, postID)
		if err != nil {
			session.AddFlash("មាន​បញ្ហា​ក្នុង​ការកែប្រែ​ការផ្សាយ!", "error")
			session.Save()
			fmt.Println("Error updating post:", err)
			return
		}
	}
	session.AddFlash("ការផ្សាយ​ត្រូវ​បាន​កែប្រែ​ដោយ​ជោគជ័យ!", "success")
	session.Save()
}

func PaginatePosts(c *gin.Context, dashboard int, query int) []Post {
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
	post := &Post{}
	mysql := `SELECT * FROM Post ORDER BY date DESC LIMIT ? OFFSET ?`
	rows, err := mydb.Query(mysql, limit, offset)
	if err != nil {
		fmt.Println("Error querying database:", err)
		return nil
	}
	defer rows.Close()
	var posts []Post
	for rows.Next() {
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Categories, &post.Thumb, &post.Date, &post.Videos, &post.Author, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		posts = append(posts, *post)
	}
	return posts
}
