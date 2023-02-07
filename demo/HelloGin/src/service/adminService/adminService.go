package adminService

import (
	"HelloGin/src/dto/reqDto"
	"HelloGin/src/dto/resDto"
	"HelloGin/src/global"
	"HelloGin/src/pojo"
	"HelloGin/src/util"
	"errors"
	"fmt"
	"reflect"
)

type AdminService struct {
}

//func NewAdminService() AdminService {
//	return AdminService{}
//}

var db = global.Db
var admin pojo.Admin
var b bool
var admins []pojo.Admin

func FindByAccount(account string) (pojo.Admin, bool) {
	r := db.Select("id,name,account,password,salt,access_token,role").Where("account=?", account).First(&admin)
	fmt.Println(admin, "影响行数")
	if r.RowsAffected != 1 {
		return admin, false
	}
	return admin, true
}
func FindByName(name string) (pojo.Admin, bool) {
	s := db.Where("name = ?", name).Find(&admin)
	if s.RowsAffected != 1 {
		return admin, false
	}
	return admin, true
}
func UpdateAdminToken(token string, id uint) bool {
	admin.ID = id
	res := db.Model(&admin).Update("access_token", token)
	if res.RowsAffected != 1 {
		return false
	}
	return true
}
func registerAdmin(addAdmin reqDto.AddAdmin) (pojo.Admin, bool) {
	admin.Name = addAdmin.Name
	admin.Password = addAdmin.Password
	admin.Salt = addAdmin.Salt
	admin.Account = addAdmin.Account
	admin.Role = 4
	fmt.Println(admin, "打印插入数据")
	res := db.Create(&admin)
	fmt.Println(admin)
	if res.RowsAffected != 1 {
		errors.New(util.ACCOUNT_EXIST_ERROR)
		//error()

		return admin, false
	}

	return admin, true
}

/**
反射控制层登录参数，查询数据库账号是否相同，
比对密码一致性，将用户信息存入jwt令牌中，签发令牌和过期时间
*/
func AdminLogin(loginReq interface{}) (tokenAndExp interface{}) {
	lp := reflect.ValueOf(loginReq)
	account := lp.FieldByName("Account").String()
	password := lp.FieldByName("Password").String()
	Revoke := lp.FieldByName("Revoke").Bool()
	admin, b = FindByAccount(account)
	if !b {
		return util.ACCOUT_NOT_EXIST_ERROR
	}
	salt := admin.Salt
	pwd, deerr := util.DePwdCode(admin.Password, salt)
	fmt.Println(pwd, deerr, "密码加密")
	if deerr != nil {
		return util.PASSWORD_RESOLUTION_ERROR
	}
	if pwd != password {
		return util.AUTH_LOGIN_PASSWORD_ERROR
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
	switch Revoke {
	case true:
		if existOldToken {
			util.DelRedis(admin.AccessToken) //清除token
		}
		token, exptime = util.SignToken(stringTokenData, global.UserLoginTime*global.DayTime)
		ok := UpdateAdminToken(tokenKey, admin.ID)
		if !ok {
			return util.AUTH_LOGIN_ERROR
		}
		redisDate := reqDto.LoginRedisDate{
			Token: token,
			Time:  exptime,
		}
		util.SetRedis(tokenKey, util.Marshal(redisDate), global.UserLoginTime)
		tokenAndExp = map[string]string{
			"token": token,
			"exp":   exptime,
		}
	case false:
		if existOldToken {
			tokenValue := util.GetRedis(admin.AccessToken)
			tokenAndExp = tokenValue
		}
		token, exptime = util.SignToken(stringTokenData, global.UserLoginTime*global.DayTime)
		ok := UpdateAdminToken(tokenKey, admin.ID)
		if !ok {
			return util.AUTH_LOGIN_ERROR
		}
		redisDate := reqDto.LoginRedisDate{
			Token: token,
			Time:  exptime,
		}
		util.SetRedis(tokenKey, util.Marshal(redisDate), global.UserLoginTime)
		tokenAndExp = map[string]string{
			"token": token,
			"exp":   exptime,
		}
	}
	return
}
func AdminInfo(id uint) pojo.Admin {
	db.Select("id,name,account,access_token,role").Where("id=?", id).First(&admin)
	return admin
}
func AdminAdd(req interface{}) interface{} {
	lp := reflect.ValueOf(req)
	name := lp.FieldByName("Name").String()
	account := lp.FieldByName("Account").String()
	password := lp.FieldByName("Password").String()
	salt := util.RandAllString()
	if password == "" {
		password = string(123456)
	}
	fmt.Println(password)
	enPwd, _ := util.EnPwdCode(password, salt)
	fmt.Println(name)
	if name != "" {
		_, tName := FindByName(name)
		if tName {
			return util.NAME_EXIST_ERROR
		}
	}
	if name == "" {
		name = "暂未命名"
	}
	_, tAccount := FindByAccount(account)
	if tAccount {
		return util.ACCOUNT_EXIST_ERROR
	}
	addAdmin := reqDto.AddAdmin{
		Name:     name,
		Account:  account,
		Password: enPwd,
		Salt:     salt,
	}
	result, tAdmin := registerAdmin(addAdmin)
	if tAdmin {
		return result
	} else {
		return util.ADD_ERROR
	}

}
func AdminList(list interface{}) interface{} {
	lp := reflect.ValueOf(list)
	take := lp.FieldByName("Take").Int()
	skip := lp.FieldByName("Skip").Uint()
	name := lp.FieldByName("Name").String()
	t1 := lp.FieldByName("Name")
	fmt.Println(name)
	fmt.Println(lp.FieldByName("Name").IsValid(), t1, "判断是否合法") //判断值是否有效
	var adminLists []resDto.AdminList
	query := db.Model(&pojo.Admin{})
	if name != "" {
		query.Where("name like ?", "%"+name+"%")
	}
	//query.Where("name like ?", "%"+name+"%")
	query.Limit(int(take)).Offset(int(skip)).Find(&adminLists)
	count := query.RowsAffected
	resList := resDto.CommonList{}
	resList.List = adminLists
	resList.Count = uint(count)

	fmt.Println(resList)
	return resList
}
