package com

import (
	"HelloGin/src/global"
	"HelloGin/src/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

var (
	fileMax  = global.FileMax
	videoMax = global.VideoMax
)

/**
 * @ClassName commonController
 * @Description 常用的公共接口
 * @Author khr
 * @Date 2023/5/8 10:53
 * @Version 1.0
 */

func Routers(e *gin.Engine) {

	commonGroup := e.Group("/api/common")
	{
		commonGroup.GET("/captcha", getCaptcha)
		commonGroup.POST("/file", upload)
		commonGroup.POST("/video", uploadVideo)
	}
}

/*
 * @MethodName getCaptcha
 * @Description 签发图片验证码
 * @Author khr
 * @Date 2023/5/8 11:00
 */

func getCaptcha(c *gin.Context) {

	err, capt := util.CreateCaptcha()
	if err != nil {
		c.Error(err)
		return
	}
	c.Set("res", capt)

	//if resData != nil {
	//	res.Success(endData)
	//	return
	//}

}

/*
 * @MethodName upload
 * @Description 上传单个图片,返回字符串
 * @Author khr
 * @Date 2023/5/8 11:02
 */

func upload(c *gin.Context) {
	res := global.NewResult(c)
	file, err := c.FormFile("file")
	if err != nil {
		res.Error(http.StatusBadRequest, util.FILE_TYPE_ERROR)
		return
	}
	if c.Request.ContentLength > fileMax {
		res.Error(http.StatusBadRequest, util.FILE_TOO_LARGE)
		return
	}
	//获取上传文件的类型
	filetype := file.Header.Get("Content-Type")
	types := strings.Split(filetype, "/")
	if types[0] != "image" {
		res.Error(http.StatusBadRequest, util.FILE_TYPE_ERROR)
		return
	}
	name := time.Now().Unix()
	filename := file.Filename
	suffix := strings.Split(filename, ".")
	nameSuffix := suffix[1]
	t := util.ExistIn(nameSuffix, global.PictureType)
	if !t {
		res.Error(http.StatusBadRequest, util.FILE_SUFFIX_ERROR)
		return
	}
	file.Filename = strconv.FormatInt(name, 10) + "." + nameSuffix

	filePath := path.Join(global.FilePath)
	_, e := os.Stat(filePath)
	if e != nil {
		os.Mkdir(global.FilePath, os.ModePerm)
	}
	pa := path.Join("./"+global.FilePath+"/", file.Filename)
	c.SaveUploadedFile(file, pa)
	res.Success(file.Filename)
	return
}

/*
 * @MethodName uploadVideo
 * @Description 上传视频
 * @Author khr
 * @Date 2023/5/8 11:03
 */

func uploadVideo(c *gin.Context) {
	res := global.NewResult(c)
	file, err := c.FormFile("video")
	//fmt.Println(err, "111111")
	if err != nil {
		res.Error(http.StatusBadRequest, util.FILE_TYPE_ERROR)
		return
	}
	if c.Request.ContentLength > videoMax {
		res.Error(http.StatusBadRequest, util.FILE_TOO_LARGE)
		return

	}
	//获取上传文件的类型
	filetype := file.Header.Get("Content-Type")
	types := strings.Split(filetype, "/")
	fmt.Println(types, "文件类型")
	if types[0] != "video" {
		res.Error(http.StatusBadRequest, util.FILE_TYPE_ERROR)
		return
	}
	name := time.Now().Unix()
	filename := file.Filename
	suffix := strings.Split(filename, ".")
	nameSuffix := suffix[1]
	t := util.ExistIn(nameSuffix, global.VideoType)
	if !t {
		res.Error(http.StatusBadRequest, util.FILE_SUFFIX_ERROR)
		return
	}
	file.Filename = strconv.FormatInt(name, 10) + "." + nameSuffix
	filePath := path.Join(global.VideoPath)
	_, e := os.Stat(filePath)
	if e != nil {
		os.Mkdir(global.VideoPath, os.ModePerm)
	}
	pa := path.Join("./"+global.VideoPath+"/", file.Filename)
	c.SaveUploadedFile(file, pa)
	res.Success(file.Filename)
	return
}
