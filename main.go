package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个 Gin 引擎
	r := gin.Default()

	// 定义路由，监听根路径并返回 PONG
	r.GET("/", func(c *gin.Context) {
		c.String(200, "PONG")
	})

	// 启动服务器，监听 8000 端口
	r.Run(":8000")
}

