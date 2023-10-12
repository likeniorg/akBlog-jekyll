package main

import (
	"akBlog/app/config"
	"akBlog/app/server"

	"github.com/gin-gonic/gin"
)

func main() {
	// 生产模式
	gin.SetMode(gin.ReleaseMode)

	// 公开服务入口
	if config.Get("isHTTPS") == "y" {
		go server.EntranceHTTPS()

	} else {
		go server.EntranceHTTP()

	}

	// 管理员后台
	server.AdminHTTP()
}
