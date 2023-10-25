package article

import (
	"akBlog/app/util"

	"github.com/gin-gonic/gin"
)

const articlePath string = "jekyll/_posts/"

func Upload() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取上传分区

		// 获取上传文件
		fileName, err := ctx.FormFile("filename")
		util.ErrprDisplay(err)

		// 保存文件
		err = ctx.SaveUploadedFile(fileName, articlePath+fileName.Filename)
		util.ErrprDisplay(err)

		ctx.String(200, "上传成功")
	}
}

func Delete() {

}
