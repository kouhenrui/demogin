package user

import "C"
import (
	"HelloGin/src/dto/reqDto"
	"HelloGin/src/global"
	interf "HelloGin/src/interface"
	"HelloGin/src/pojo"
	"HelloGin/src/service/userService"
	"HelloGin/src/util"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"reflect"
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

var trans ut.Translator

var userservice = userService.NewUserService()
var login = interf.LoginService()
var use pojo.User

//func addUser(c *gin.Context) {
//	result := global.NewResult(c)
//	var add reqDto.AddUser
//	if err := c.ShouldBindJSON(&add); err != nil {
//		errs, ok := err.(validator.ValidationErrors)
//		if !ok {
//			result.DiyErr(http.StatusBadRequest, err.Error())
//			return
//		}
//		result.DiyErr(http.StatusBadRequest, global.Translate(errs))
//		return
//	}
//	//go func() {
//	//	log.Println("信道传输开始")
//	//	s := interf.Login(add.Account, add.Password)
//	//	//
//	//	a:=<-s
//	//	if <-s { //读取信道
//	//		result.DiyErr(http.StatusGatewayTimeout, util.NAME_AND_ACCOUNT_EXIST)
//	//		return
//	//	}
//	//	log.Println("信道传输结束")
//	//}()
//	//log.Println("信道传输未开始")
//	////生成16位随机字符串
//	//salt := util.RandAllString()
//	//pwd, pwderr := util.EnPwdCode(add.Password, salt)
//	//if pwderr != nil {
//	//	log.Println(pwderr, "打印错误")
//	//	result.DiyErr(http.StatusBadRequest, pwderr)
//	//	return
//	//}
//
//	//go func() {
//	//	b := userservice.AddUser(user)
//	//	log.Println(<-b, "传输掺入信道输出参数true false")
//	//	if <-b {
//	//		result.DiyErr(http.StatusGatewayTimeout, util.NAME_AND_ACCOUNT_EXIST)
//	//		return
//	//	}
//	//}()
//
//	token := util.SignToken()
//	log.Println(token, "token")
//	result.Success(gin.H{"token": token})
//	return
//}
func rejisterUser(c *gin.Context) {
	log.Println("注册请求接收到")
	res := global.NewResult(c)
	var a reqDto.AddUser
	if err := c.BindJSON(&a); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			res.Error(http.StatusInternalServerError, err.Error())
			return
		}
		res.Error(http.StatusBadRequest, global.Translate(errs))
		return
	}
	salt := util.RandAllString()
	sa, _ := util.EnPwdCode(a.Password, salt)
	a.Password = sa
	a.Salt = salt
	bol, msg := userservice.AddUser(a.Name, a.Account, a.Password, a.Salt)
	if !bol {
		res.Error(http.StatusBadRequest, msg)
		return
	}
	res.Success(util.SUCCESS)
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
	bol, msg := login.UserLogin(js)
	if !bol {
		res.Error(http.StatusBadRequest, msg)
		return
	}
	tokenDate := reflect.ValueOf(msg)
	old_access := tokenDate.FieldByName("AccessToken").String()
	id := tokenDate.FieldByName("Id").Uint()
	a := util.ExistRedis(old_access)
	access_token := util.Rand6String6()
	stringTokenData := util.UserClaims{
		Id:      uint(tokenDate.FieldByName("Id").Uint()),
		Name:    tokenDate.FieldByName("Name").String(),
		Account: tokenDate.FieldByName("Account").String(),
		Role:    int(tokenDate.FieldByName("Role").Int()),
	}
	var token string
	var exptime string
	switch js.Revoke {
	case true:
		if a {
			util.DelRedis(old_access) //删除老token
		} //强制删除现存登陆信息
		ok := userservice.UpdateUserToken(access_token, uint(id))
		if !ok {
			res.Error(http.StatusInternalServerError, util.AUTH_LOGIN_ERROR)
			return
		}
		token, exptime = util.SignToken(stringTokenData, global.UserLoginTime*global.DayTime)
		redisDate := reqDto.LoginRedisDate{
			Token: token,
			Time:  exptime,
		}
		util.SetRedis(access_token, util.Marshal(redisDate), global.UserLoginTime)
	case false:
		if a {
			d := util.GetRedis(old_access)
			var red reqDto.LoginRedisDate
			er := json.Unmarshal([]byte(d), &red)
			fmt.Print(er)
			if er != nil {
				res.Error(http.StatusInternalServerError, util.JSON_UNMARSHAL_ERROR)
				return
			}
			rf := reflect.ValueOf(red)
			token = rf.FieldByName("Token").String()
			exptime = rf.FieldByName("Time").String()
		} else {
			ok := userservice.UpdateUserToken(access_token, uint(id))
			if !ok {
				res.Error(http.StatusInternalServerError, util.AUTH_LOGIN_ERROR)
				return
			}
			token, exptime = util.SignToken(stringTokenData, global.UserLoginTime*global.DayTime)
			redisDate := reqDto.LoginRedisDate{
				Token: token,
				Time:  exptime,
			}
			if b := util.SetRedis(access_token, util.Marshal(redisDate), global.UserLoginTime); !b {
				log.Println(b, "redis返回值")
				res.Error(http.StatusInternalServerError, util.REDIS_INFORMATION_ERROR)
				return
			}
		}
	}
	res.Success(gin.H{"token": token, "exptime": exptime})
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
