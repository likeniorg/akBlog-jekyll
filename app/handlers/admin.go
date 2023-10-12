package handlers

import (
	"akBlog/app/config"
	filehashchecking "akBlog/cmd/fileHashChecking"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func verifyIP() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.ClientIP() != config.Get("serverIP") {
			fmt.Println("c" + ctx.ClientIP())
			ctx.String(200, "亲亲你涉嫌非法访问，请不要继续尝试了")
			ctx.Abort()
		}
	}
}

// 管理员后台服务器
func AdminServer() http.Handler {
	// 使用默认路由
	adminR := gin.New()

	// 导入mirrors文件
	adminR.LoadHTMLFiles("web/mirrors.html")

	// index页面
	adminR.GET("/", verifyIP(), func(ctx *gin.Context) {

		// 验证文件哈希值
		success, fail := filehashchecking.CheckingHash()

		// 返回网页验证记录
		ctx.HTML(200, "mirrors.html", gin.H{"success": success, "fail": fail})
	})

	// 文件上传
	adminR.POST("/saveFile", verifyIP(), func(ctx *gin.Context) {
		// 单独文件上传
		file, _ := ctx.FormFile("file")

		//保存文件
		ctx.SaveUploadedFile(file, filehashchecking.ScanDirPath+file.Filename)

		// 写入HASH
		filehashchecking.AddHash(file.Filename)
	})

	return adminR
}
