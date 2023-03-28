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

var userInfo util.UserClaims
var permissionServiceImpl = pojo.RbacPermission()

func GolbalMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		res := global.NewResult(c)
		fmt.Println("身份认证开始执行")
		t := time.Now()
		requestUrl := c.Request.URL.String()
		reqUrl := strings.Split(requestUrl, "/api/")
		paths := global.ReuqestPaths
		pathIsExist := util.ExistIn(reqUrl[1], paths)
		if !pathIsExist {
			_, ok := c.Get("ok")
			fmt.Println("OK", ok)
			if ok {
				c.Next()
			}
			judge := util.AnalysyToken(c)
			if !judge {
				fmt.Println("身份验证未通过，没有token")
				c.Abort()
				res.Err(util.NO_AUTHORIZATION)
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
func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		res := global.NewResult(c)
		requestUrl := c.Request.URL.String()
		reqUrl := strings.Split(requestUrl, "/api/")
		rolename := global.RoleName
		paths := global.ReuqestPaths
		pathIsExist := util.ExistIn(reqUrl[1], paths)
		//登录跳过权限验证
		if pathIsExist {
			c.Next()
		}
		//验证身份
		_, y := c.Get("ok")
		//通过身份验证
		if !y {
			c.Abort()
			res.Err(util.NO_AUTH_ERROR)
			return
		} else {
			roleName := c.GetString("role_name")
			role := c.GetInt("role")
			if !util.ExistIn(roleName, rolename) {
				t, permission := permissionServiceImpl.FindPermissionByPath(reqUrl[1])
				if !t {
					c.Abort()
					fmt.Println("请求地址权限未找到")
					res.Err(util.INSUFFICIENT_PERMISSION_ERROR)
					return
				}
				allowRole := permission.AuthorizedRoles
				roleList := strings.Split(allowRole, ",")
				roleExist := util.ExistIn(string(role), roleList)
				if !roleExist {
					c.Abort()
					fmt.Println("请求地址不包含该权限权限")
					res.Err(util.INSUFFICENT_PERMISSION)
					return
				}
			}
			fmt.Println("检测到是超级管理员，可以直接操作，不需要判断")
			//c.Next()
		}

		//fmt.Println("身份认证执行结束")
	}
}
