package handler

import (
	"database/sql"
	"embed"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	"samuelemusiani/project_90107/config"
)

//go:embed dist
var frontend embed.FS

// Istanza database globale
var db *sql.DB

func InitAndServe(conf *config.Config, database *sql.DB) error {
	db = database

	router := gin.Default()

	// middleware for static files (frontend)
	router.Use(static.Serve("/", static.EmbedFolder(frontend, "dist")))

	api := router.Group("/api")
	api.GET("/", root)
	api.GET("/users", users)
	api.POST("/users", insertUser)

	// Read index.html into memory
	index, err := frontend.ReadFile("dist/index.html")
	if err != nil {
		return err
	}

	// If no route match is probably a vue route. So we return the index.html
	// and the vue-router takes from here
	router.NoRoute(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.String(), "/api") {
			c.JSON(http.StatusNotFound, "")
			return
		}
		c.Data(http.StatusOK, "text/html", index)
	})

	router.Run(fmt.Sprintf(":%d", conf.Server.Port))
	return nil
}

func root(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func users(c *gin.Context) {
	rows, err := db.Query("SELECT * FROM users")

	if err != nil {
		slog.With("err", err).Error("Executing query")
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	defer rows.Close()

	var users []string
	for rows.Next() {
		var username string
		err = rows.Scan(&username)
		if err != nil {
			slog.With("err", err).Error("Scanning row")
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		users = append(users, username)
	}

	c.JSON(200, gin.H{
		"users": users,
	})
}

func insertUser(c *gin.Context) {
	var user struct {
		User string
	}

	if err := c.BindJSON(&user); err != nil {
		slog.With("err", err).Error("Binding JSON")
		c.JSON(400, gin.H{
			"error": "Username not found",
		})
		return
	}

	_, err := db.Exec("INSERT INTO users (username) VALUES ($1)", user.User)
	if err != nil {
		slog.With("err", err).Error("Inserting user")
		c.JSON(500, gin.H{
			"error": "Adding user to db",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "User added",
	})
}
