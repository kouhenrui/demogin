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
		//userGroup.POST("/post/message", postMessage)
		//userGroup.PUT("/put/user", updateUser)
		adminGroup.POST("/register", registerAdmin)
	}
}

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
	result := adminService.AdminLogin(js)
	res.Success(result)
	return
}
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
func registerAdmin(c *gin.Context) {
	fmt.Println("进入控制层")
	res := global.NewResult(c)
	var re reqDto.AddAdmin
	fmt.Println(&re, "打印请求书数据")
	if err := c.BindJSON(&re); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			res.Error(http.StatusBadRequest, err.Error())
			return
		}
		res.DiyErr(http.StatusBadRequest, global.Translate(errs))
		return
	}
	result := adminService.AdminAdd(re)
	fmt.Println(result)
	res.Success(result)
	return
}
