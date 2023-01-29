package upload

import (
	"HelloGin/src/global"
	"HelloGin/src/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
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
	pa := path.Join("./static/", file.Filename)
	c.SaveUploadedFile(file, pa)
	res.Success(file.Filename)
	return
}
