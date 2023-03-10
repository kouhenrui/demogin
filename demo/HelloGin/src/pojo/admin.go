package pojo

import (
	"HelloGin/src/dto/reqDto"
	"HelloGin/src/dto/resDto"
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Name        string `json:"name" gorm:"not null"`
	Password    string `json:"password" `
	Salt        string `json:"salt"`
	Account     string `json:"account" gorm:"unique:true"`
	AccessToken string `json:"access_token"`
	Revoke      bool   `json:"revoke" gorm:"default:false"`
	Role        int    `json:"role"`
}

func AdminServiceImpl() Admin {
	return Admin{}
}

var (
	//admins=[]Admin{}
	admin        = Admin{}
	resAdminList = []resDto.AdminList{} //要查询的字段
)

// 分页,模糊查询用户
func (a *Admin) AdminList(list reqDto.AdminList) resDto.CommonList {
	query := db.Model(&a)
	if list.Name != "" {
		query.Where("name like ?", "%"+list.Name+"%")
	}
	query.Limit(list.Take).Offset(int(list.Skip)).Find(&resAdminList)
	reslist.Count = uint(query.RowsAffected)
	reslist.List = resAdminList
	return reslist
}

// 查询账号
func (a *Admin) CheckByAccount(account string) (Admin, bool) {
	res := db.Where("account =?", account).Find(&admin)
	if res.RowsAffected != 1 {
		return admin, false
	}
	return admin, true
}

// 查询名称
func (a *Admin) CheckByName(name string) (Admin, bool) {
	res := db.Where("name =?", name).Find(&a)
	if res.RowsAffected != 1 {
		return admin, false
	}
	return admin, true
}

// 更新token数据
func (a *Admin) UpdateToken(access_token string, id uint) bool {
	a.ID = id
	res := db.Model(&a).Update("access_token", access_token)
	if res.RowsAffected != 1 {
		return false
	}
	return true
}

// 增加用户
func (a *Admin) AddAdmin(admins Admin) bool {
	//fmt.Println(admin)
	res := db.Create(&admins)
	if res.RowsAffected != 1 {
		return false
	}
	return true
}
