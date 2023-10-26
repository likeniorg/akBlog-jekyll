package handlers

import (
	"akBlog/app/config"
	"akBlog/app/mirrors"
	"akBlog/app/util"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func verifyIP() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.ClientIP() == config.Get("serverIP") || ctx.ClientIP() == "::1" {
			ctx.Next()
		} else {
			ctx.String(200, "亲亲你涉嫌非法访问，请不要继续尝试了")
			ctx.Abort()
		}
	}
}

// 管理员后台服务器
func AdminServer() http.Handler {
	// 使用默认路由
	adminR := gin.New()

	// 所有页面都要验证IP
	adminR.Use(verifyIP())

	// index页面
	adminR.GET("/", func(ctx *gin.Context) {
		// 获取_posts路径下文件夹
		dir, err := os.ReadDir("jekyll/_posts/")
		util.ErrprDisplay(err)
		// 保存路径名字
		articleType := []string{}
		for _, v := range dir {
			articleType = append(articleType, v.Name())
			fmt.Println(articleType)
		}

		ctx.HTML(200, "admin.html", gin.H{"articleType": articleType})
	})

	// 开启镜像站及文件上传功能
	mirrors.StartAllHandlers(adminR)
	return adminR
}
