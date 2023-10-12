package handlers

import (
	filehashchecking "akBlog/cmd/fileHashChecking"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

// 前端页面生成
func WebPage(r *gin.Engine) {
	// 加载mirrors文件
	r.LoadHTMLFiles("web/mirrors.html")

	r.GET("/", func(ctx *gin.Context) {
		ctx.File("./web/index.html")
	})

	r.GET("/categories", func(ctx *gin.Context) {
		ctx.File("./web/categories/index.html")
	})

	r.GET("/tags", func(ctx *gin.Context) {
		ctx.File("./web/tags/index.html")
	})

	r.GET("/archives", func(ctx *gin.Context) {
		ctx.File("./web/archives/index.html")
	})

	r.GET("/about", func(ctx *gin.Context) {
		ctx.File("./web/about/index.html")
	})

	r.GET("/mirrors", func(ctx *gin.Context) {
		success, fail := filehashchecking.CheckingHash()
		ctx.HTML(200, "mirrors.html", gin.H{"sucess": success, "fail": fail})
	})

	// 文章页面路径生成
	func() {
		article := r.Group("/posts/")
		dir, _ := os.ReadDir("./web/posts/")
		// var page []string

		// 获取文章路径
		for _, root := range dir {
			if root.IsDir() {
				article.GET(root.Name(), func(ctx *gin.Context) {
					// 这里是进入页面后单独执行，变量与上下文无关

					// 判断url路径对应文章目录是否存在。避免导入其他文件
					isExistDir := false

					dir, _ := os.ReadDir("./web/posts/")
					for _, root := range dir {
						if root.IsDir() {
							fmt.Println(ctx.Request.URL.Path)
							// 从请求路径来判断导入哪个文章
							if ctx.Request.URL.Path == "/posts/"+root.Name() {
								isExistDir = true
							}
						}
					}
					if isExistDir {
						ctx.File("./web" + ctx.Request.URL.Path + "/index.html")
					} else {
						ctx.String(200, "非法请求页面")
					}

				})
			}
		}

	}()
}
