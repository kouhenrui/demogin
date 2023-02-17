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

var judje bool

func Routers(e *gin.Engine) {

	adminGroup := e.Group("/api/admin")
	{
		adminGroup.POST("/login", adminLogin)
		adminGroup.GET("/info", getAdminInfo)
		//adminGroup.POST("/register", registerAdmin)
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
		res.DiyErr(http.StatusBadRequest, global.Translate(errs))
		return
	}
	jude, result := adminService.AdminLogin(js)
	if jude {
		res.Err(result)
		return
	}
	res.Success(result)
	return
}

//获取详情接口
func getAdminInfo(c *gin.Context) {
	res := global.NewResult(c)
	user, _ := c.Get("user")
	//info:=adminService.AdminInfo()
	fmt.Println("request", user)
	res.Success(gin.H{
		"message": "hello gin",
		"request": user,
	})
	return
}

//增加管理员接口
//func registerAdmin(c *gin.Context) {
//	fmt.Println("进入控制层")
//	res := global.NewResult(c)
//	var re reqDto.AddAdmin
//	fmt.Println(&re, "打印请求书数据")
//	if err := c.BindJSON(&re); err != nil {
//		errs, ok := err.(validator.ValidationErrors)
//		if !ok {
//			res.Error(http.StatusBadRequest, err.Error())
//			return
//		}
//		res.DiyErr(http.StatusBadRequest, global.Translate(errs))
//		return
//	}
//	result := adminService.AdminAdd(re)
//	fmt.Println(result)
//	res.Success(result)
//	return
//}

//管理员列表
func adminList(c *gin.Context) {
	res := global.NewResult(c)
	var ls reqDto.AdminList
	//fmt.Println(ls, "请求参数")
	if err := c.BindJSON(&ls); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			res.Error(http.StatusBadRequest, err.Error())
			return
		}
		res.DiyErr(http.StatusBadRequest, global.Translate(errs))
		return
	}
	//fmt.Println(ls)
	list := adminService.AdminList(ls)

	res.Success(list)
	return

}

func userList(c *gin.Context) {
	res := global.NewResult(c)
	var ls reqDto.UserList
	if err := c.BindJSON(&ls); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			res.Error(http.StatusBadRequest, err.Error())
			return
		}
		res.DiyErr(http.StatusBadRequest, global.Translate(errs))
		return
	}
	list := adminService.UserList(ls)
	res.Success(list)
	return
}
