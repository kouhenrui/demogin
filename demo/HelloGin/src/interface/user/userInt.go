package userDao

import (
	"HelloGin/src/dto/reqDto"
	"HelloGin/src/dto/resDto"
	"HelloGin/src/global"
	"HelloGin/src/pojo"
)

var (
	db           = global.Db           //引用全局的数据连接
	user         = pojo.User{}         //单个user返回
	userList     = &[]pojo.User{}      //多个user返回
	resUsersList = []resDto.UserList{} //要查询的字段
	reslist      = resDto.CommonList{} //返回的列表，包含数据和数量
)

type UserDao struct{}

func UserServiceImpl() UserDao {
	return UserDao{}
}

// 分页,模糊查询用户
func (u *UserDao) UserList(list reqDto.UserList) resDto.CommonList {
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
func (u *UserDao) CheckByAccount(account string) (pojo.User, bool) {
	res := db.Model(&user).First(&user).Where("account =?", account)
	if res.RowsAffected <= 0 {
		return user, false
	}
	return user, true
}

// 查询名称
func (u *UserDao) CheckByName(name string) (pojo.User, bool) {
	res := db.Model(&user).First(&user).Where("name =?", name)
	if res.RowsAffected <= 0 {
		return user, false
	}
	return user, true
}

// 更新token数据
func (u *UserDao) UpdateToken(access_token string, id uint) bool {
	user.ID = id
	res := db.Model(&user).Update("access_token", access_token)
	if res.RowsAffected != 1 {
		return false
	}
	return true
}
