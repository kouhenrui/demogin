package middleWare

import (
	"HelloGin/src/global"
	"HelloGin/src/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

var (
	userInfo = &util.UserClaims{}
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
		//fmt.Println(requestUrl)
		if strings.Contains(requestUrl, "/ws/") {
			//ts := time.Since(t)
			//fmt.Println(requestUrl)
			fmt.Println("ws身份认证执行结束")
		} else {
			reqUrl := strings.Split(requestUrl, "/api/")
			paths := global.ReuqestPaths
			pathIsExist := util.ExistIn(reqUrl[1], paths)
			if !pathIsExist {
				fmt.Println("身份验证")
				user := util.AnalysyToken(c)
				//fmt.Println("userInfo", user)
				c.Set("user", user)
			}
			ts := time.Since(t)
			fmt.Println("time", ts)
			fmt.Println("身份认证执行结束")
		}

	}
}
