package interf

import (
	"HelloGin/src/dto/reqDto"
	"HelloGin/src/pojo"
	"HelloGin/src/service/userService"
	"HelloGin/src/util"
	"fmt"
	"reflect"
)

var user = userService.NewUserService()

type Login struct{}

func LoginService() Login {
	return Login{}
}

var u pojo.User
var ue bool

func (l *Login) UserLogin(i interface{}) (bool, interface{}) {
	lp := reflect.ValueOf(i)
	method := lp.FieldByName("Method").String()
	name := lp.FieldByName("Name").String()
	account := lp.FieldByName("Account").String()
	password := lp.FieldByName("Password").String()
	switch method {
	case "name":
		u, ue = user.FindByName(name)
	case "account":
		u, ue = user.FindByAccount(account)
	default:
		return false, util.METHOD_NOT_FILLED_ERROR
	}
	if !ue {
		return false, util.ACCOUT_NOT_EXIST_ERROR
	}
	enpwd := u.Password
	salt := u.Salt
	fmt.Println(salt)
	pwd, deerr := util.DePwdCode(enpwd, salt)
	if deerr != nil {
		fmt.Println(deerr, "加密方法")
		return false, util.PASSWORD_RESOLUTION_ERROR
	}
	fmt.Println(pwd, password)
	if pwd != password {
		return false, util.AUTH_LOGIN_PASSWORD_ERROR
	}
	tokenData := reqDto.TokenDate{
		Id:          u.ID,
		Name:        u.Name,
		Account:     u.Account,
		Salt:        u.Salt,
		AccessToken: u.AccessToken,
		Role:        u.Role,
	}
	return true, tokenData
}
func (l *Login) UserInfo(account string) interface{} {
	u, _ := user.FindByAccount(account)
	return u
}
