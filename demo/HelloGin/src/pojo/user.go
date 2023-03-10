package pojo

import (
	"HelloGin/src/dto/reqDto"
	"HelloGin/src/dto/resDto"
	"fmt"
	"gorm.io/gorm"
)

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

func UserServiceImpl() User {
	return User{}
}

var (
	userList     = &[]User{} //多个user返回
	user         = User{}
	resUsersList = []resDto.UserList{} //要查询的字段
)

// 分页,模糊查询用户
func (u *User) UserList(list reqDto.UserList) resDto.CommonList {
	query := db.Model(user)
	if list.Name != "" {
		query.Where("name like ?", "%"+list.Name+"%")
	}
	query.Limit(list.Take).Offset(int(list.Skip)).Find(&resUsersList)
	reslist.Count = uint(query.RowsAffected)
	reslist.List = resUsersList
	return reslist
}

// 查询账号
func (u *User) CheckByAccount(account string) (User, bool) {
	res := db.Model(&u).First(&user).Where("account =?", account)
	if res.RowsAffected <= 0 {
		return user, false
	}
	return user, true
}

// 查询名称
func (u *User) CheckByName(name string) (User, bool) {
	res := db.Model(&user).First(&user).Where("name =?", name)
	if res.RowsAffected <= 0 {
		return user, false
	}
	return user, true
}

// 更新token数据
func (u *User) UpdateToken(access_token string, id uint) bool {
	user.ID = id
	res := db.Model(&user).Update("access_token", access_token)
	if res.RowsAffected != 1 {
		return false
	}
	return true
}

// 增加用户
func (u *User) AddUser(user User) bool {
	fmt.Println(user)
	res := db.Create(&user)
	if res.RowsAffected != 1 {
		return false
	}
	return true
}
