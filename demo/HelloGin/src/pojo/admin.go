package pojo

import "gorm.io/gorm"

//var db = global.Db

type Admin struct {
	gorm.Model
	//Base        Base   `gorm:"embedded"`
	Name        string `json:"name" gorm:"default:admin;unique:true"`
	Password    string `json:"password" `
	Salt        string `json:"salt"`
	Account     string `json:"account" gorm:"unique:true"`
	AccessToken string `json:"access_token"`
	Revoke      bool   `json:"revoke" gorm:"default:false"`
	Role        int    `json:"role" gorm:"type:not null" `
}

type AdminInterface interface {
	FindByAccount(string) (User, error)
	FindByName(string) (User, error)
}

var admin Admin

func (a Admin) FindByAccount(account string) (Admin, error) {
	//a.Account = account

	db.Select("id,name,account,password,salt").Where("account=?", account).First(&admin)
	return admin, nil
}
func (a Admin) FindByName(name string) (Admin, error) {
	db.Select("id,name,account,password,salt").Where("name=?", name).First(&admin)
	return admin, nil
}

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
