package admin

import (
	"HelloGin/src/dto/reqDto"
	"HelloGin/src/global"
	"HelloGin/src/service/adminService"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

func Routers(e *gin.Engine) {
	adminGroup := e.Group("/api/admin")
	{
		adminGroup.POST("/register", registerAdmin)
		adminGroup.POST("/login", adminLogin)
		adminGroup.GET("/info", getAdminInfo)
		adminGroup.GET("/logout", logout)
		adminGroup.POST("/list", adminList)
		adminGroup.POST("/users/list", userList)
		adminGroup.POST("/permission/list", permissionList)
		adminGroup.POST("/permission/add", permissionAdd)
		adminGroup.PUT("/permission/update", permissionUpdate)
		adminGroup.DELETE("/permission/del", permissionDelete) //permission/del?id=8
		adminGroup.GET("/permission/info", permissionInfo)     //permission/info?id=
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

// 登出
func logout(c *gin.Context) {
	res := global.NewResult(c)
	go adminService.AdminLogout()
	res.Succ()
}

// 获取详情接口
func getAdminInfo(c *gin.Context) {
	res := global.NewResult(c)
	id := c.GetInt("id")
	name := c.GetString("name")
	p, info := adminService.AdminInfo(id, name)
	if p {
		res.Success(info)
		return
	}
	res.Err(info)
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
	fmt.Println("请求参数：", ls)
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
	fmt.Println("list:", list)
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

// 权限列表
func permissionList(c *gin.Context) {
	res := global.NewResult(c)
	var ls reqDto.PermissionList
	fmt.Println("请求体：", ls.Take, ls.Skip)
	if err := c.ShouldBindJSON(&ls); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		fmt.Println("参数验证出错：", errs)
		if !ok {
			res.Error(http.StatusBadRequest, err.Error())
			return
		}
		res.Error(http.StatusBadRequest, global.Translate(errs))
		return
	}
	list := adminService.PermissionList(ls)
	res.Success(list)
	return
}

// 权限增加
func permissionAdd(c *gin.Context) {
	res := global.NewResult(c)
	var ls reqDto.PermissionAdd
	fmt.Println(ls, &ls)
	if err := c.BindJSON(&ls); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			res.Error(http.StatusBadRequest, err.Error())
			//c.Abort()
			return
		}
		res.Error(http.StatusBadRequest, global.Translate(errs))
		//c.Abort()
		return
	}
	list := adminService.PermissionAdd(ls)
	res.Success(list)
	return
}

/* 权限修改*/
func permissionUpdate(c *gin.Context) {
	res := global.NewResult(c)
	var ls reqDto.PermissionUpdate
	//id := c.Query("id")
	//i, _ := strconv.Atoi(id)
	if err := c.BindJSON(&ls); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			res.Error(http.StatusBadRequest, err.Error())
			return
		}
		res.Error(http.StatusBadRequest, global.Translate(errs))
		return
	}
	list := adminService.PermissionUpdate(ls)
	res.Success(list)
	return
}

// 权限删除
func permissionDelete(c *gin.Context) {
	res := global.NewResult(c)
	id := c.Query("id")
	i, _ := strconv.Atoi(id)
	ok, result := adminService.Permissiondel(i)
	if ok {
		res.Success(result)
		return
	}
	res.Err(result)
	return
}

/*权限详情*/
func permissionInfo(c *gin.Context) {
	res := global.NewResult(c)
	id := c.Query("id")
	i, _ := strconv.Atoi(id) //string转int
	ok, result := adminService.PermissionIndo(i)
	if ok {
		res.Success(result)
		return
	}
	res.Err(result)
	return
}
