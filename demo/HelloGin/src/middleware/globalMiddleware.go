package middleWare

import (
	"HelloGin/src/global"
	"HelloGin/src/pojo"
	"HelloGin/src/util"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

var userInfo util.UserClaims
var permissionServiceImpl = pojo.RbacPermission()

func GolbalMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		//res := global.NewResult(c)
		fmt.Println("身份认证开始执行")
		t := time.Now()
		requestUrl := c.Request.URL.String()
		reqUrl := strings.Split(requestUrl, "/api/")
		paths := global.ReuqestPaths
		pathIsExist := util.ExistIn(reqUrl[1], paths)
		if !pathIsExist {
			fmt.Println("身份验证")
			judge := util.AnalysyToken(c)
			if !judge {
				//res.Err(util.NO_AUTHORIZATION)
				c.Abort()
				//return

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
		//res := global.NewResult(c)

		requestUrl := c.Request.URL.String()
		reqUrl := strings.Split(requestUrl, "/api/")
		k, y := c.Get("ok")
		fmt.Println(k, y)
		//tokrn:=c.GetHeader("Authorization")
		//paths := global.ReuqestPaths
		//pathIsExist := util.ExistIn(reqUrl[1], paths)
		if !y {
			errors.New(util.NO_AUTH_ERROR)
			c.Abort()
			//res.Err(util.NO_AUTH_ERROR)
		} else {
			role_name, ok := c.Get("role_name")
			role, _ := c.Get("role")
			fmt.Println(role_name, "用户数据")
			fmt.Println(ok, "ok")
			if ok {
				if role_name == "admin" {
					c.Next()
				}
				t, permission := permissionServiceImpl.FindPermissionByPath(reqUrl[1])
				fmt.Println(t)
				if !t {
					//res.Err(util.INSUFFICIENT_PERMISSION_ERROR)
					c.Abort()
					//return
				}

				allowRole := permission.AuthorizedRoles
				roleList := strings.Split(allowRole, ",")
				fmt.Println("allowRole:", roleList[1])
				//fmt.Fprintf("allowrole is %T",allowRole)
				fmt.Println("role:", role)
				//if allowRole != role {
				//	res.Err(util.INSUFFICENT_PERMISSION)
				//	c.Abort()
				//	return
				//}
			}
		}

		//fmt.Println("身份认证执行结束")
	}
}
