package middleWare

import (
	"HelloGin/src/global"
	"HelloGin/src/logger"
	"HelloGin/src/pojo"
	"HelloGin/src/util"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"math"
	"net/http"
	"strings"
	"time"
)

var userInfo util.UserClaims
var permissionServiceImpl = pojo.RbacPermission()

// 全局路由中间件
func GolbalMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("身份认证开始执行")
		t := time.Now()
		requestUrl := c.Request.URL.String()
		reqUrl := strings.Split(requestUrl, "/api/")
		paths := global.ReuqestPaths
		pathIsExist := util.ExistIn(reqUrl[1], paths)
		if !pathIsExist {
			_, ok := c.Get("ok")
			//fmt.Println("OK", ok)
			if ok {
				c.Next()
				//c.Abort()
				//res.Err(util.AUTHENTICATION_FAILED)
				//return
			}
			judge := util.AnalysyToken(c)
			//fmt.Println("验证token是否存在，", judge)
			if !judge {
				c.AbortWithStatusJSON(http.StatusUnauthorized, util.NO_AUTHORIZATION)
				return
			}
			userInfo = util.ParseToken(c.GetHeader("Authorization"))
			c.Set("id", userInfo.Id)
			c.Set("name", userInfo.Name)
			c.Set("role", userInfo.Role)
			c.Set("account", userInfo.Account)
			c.Set("role_name", userInfo.RoleName)
			c.Set("ok", true)
			c.Next()
		}
		ts := time.Since(t)
		fmt.Println("time", ts)
		fmt.Println("身份认证执行结束")

	}
}

// 权限路由中间件
func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestUrl := c.Request.URL.String()
		reqUrl := strings.Split(requestUrl, "/api/")
		rolename := global.RoleName
		paths := global.ReuqestPaths
		pathIsExist := util.ExistIn(reqUrl[1], paths)
		//登录跳过权限验证
		if !pathIsExist {
			//验证身份
			_, y := c.Get("ok")
			//通过身份验证
			if !y {
				c.AbortWithStatusJSON(http.StatusUnauthorized, util.NO_AUTH_ERROR)
				return
			} else {
				roleName := c.GetString("role_name")
				role := c.GetInt("role")
				if !util.ExistIn(roleName, rolename) {
					err, permission := permissionServiceImpl.FindPermissionByPath(reqUrl[1])
					if err != nil {
						c.AbortWithStatusJSON(http.StatusAccepted, util.INSUFFICIENT_PERMISSION_ERROR)
						return
					}
					allowRole := permission.AuthorizedRoles
					roleList := strings.Split(allowRole, ",")
					roleExist := util.ExistIn(string(role), roleList)
					if !roleExist {
						//c.Abort()
						//fmt.Println("请求地址不包含该权限权限")
						c.AbortWithStatusJSON(http.StatusAccepted, util.INSUFFICENT_PERMISSION)
						//res.Err(util.INSUFFICENT_PERMISSION)
						return
					}
				}
				fmt.Println("检测到是超级管理员，可以直接操作，不需要判断")
			}
		}
	}
}

// 全局日志中间件
func LoggerMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		requestBody, _ := c.GetRawData()
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))
		rbody := string(requestBody)
		query := c.Request.URL.RawQuery
		c.Next() // 调用该请求的剩余处理程序
		stopTime := time.Since(startTime)
		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds()/1000000))))
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		method := c.Request.Method
		url := c.Request.RequestURI
		Log := logger.Logger.WithFields(
			logrus.Fields{
				"SpendTime": spendTime,       //接口花费时间
				"path":      url,             //请求路径
				"Method":    method,          //请求方法
				"status":    statusCode,      //接口返回状态
				"proto":     c.Request.Proto, //http请求版本
				"Ip":        clientIP,        //IP地址
				"body":      rbody,           //请求体
				"query":     query,           //请求query
				"message":   c.Errors,        //返回错误信息
			})
		if len(c.Errors) > 0 { // 矿建内部错误
			Log.Error(c.Errors.ByType(gin.ErrorTypePrivate))
		}
		if statusCode > 200 {
			Log.Error()
		} else {
			Log.Info()
		}
	}
}
