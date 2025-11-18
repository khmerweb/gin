// frontend/frontend.go
package frontend

import (
	"encoding/json"
	"fmt"
	"gin/backend"
	"gin/db"
	"html/template"
	"strconv"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/", func(c *gin.Context) {
		siteTitle := backend.Setup().Title
		limit := backend.Setup().Playlist
		playlists := db.GetPlaylists(limit)
		countPlaylists := db.CountPlaylists()
		frontend := backend.Setup().Frontend
		c.HTML(200, "home", gin.H{
			"Title":          siteTitle,
			"Playlists":      playlists,
			"CountPlaylists": countPlaylists,
			"Frontend":       frontend,
			"PageUrl":        "/",
		})
	})

	router.GET("/:category/:page", func(c *gin.Context) {
		siteTitle := backend.Setup().Title
		frontend := backend.Setup().Frontend
		limit := backend.Setup().Playlist
		posts := db.PaginatePosts(c, limit, 0)
		category := c.Param("category")
		pageStr := c.Param("page")
		currentPage, _ := strconv.Atoi(pageStr)
		var pageURL string

		if category == "national" {
			pageURL = "news"
		} else if category == "global" {
			pageURL = "news"
		} else if category == "opinion" {
			pageURL = "opinion"
		} else if category == "documentary" {
			pageURL = "documentary"
		} else if category == "sport" {
			pageURL = "sport"
		} else if category == "Khmer" {
			pageURL = "movie"
		} else if category == "Thai" {
			pageURL = "movie"
		} else if category == "Chinese" {
			pageURL = "movie"
		} else if category == "Korean" {
			pageURL = "movie"
		} else if category == "world" {
			pageURL = "movie"
		} else if category == "movies" {
			pageURL = "movie"
		} else if category == "travel" {
			pageURL = "travel"
		} else if category == "simulation" {
			pageURL = "simulation"
		} else if category == "food" {
			pageURL = "entertainment"
		} else if category == "music" {
			pageURL = "entertainment"
		} else if category == "game" {
			pageURL = "entertainment"
		}

		count := db.CountPostsByCategory(category)
		pageNumbers := make([]int, 0)
		pageCount := (count + frontend - 1) / frontend
		for i := 0; i < pageCount; i++ {
			pageNumbers = append(pageNumbers, i+1)
		}

		c.HTML(200, "category-frontend", gin.H{
			"Title":       siteTitle,
			"Posts":       posts,
			"Frontend":    frontend,
			"Category":    category,
			"PageNumbers": pageNumbers,
			"CurrentPage": currentPage,
			"PageUrl":     pageURL,
		})
	})

	router.GET("/post/:id", func(c *gin.Context) {
		session := sessions.Default(c)
		siteTitle := backend.Setup().Title
		id := c.Param("id")
		post := db.GetPost(id)
		parts := strings.Split(post.Categories, ",")
		category := parts[0]
		randomPosts := db.GetRandomPosts(7, category, id)
		author := db.GetUser(post.Author)
		authorName := author.Title
		userRole := session.Get("userRole")

		jsonDataString, err := json.Marshal(post)
		if err != nil {
			fmt.Println("Error marshalling string data:", err)
		}
		Post := string(jsonDataString)
		SafeHTML := template.HTML(post.Content)
		c.HTML(200, "post", gin.H{
			"Title":       siteTitle,
			"POST":        post,
			"Post":        Post,
			"AuthorName":  authorName,
			"UserRole":    userRole,
			"SafeHTML":    SafeHTML,
			"RandomPosts": randomPosts,
		})
	})
}
