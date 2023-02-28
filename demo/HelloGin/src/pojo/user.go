package pojo

import "gorm.io/gorm"

//var db = global.Db

type User struct {
	gorm.Model
	Name        string `json:"name" gorm:"default:隔壁老王"`
	Password    string `json:"password"`
	Salt        string `json:"salt"`
	Account     string `json:"account" tag:"unique"`
	AccessToken string `json:"access_token"`
	Revoke      bool   `json:"revoke" gorm:"default:false"`
	Role        int    `json:"role" gorm:"default:5;type:int"`
}

//type UserInterface interface {
//	FindByAccount(string) (User, error)
//	FindByName(string) (User, error)
//}
//
////var use User
//
//func (u User) FindByAccount(account string) (User, error) {
//
//	u.Account = account
//	db.Select("id,name,account,password,salt").Where("account=?", account).First(&u)
//	return u, nil
//}
//func (u User) FindByName(name string) (User, error) {
//	u.Name = name
//	db.Select("id,name,account,password,salt").Where("name=?", name).First(&u)
//	return u, nil
//}

//func (u User) LoginByAccountPws(account string, password string) (s chan bool) {
//	db.Select("id,name,account,password,salt").Where("account =?", account).First(User{})
//	if u.ID == 0 {
//		s <- false
//		return
//	}
//	unenc, error := util.DePwdCode(u.Password, u.Salt)
//	if error != nil {
//		s <- false
//		return
//	}
//	if password != unenc {
//		s <- false
//		return
//	}
//	s <- true
//	return
//}
