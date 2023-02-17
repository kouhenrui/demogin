package adminDao

import (
	"HelloGin/src/dto/reqDto"
	"HelloGin/src/dto/resDto"
	"HelloGin/src/global"
	"HelloGin/src/pojo"
)

/**
* @program: work_space
*
* @description:dao层,数据连接的地方
*
* @author: khr
*
* @create: 2023-02-17 14:39
**/
var (
	db    = global.Db    //引用全局的数据连接
	admin = pojo.Admin{} //单个user返回
	//adminList    = &[]pojo.Admin{}      //多个user返回
	resAdminList = []resDto.AdminList{} //要查询的字段
	reslist      = resDto.CommonList{}  //返回的列表，包含数据和数量
)

type AdminDao struct{}

func AdminServiceImpl() AdminDao {
	return AdminDao{}
}

//分页,模糊查询用户
func (a *AdminDao) AdminList(list reqDto.AdminList) resDto.CommonList {
	query := db.Model(&admin)
	if list.Name != "" {
		query.Where("name like ?", "%"+list.Name+"%")
	}
	query.Limit(list.Take).Offset(int(list.Skip)).Find(&resAdminList)
	reslist.Count = uint(query.RowsAffected)
	reslist.List = resAdminList
	return reslist
}
func (a *AdminDao) CheckByAccount(account string) (pojo.Admin, bool) {
	res := db.Model(&admin).First(&admin).Where("account =?", account)
	if res.RowsAffected <= 0 {
		return admin, false
	}
	return admin, true
}

func (a *AdminDao) CheckByName(name string) (pojo.Admin, bool) {
	res := db.Model(&admin).First(&admin).Where("name =?", name)
	if res.RowsAffected <= 0 {
		return admin, false
	}
	return admin, true
}

func (a *AdminDao) UpdateToken(access_token string, id uint) bool {
	admin.ID = id
	res := db.Model(&admin).Update("access_token", access_token)
	if res.RowsAffected != 1 {
		return false
	}
	return true
}
