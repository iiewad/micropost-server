package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iiewad/micropost-server/common"
	"github.com/iiewad/micropost-server/controller/micropost"
	"github.com/iiewad/micropost-server/controller/user"
	"github.com/iiewad/micropost-server/middleware"
	"github.com/iiewad/micropost-server/models"
)

func main() {
	r := gin.Default()
	common.Init()
	models.Init()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	v1 := r.Group("/api/v1")
	{
		v1.POST("/signup", user.Create)
		v1.POST("/login", user.Login)
	}
	v1.Use(middleware.Auth())
	{
		v1.POST("/micropost", micropost.Create)
	}
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
