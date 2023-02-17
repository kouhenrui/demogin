package adminService

import (
	"HelloGin/src/dto/reqDto"
	"HelloGin/src/dto/resDto"
	"HelloGin/src/global"
	adminDao "HelloGin/src/interface/admin"
	userDao "HelloGin/src/interface/user"
	"HelloGin/src/pojo"
	"HelloGin/src/util"
	"errors"
	"fmt"
)

type AdminService struct {
}

var admin pojo.Admin
var judge bool

//var b bool
//var admins []pojo.Admin
var userServiceImpl = userDao.UserServiceImpl()
var adminServiceImpl = adminDao.AdminServiceImpl()

//var userDao UserDao
//func FindByAccount(account string) (pojo.Admin, bool) {
//	r := db.Select("id,name,account,password,salt,access_token,role").Where("account=?", account).First(&admin)
//	fmt.Println(admin, "影响行数")
//	if r.RowsAffected != 1 {
//		return admin, false
//	}
//	return admin, true
//}
//func FindByName(name string) (pojo.Admin, bool) {
//	s := db.Where("name = ?", name).Find(&admin)
//	if s.RowsAffected != 1 {
//		return admin, false
//	}
//	return admin, true
//}
//func UpdateAdminToken(token string, id uint) bool {
//	admin.ID = id
//	res := db.Model(&admin).Update("access_token", token)
//	if res.RowsAffected != 1 {
//		return false
//	}
//	return true
//}
//func registerAdmin(addAdmin reqDto.AddAdmin) (pojo.Admin, bool) {
//	admin.Name = addAdmin.Name
//	admin.Password = addAdmin.Password
//	admin.Salt = addAdmin.Salt
//	admin.Account = addAdmin.Account
//	admin.Role = 4
//	fmt.Println(admin, "打印插入数据")
//	res := db.Create(&admin)
//	fmt.Println(admin)
//	if res.RowsAffected != 1 {
//		errors.New(util.ACCOUNT_EXIST_ERROR)
//		//error()
//
//		return admin, false
//	}
//
//	return admin, true
//}

/**
反射控制层登录参数，查询数据库账号是否相同，
比对密码一致性，将用户信息存入jwt令牌中，签发令牌和过期时间
*/
func AdminLogin(list reqDto.AdminLogin) (a bool, tokenAndExp interface{}) {
	switch list.Method {
	case "name":
		admin, judge = adminServiceImpl.CheckByName(list.Name)
	case "account":
		admin, judge = adminServiceImpl.CheckByAccount(list.Account)
	default:
		return false, util.METHOD_NOT_FILLED_ERROR
	}
	if judge {
		return false, util.ACCOUT_NOT_EXIST_ERROR
	}
	enpwd := admin.Password
	salt := admin.Salt
	pwd, deerr := util.DePwdCode(enpwd, salt)
	if deerr != nil {
		fmt.Println(deerr, "加密方法")
		errors.New(util.PASSWORD_RESOLUTION_ERROR)
		//return false, util.PASSWORD_RESOLUTION_ERROR
	}
	fmt.Println(pwd, list.Password, "比对密码")
	if pwd != list.Password {
		return false, util.AUTH_LOGIN_PASSWORD_ERROR
	}
	existOldToken := util.ExistRedis(admin.AccessToken)
	tokenKey := util.Rand6String6()
	var token string
	var exptime string
	stringTokenData := util.UserClaims{
		Id:      admin.ID,
		Name:    admin.Name,
		Account: admin.Account,
		Role:    admin.Role,
	}
	switch list.Revoke {
	case true:
		if existOldToken {
			util.DelRedis(admin.AccessToken) //清除token
		}
		token, exptime = util.SignToken(stringTokenData, global.UserLoginTime*global.DayTime)
		ok := adminServiceImpl.UpdateToken(tokenKey, admin.ID)
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
			tokenValue := util.GetRedis(admin.AccessToken)
			tokenAndExp = tokenValue
		}
		token, exptime = util.SignToken(stringTokenData, global.UserLoginTime*global.DayTime)
		ok := adminServiceImpl.UpdateToken(tokenKey, admin.ID)
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

//func AdminInfo(id uint) pojo.Admin {
//	db.Select("id,name,account,access_token,role").Where("id=?", id).First(&admin)
//	return admin
//}
//func AdminAdd(req interface{}) interface{} {
//	lp := reflect.ValueOf(req)
//	name := lp.FieldByName("Name").String()
//	account := lp.FieldByName("Account").String()
//	password := lp.FieldByName("Password").String()
//	salt := util.RandAllString()
//	if password == "" {
//		password = string(123456)
//	}
//	fmt.Println(password)
//	enPwd, _ := util.EnPwdCode(password, salt)
//	fmt.Println(name)
//	if name != "" {
//		_, tName := FindByName(name)
//		if tName {
//			return util.NAME_EXIST_ERROR
//		}
//	}
//	if name == "" {
//		name = "暂未命名"
//	}
//	_, tAccount := FindByAccount(account)
//	if tAccount {
//		return util.ACCOUNT_EXIST_ERROR
//	}
//	addAdmin := reqDto.AddAdmin{
//		Name:     name,
//		Account:  account,
//		Password: enPwd,
//		Salt:     salt,
//	}
//	result, tAdmin := registerAdmin(addAdmin)
//	if tAdmin {
//		return result
//	} else {
//		return util.ADD_ERROR
//	}
//
//}
func AdminList(list reqDto.AdminList) interface{} {
	//lp := reflect.ValueOf(list)
	//take := lp.FieldByName("Take").Int()
	//skip := lp.FieldByName("Skip").Uint()
	//name := lp.FieldByName("Name").String()
	//t1 := lp.FieldByName("Name")
	//fmt.Println(name)
	//fmt.Println(lp.FieldByName("Name").IsValid(), t1, "判断是否合法") //判断值是否有效
	res := adminServiceImpl.AdminList(list)
	return res
}

func UserList(list reqDto.UserList) interface{} {
	//fmt.Println(list)
	res := userServiceImpl.UserList(list)
	fmt.Println(res, "service返回的结果")
	return res
}
