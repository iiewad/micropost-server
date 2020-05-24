package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iiewad/micropost-server/db"
	"github.com/iiewad/micropost-server/models"
)

func main() {
	r := gin.Default()
	db.Init()
	models.Init()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
