package userDao

import (
	"HelloGin/src/dto/reqDto"
	"HelloGin/src/dto/resDto"
	"HelloGin/src/global"
	"HelloGin/src/pojo"
)

var (
	db           = global.Db           //引用全局的数据连接
	user         = &pojo.User{}        //单个user返回
	userList     = &[]pojo.User{}      //多个user返回
	resUsersList = []resDto.UserList{} //要查询的字段
	reslist      = resDto.CommonList{} //返回的列表，包含数据和数量
)

type UserDao struct{}

func UserServiceImpl() UserDao {
	return UserDao{}
}

//分页,模糊查询用户
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
