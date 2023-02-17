package user

import "C"
import (
	"HelloGin/src/dto/reqDto"
	"HelloGin/src/global"
	interf "HelloGin/src/interface"
	"HelloGin/src/pojo"
	"HelloGin/src/service/userService"
	"HelloGin/src/util"
	"fmt"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
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
		//userGroup.POST("/register/user", rejisterUser)
	}

}

var trans ut.Translator

// var userservice = userService.NewUserService()
var login = interf.LoginService()
var use pojo.User

//	func addUser(c *gin.Context) {
//		result := global.NewResult(c)
//		var add reqDto.AddUser
//		if err := c.ShouldBindJSON(&add); err != nil {
//			errs, ok := err.(validator.ValidationErrors)
//			if !ok {
//				result.DiyErr(http.StatusBadRequest, err.Error())
//				return
//			}
//			result.DiyErr(http.StatusBadRequest, global.Translate(errs))
//			return
//		}
//		//go func() {
//		//	log.Println("信道传输开始")
//		//	s := interf.Login(add.Account, add.Password)
//		//	//
//		//	a:=<-s
//		//	if <-s { //读取信道
//		//		result.DiyErr(http.StatusGatewayTimeout, util.NAME_AND_ACCOUNT_EXIST)
//		//		return
//		//	}
//		//	log.Println("信道传输结束")
//		//}()
//		//log.Println("信道传输未开始")
//		////生成16位随机字符串
//		//salt := util.RandAllString()
//		//pwd, pwderr := util.EnPwdCode(add.Password, salt)
//		//if pwderr != nil {
//		//	log.Println(pwderr, "打印错误")
//		//	result.DiyErr(http.StatusBadRequest, pwderr)
//		//	return
//		//}
//
//		//go func() {
//		//	b := userservice.AddUser(user)
//		//	log.Println(<-b, "传输掺入信道输出参数true false")
//		//	if <-b {
//		//		result.DiyErr(http.StatusGatewayTimeout, util.NAME_AND_ACCOUNT_EXIST)
//		//		return
//		//	}
//		//}()
//
//		token := util.SignToken()
//		log.Println(token, "token")
//		result.Success(gin.H{"token": token})
//		return
//	}
//
//	func rejisterUser(c *gin.Context) {
//		log.Println("注册请求接收到")
//		res := global.NewResult(c)
//		var a reqDto.AddUser
//		if err := c.BindJSON(&a); err != nil {
//			errs, ok := err.(validator.ValidationErrors)
//			if !ok {
//				res.Error(http.StatusInternalServerError, err.Error())
//				return
//			}
//			res.Error(http.StatusBadRequest, global.Translate(errs))
//			return
//		}
//		salt := util.RandAllString()
//		sa, _ := util.EnPwdCode(a.Password, salt)
//		a.Password = sa
//		a.Salt = salt
//		bol, msg := userservice.AddUser(a.Name, a.Account, a.Password, a.Salt)
//		if !bol {
//			res.Error(http.StatusBadRequest, msg)
//			return
//		}
//		res.Success(util.SUCCESS)
//		return
//	}
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
