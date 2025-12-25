package main

import (
	"ikbs/internal/myValidator"
	"ikbs/internal/router"
	"ikbs/lib/basic"
	"ikbs/lib/config"
	"ikbs/lib/db"
	"ikbs/lib/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	basic.Init()
	config.Init()
	logger.Init()
	db.Init()
	myValidator.Init()

	// 创建带默认中间件（日志与恢复）的 Gin 路由器
	r := gin.Default()

	//注册路由
	router.Register(r)

	// 定义简单的 GET 路由
	r.GET("/ping", func(c *gin.Context) {
		// 返回 JSON 响应
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// 默认端口 8080 启动服务器
	// 监听 0.0.0.0:8080（Windows 下为 localhost:8080）
	r.Run()
}
