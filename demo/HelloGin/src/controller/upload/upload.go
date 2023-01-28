package upload

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path"
)

func Routers(e *gin.Engine) {

	e.MaxMultipartMemory = 8 << 20
	e.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")

		if err != nil {
			fmt.Println("文件类型错误", err)
			c.String(1000, "文件格式出错")
		}
		pa := path.Join("./static/", file.Filename)

		log.Println("path:", pa)

		//c.SaveUploadedFile(file, "../../static"+file.Filename)
		c.SaveUploadedFile(file, pa)
		//fmt.Print(file, "打印文件")
		c.String(http.StatusOK, file.Filename)

		//types := c.DefaultPostForm("type", "post")
		//name := c.PostForm("name")
		//pwd := c.PostForm("pwd")
		//c.String(http.StatusOK, fmt.Sprintf("name:%s ,pwd:%s,type:%s", name, pwd, types))
	})
}
