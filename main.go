package main

import (
	"akBlog/app/config"
	"akBlog/app/server"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// 生产模式
	gin.SetMode(gin.ReleaseMode)

	// 公开服务入口
	if config.Get("isHTTPS") == "y" {
		go server.EntranceHTTPS()

		// 开启https时通常会启用http端口来跳转，避免http时错误不能打开
		isHTTP := ""
		fmt.Println("是否开启80端口,需要root权限(y/n)")
		fmt.Scanln(&isHTTP)
		if isHTTP == "y" {
			go server.EntranceHTTP("80")
			fmt.Println("成功开启80端口")
		}

	}

	// 管理员后台
	server.AdminHTTP()
}
