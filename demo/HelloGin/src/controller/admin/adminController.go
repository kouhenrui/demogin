package admin

import (
	"HelloGin/src/dto/reqDto"
	"HelloGin/src/global"
	"HelloGin/src/service/adminService"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func Routers(e *gin.Engine) {
	adminGroup := e.Group("/api/admin")
	{
		adminGroup.POST("/login", adminLogin)
		adminGroup.GET("/info", getAdminInfo)
		adminGroup.POST("/register", registerAdmin)
		adminGroup.POST("/list", adminList)
		adminGroup.POST("/users/list", userList)
	}
}

// 登录接口
func adminLogin(c *gin.Context) {
	res := global.NewResult(c)
	var js reqDto.AdminLogin
	if err := c.BindJSON(&js); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			res.Error(http.StatusBadRequest, err.Error())
			return
		}
		res.Error(http.StatusBadRequest, global.Translate(errs))
		return
	}
	jude, result := adminService.AdminLogin(js)
	if !jude {
		res.Err(result)
		return
	}
	res.Success(result)
	return
}

// 获取详情接口
func getAdminInfo(c *gin.Context) {
	res := global.NewResult(c)

	//var u=&util.UserClaims{}
	user, _ := c.Get("user")
	fmt.Println("request", user)
	res.Success(user)
	//result := adminService.AdminInfo()
	//res.Success(result)
	return
}

// 增加管理员接口
func registerAdmin(c *gin.Context) {
	fmt.Println("进入控制层")
	res := global.NewResult(c)
	var re reqDto.AddAdmin
	if err := c.BindJSON(&re); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			res.Error(http.StatusBadRequest, err.Error())
			return
		}
		res.Error(http.StatusBadRequest, global.Translate(errs))
		return
	}
	judje, result := adminService.AdminAdd(re)
	if judje {
		res.Success(result)
		return
	} else {
		res.Err(result)
		return
	}
}

// 管理员列表
func adminList(c *gin.Context) {
	res := global.NewResult(c)
	var ls reqDto.AdminList
	if err := c.BindJSON(&ls); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			res.Error(http.StatusBadRequest, err.Error())
			return
		}
		res.Error(http.StatusBadRequest, global.Translate(errs))
		return
	}
	list := adminService.AdminList(ls)
	res.Success(list)
	return

}

// 用户列表
func userList(c *gin.Context) {
	res := global.NewResult(c)
	var ls reqDto.UserList
	if err := c.BindJSON(&ls); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			res.Error(http.StatusBadRequest, err.Error())
			return
		}
		res.Error(http.StatusBadRequest, global.Translate(errs))
		return
	}
	list := adminService.UserList(ls)
	res.Success(list)
	return
}
