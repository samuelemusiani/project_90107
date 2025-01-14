package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"samuelemusiani/project_90107/config"
)

func InitAndServe(conf *config.Config) {
	r := gin.Default()
	api := r.Group("/api")
	api.GET("/", root)

	r.Run(fmt.Sprintf(":%d", conf.Server.Port))
}

func root(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}
