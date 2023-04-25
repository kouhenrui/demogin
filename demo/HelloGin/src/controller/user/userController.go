package user

import (
	"HelloGin/src/dto/reqDto"
	"HelloGin/src/global"
	"HelloGin/src/service/userService"
	"HelloGin/src/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func Routers(e *gin.Engine) {

	userGroup := e.Group("/api/user")
	{
		userGroup.POST("/login", userLogin)
		userGroup.GET("/info", getUserInfo)
		userGroup.POST("/post/message", postMessage)
		userGroup.PUT("/put/user", updateUser)
		userGroup.POST("/register/user", rejisterUser)
	}

}

func rejisterUser(c *gin.Context) {
	res := global.NewResult(c)
	var add reqDto.AddUser
	if err := c.BindJSON(&add); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			res.Error(http.StatusInternalServerError, err.Error())
			return
		}
		res.Error(http.StatusBadRequest, global.Translate(errs))
		return
	}
	bol, msg := userService.UserRejist(add)
	if !bol {
		res.Err(msg)
		return
	}
	res.Success(msg)
	return
}
func userLogin(c *gin.Context) {
	res := global.NewResult(c)
	var js reqDto.UserLogin
	if err := c.BindJSON(&js); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			res.Error(http.StatusBadRequest, err.Error())
			return
		}
		res.Error(http.StatusBadRequest, global.Translate(errs))
		return
	}
	bol, msg := userService.UserLogin(js)
	if !bol {
		res.Err(msg)
		return
	}
	res.Success(msg)
	return

}
func postMessage(c *gin.Context) {
	types := c.DefaultPostForm("type", "post")
	name := c.PostForm("name")
	pwd := c.PostForm("pwd")
	fmt.Println(name, pwd, "传递参数")
	c.String(http.StatusOK, fmt.Sprintf("name:%s ,pwd:%s,type:%s", name, pwd, types))
}

func getUserInfo(c *gin.Context) {
	res := global.NewResult(c)
	user, _ := c.Get("user")
	fmt.Println("request", user)
	res.Success(gin.H{
		"message": "hello gin",
		"request": user,
	})
	return

}

func updateUser(c *gin.Context) {
	result := global.NewResult(c)

	result.Success(util.MODIFICATION_SUCCESSE)
	return
}
