package handlers

import (
	"akBlog/app/config"
	"akBlog/app/util"
	filehashchecking "akBlog/cmd/fileHashChecking"
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

	// 导入mirrors文件
	adminR.LoadHTMLFiles("web/mirrors.html", "web/admin.html")

	// 所有页面都要验证IP
	adminR.Use(verifyIP())

	// index页面
	adminR.GET("/", func(ctx *gin.Context) {
		// 获取_posts路径下文件夹
		dir, err := os.ReadDir("jekyll/_posts/")
		util.ErrprDisplay(err)
		fmt.Println(dir)
		// 保存路径名字
		articleType := []string{}
		for _, v := range dir {
			articleType = append(articleType, v.Name())
			fmt.Println(articleType)
		}

		ctx.HTML(200, "admin.html", gin.H{"articleType": articleType})
	})

	// 文件上传
	adminR.POST("/saveFile", func(ctx *gin.Context) {
		// 单独文件上传
		file, _ := ctx.FormFile("file")

		// 写入HASH
		filehashchecking.AddHash(file.Filename)

		//保存文件
		ctx.SaveUploadedFile(file, filehashchecking.ScanDirPath+file.Filename)

		ctx.Redirect(200, "/mirrors")
	})

	// index页面
	adminR.GET("/mirrors", func(ctx *gin.Context) {

		// 验证文件哈希值
		success, fail := filehashchecking.CheckingHash()

		// 返回网页验证记录
		ctx.HTML(200, "mirrors.html", gin.H{"success": success, "fail": fail})
	})

	return adminR
}
