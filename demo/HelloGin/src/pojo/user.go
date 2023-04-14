package pojo

import (
	"HelloGin/src/dto/reqDto"
	"HelloGin/src/dto/resDto"
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
func (u *User) UserList(list reqDto.UserList) *resDto.CommonList {
	query := db.Model(user)
	if list.Name != "" {
		query.Where("name like ?", "%"+list.Name+"%")
	}
	query.Limit(list.Take).Offset(int(list.Skip)).Find(&resUsersList)
	reslist.Count = uint(query.RowsAffected)
	reslist.List = resUsersList
	return &reslist
}

// 查询账号
func (u *User) CheckByAccount(account string) (error, *resDto.UserInfo) {
	var userInfo resDto.UserInfo
	userInfo.Account = account
	err := db.Model(&u).First(&user).Error
	if err != nil {
		return err, nil
	}
	return nil, &userInfo
}

// 查询名称
func (u *User) CheckByName(name string) (error, *resDto.UserInfo) {
	var userInfo resDto.UserInfo
	userInfo.Name = name
	err := db.Model(&user).First(&userInfo).Error
	if err != nil {
		return err, nil
	}
	return nil, &userInfo
}

// 更新token数据
func (u *User) UpdateToken(access_token string, id uint) error {

	u.ID = id
	return db.Model(&u).Update("access_token", access_token).Error
}

// 增加用户
func (u *User) AddUser(user User) error {
	return db.Create(&user).Error
}
