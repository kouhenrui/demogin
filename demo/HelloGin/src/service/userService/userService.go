package userService

import (
	"HelloGin/src/global"
	"HelloGin/src/pojo"
	"HelloGin/src/util"
)

type UserService struct {
}

func NewUserService() UserService {
	return UserService{}
}

var db = global.Db
var user pojo.User

func (c *UserService) FindByAccount(account string) (pojo.User, bool) {
	//user.Account = account
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
