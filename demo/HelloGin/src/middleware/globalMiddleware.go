package middleWare

import (
	"HelloGin/src/global"
	"HelloGin/src/pojo"
	"HelloGin/src/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

var (
	user = &pojo.User{}
)

func GolbalMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		//t := time.Now()
		//fmt.Println("全局中间件开始执行")
		//requestUrl := c.Request.URL.String()
		//reqUrl := strings.Split(requestUrl, "/api/")
		//
		//paths := global.ReuqestPaths
		//pathIsExist := existIn(reqUrl[1], paths)
		//ts := time.Since(t)
		//fmt.Println("time", ts)
		//fmt.Println("全局中间件执行结束")

	}
}
func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("身份认证开始执行")
		t := time.Now()
		requestUrl := c.Request.URL.String()
		reqUrl := strings.Split(requestUrl, "/api/")

		paths := global.ReuqestPaths
		pathIsExist := existIn(reqUrl[1], paths)
		fmt.Println(pathIsExist, "打印是否存在白名单")
		if !pathIsExist {
			fmt.Println("身份验证")
			user := util.AnalysyToken(c)
			c.Set("user", user)
		}
		//fmt.Println("身份验证不需要，存在白名单中")
		ts := time.Since(t)
		fmt.Println("time", ts)
		fmt.Println("身份认证执行结束")
	}
}

func existIn(requestUrl string, paths []string) bool {
	for _, v := range paths {
		fmt.Println(v, requestUrl)
		if requestUrl == v {
			return true
		}
	}
	return false
}
