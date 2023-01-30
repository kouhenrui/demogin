package upload

import (
	"HelloGin/src/global"
	"HelloGin/src/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
	"strconv"
	"strings"
	"time"
)

func Routers(e *gin.Engine) {

	uploadGroup := e.Group("/api/upload")
	{
		uploadGroup.POST("/file", upload)
	}

}

//e.MaxMultipartMemory = 8 << 20
func upload(c *gin.Context) {
	res := global.NewResult(c)
	file, err := c.FormFile("file")
	if err != nil {
		res.Error(http.StatusBadRequest, util.FILE_TYPE_ERROR)
		return
	}
	//获取上传文件的类型
	filetype := file.Header.Get("Content-Type")
	types := strings.Split(filetype, "/")

	//fmt.Println(types, "文件类型")
	if types[0] != "image" {
		res.DiyErr(http.StatusBadRequest, util.FILE_TYPE_ERROR)
		return
	}
	name := time.Now().Unix()
	filename := file.Filename
	suffix := strings.Split(filename, ".")
	nameSuffix := suffix[1]
	t := util.ExistIn(nameSuffix, global.PictureType)
	if !t {
		res.DiyErr(http.StatusBadRequest, util.FILE_SUFFIX_ERROR)
		return
	}
	file.Filename = strconv.FormatInt(name, 10) + "." + nameSuffix
	pa := path.Join("./static/", file.Filename)
	c.SaveUploadedFile(file, pa)
	res.Success(file.Filename)
	return
}
