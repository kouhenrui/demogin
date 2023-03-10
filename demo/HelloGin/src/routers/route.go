package routers

import (
	"HelloGin/src/global"
	middleWare "HelloGin/src/middleware"
	"HelloGin/src/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Option func(engine *gin.Engine)

var options = []Option{}

func Include(opts ...Option) {
	options = append(options, opts...)
}
func InitRoute() *gin.Engine {
	log.Println("路由初始化调用")
	r := gin.New()
	r.Use(Cors())                               //民间跨域
	r.StaticFS("/img", http.Dir("./static"))    //加载静态资源，一般是上传的资源，例如用户上传的图片
	r.StaticFS("/dynamic", http.Dir("./video")) //加载静态资源，一般是上传的资源，例如用户上传的图片
	r.Use(middleWare.LoggerMiddleWare())        //日志中间件
	r.Use(middleWare.GolbalMiddleWare())        //全局中间件
	r.Use(middleWare.AuthMiddleWare())          //身份认证
	r.Use(middleWare.Recover)                   //错误捕捉
	r.NoRoute(HandleNotFound)                   //路由未找到
	r.NoMethod(HandleNotAllowed)                //方法未找到
	r.MaxMultipartMemory = 64 << 20             //64Mb
	for _, ii := range options {
		ii(r)
	} //挂载模块
	return r

}

// 404
func HandleNotFound(c *gin.Context) {
	res := global.NewResult(c)
	res.Error(http.StatusNotFound, util.RESOURCE_NOT_FOUND_ERROR)
	return
}

func HandleNotAllowed(c *gin.Context) {
	res := global.NewResult(c)
	res.Error(http.StatusMethodNotAllowed, util.REQUEST_METHOD_NOT_ALLOWED_ERROE)
	return
}

func Cors() gin.HandlerFunc {
	//
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Set("content-type", "application/json")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
