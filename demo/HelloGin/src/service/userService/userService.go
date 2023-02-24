package userService

import (
	"HelloGin/src/dto/reqDto"
	"HelloGin/src/dto/resDto"
	"HelloGin/src/global"
	userDao "HelloGin/src/interface/user"
	"HelloGin/src/pojo"
	"HelloGin/src/util"
	"fmt"
)

type UserService struct {
}

func NewUserService() UserService {
	return UserService{}
}

var db = global.Db
var user pojo.User
var userServiceImpl = userDao.UserServiceImpl()
var judge bool

func (c *UserService) FindByAccount(account string) (pojo.User, bool) {
	db.Select("id,name,account,password,salt,access_token,role").Where("account=?", account).First(&user)
	return user, true
}
func (c *UserService) FindByName(name string) (pojo.User, bool) {
	//user.Name = name
	db.Select("id,name,account,password,salt,access_token,role").Where("name=?", name).First(&user)
	return user, true
}
func (c *UserService) AddUser(name string, account string, password string, salt string) (bool, string) {
	user.Name = name
	user.Salt = salt
	user.Password = password
	user.Account = account
	res := db.Create(&user)
	if res.RowsAffected == 0 {
		return false, util.INSET_USER_ERROR
	}
	return true, util.SUCCESS
}
func (c *UserService) UpdateUserToken(accessToken string, id uint) bool {
	user.ID = id
	res := db.Model(&user).Update("AccessToken", accessToken)
	if res.RowsAffected != 1 {
		return false
	}
	return true
}

// 用户列表
func UserList(list reqDto.UserList) interface{} {
	res := userServiceImpl.UserList(list)
	fmt.Println(res, "service返回的结果")
	return res
}

func UserLogin(list reqDto.UserLogin) (a bool, tokenAndExp interface{}) {
	switch list.Method {
	case "name":
		user, judge = userServiceImpl.CheckByName(list.Name)
	case "account":
		user, judge = userServiceImpl.CheckByAccount(list.Account)
	default:
		return false, util.METHOD_NOT_FILLED_ERROR
	}
	if judge {
		return false, util.ACCOUT_NOT_EXIST_ERROR
	}
	enpwd := user.Password
	salt := user.Salt
	pwd, deerr := util.DePwdCode(enpwd, salt)
	if deerr != nil {
		fmt.Println(deerr, "加密方法")
		//errors.New(util.PASSWORD_RESOLUTION_ERROR)
		return false, util.PASSWORD_RESOLUTION_ERROR
	}
	if pwd != list.Password {
		return false, util.AUTH_LOGIN_PASSWORD_ERROR
	}
	existOldToken := util.ExistRedis(user.AccessToken)
	tokenKey := util.Rand6String6()
	var token string
	var exptime string
	stringTokenData := util.UserClaims{
		Id:      user.ID,
		Name:    user.Name,
		Account: user.Account,
		Role:    user.Role,
	}
	switch list.Revoke {
	case true:
		if existOldToken {
			util.DelRedis(user.AccessToken) //清除token
		}
		token, exptime = util.SignToken(stringTokenData, global.UserLoginTime*global.DayTime)
		ok := userServiceImpl.UpdateToken(tokenKey, user.ID)
		if !ok {
			return false, util.AUTH_LOGIN_ERROR
		}
		redisDate := reqDto.LoginRedisDate{
			Token: token,
			Time:  exptime,
		}
		util.SetRedis(tokenKey, util.Marshal(redisDate), global.UserLoginTime)
		tokenAndExp = resDto.TokenAndExp{
			token,
			exptime,
		}
	case false:
		if existOldToken {
			tokenValue := util.GetRedis(user.AccessToken)
			mp := make(map[string]interface{})
			_, cs := util.UnMarshal([]byte(tokenValue), &mp)
			return true, cs
		}
		token, exptime = util.SignToken(stringTokenData, global.UserLoginTime*global.DayTime)
		ok := userServiceImpl.UpdateToken(tokenKey, user.ID)
		if !ok {
			return false, util.AUTH_LOGIN_ERROR
		}
		redisDate := reqDto.LoginRedisDate{
			Token: token,
			Time:  exptime,
		}
		util.SetRedis(tokenKey, util.Marshal(redisDate), global.UserLoginTime)
		tokenAndExp = resDto.TokenAndExp{
			token,
			exptime,
		}
	}
	return true, tokenAndExp
}

//func UserByNameAndAccount(query string) bool {
//	result := db.Where(query).Take(&u)
//	if result.Error != nil {
//		return false
//	}
//	return true
//}
//
////func JudgeUserExist(name string, account string) (a chan bool) {
////	//ua := pojo.User{}
////	db.Select("id", "name", "account").Where("name = ? or account=?", name, account).First(&u)
////	log.Println(u, "打印数据公告", u.ID, "查询id")
////	if u.ID == 0 {
////		a <- true
////		return
////	}
////	a <- false
////	return
////}
////
