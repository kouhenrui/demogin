package adminService

import (
	"HelloGin/src/dto/reqDto"
	"HelloGin/src/dto/resDto"
	"HelloGin/src/global"
	"HelloGin/src/pojo"
	"HelloGin/src/service/userService"
	"HelloGin/src/util"
	"fmt"
)

//type AdminService struct {
//}

var admin pojo.Admin
var judge bool

// 引入dao层
var (
	permissionServiceImpl = pojo.RbacPermission()
	roleServiceImpl       = pojo.RbacRule()
	adminServiceImpl      = pojo.AdminServiceImpl()
)

/*
*
反射控制层登录参数，查询数据库账号是否相同，
比对密码一致性，将用户信息存入jwt令牌中，签发令牌和过期时间
*/
func AdminLogin(list reqDto.AdminLogin) (a bool, tokenAndExp interface{}) {
	//var ad=pojo.Admin{}
	switch list.Method {
	case "name":
		admin, judge = adminServiceImpl.CheckByName(list.Name)
	case "account":
		admin, judge = adminServiceImpl.CheckByAccount(list.Account)
	default:
		return false, util.METHOD_NOT_FILLED_ERROR
	}
	if !judge {
		return false, util.ACCOUT_NOT_EXIST_ERROR
	}
	pwd, deerr := util.DePwdCode(admin.Password, admin.Salt)
	if deerr != nil {
		return false, util.PASSWORD_RESOLUTION_ERROR
	}
	if pwd == "" {
		return false, util.PASSWORD_RESOLUTION_ERROR
	}
	if pwd != list.Password {
		return false, util.AUTH_LOGIN_PASSWORD_ERROR
	}
	_, role_name := roleServiceImpl.FindRoleName(uint(admin.Role))
	existOldToken := util.ExistRedis(admin.AccessToken)
	tokenKey := util.Rand6String6()
	var token string
	var exptime string
	stringTokenData := util.UserClaims{
		Id:       admin.ID,
		Name:     admin.Name,
		Account:  admin.Account,
		Role:     admin.Role,
		RoleName: role_name.Name,
	}
	switch list.Revoke {
	case true:
		if existOldToken {
			util.DelRedis(admin.AccessToken) //清除token
		}
		token, exptime = util.SignToken(stringTokenData, global.AdminLoginTime*global.DayTime)
		ok := adminServiceImpl.UpdateToken(tokenKey, admin.ID)
		if !ok {
			return false, util.AUTH_LOGIN_ERROR
		}
		redisDate := reqDto.LoginRedisDate{
			Token: token,
			Time:  exptime,
		}
		util.SetRedis(tokenKey, util.Marshal(redisDate), global.AdminLoginTime)
		tokenAndExp = resDto.TokenAndExp{
			token,
			exptime,
		}
	case false:
		if existOldToken {
			tokenValue := util.GetRedis(admin.AccessToken)
			mp := make(map[string]interface{})
			_, cs := util.UnMarshal([]byte(tokenValue), &mp)
			return true, cs
		}
		//token过期时
		token, exptime = util.SignToken(stringTokenData, global.AdminLoginTime*global.DayTime)
		ok := adminServiceImpl.UpdateToken(tokenKey, admin.ID)
		if !ok {
			return false, util.AUTH_LOGIN_ERROR
		}
		redisDate := reqDto.LoginRedisDate{
			Token: token,
			Time:  exptime,
		}
		util.SetRedis(tokenKey, util.Marshal(redisDate), global.AdminLoginTime)
		tokenAndExp = resDto.TokenAndExp{
			token,
			exptime,
		}
		break
	}
	return true, tokenAndExp
}

func AdminInfo(id int, name string) (bool, resDto.AdminInfo) {
	var adminInfo = resDto.AdminInfo{}
	var ok bool
	adminInfo, ok = adminServiceImpl.AdminInfo(id, name)
	if ok {
		return true, adminInfo
	}
	return false, adminInfo
}

// 分页模糊查询管理员
func AdminList(list reqDto.AdminList) interface{} {
	res := adminServiceImpl.AdminList(list)
	fmt.Println("adminlist:", res)
	return res
}

// 增加admin
func AdminAdd(add reqDto.AddAdmin) (bool, interface{}) {
	add.Salt = util.RandAllString()
	var pwd = add.Password
	//校验是否有密码，没有则为123456
	if add.Password == "" {
		pwd = string(123456)
	}
	//调用加密方法
	enPwd, _ := util.EnPwdCode(pwd, add.Salt)
	//加密密码
	add.Password = enPwd
	//检查名称是否重复
	if add.Name != "" {
		_, judge = adminServiceImpl.CheckByName(add.Name)
		if judge {
			return false, util.NAME_EXIST_ERROR
		}
	}
	if add.Name == "" {
		add.Name = "暂未命名"
	}
	_, judge = adminServiceImpl.CheckByAccount(add.Account)
	if judge {
		return false, util.ACCOUNT_EXIST_ERROR
	}
	ad := pojo.Admin{
		Salt:     add.Salt,
		Password: add.Password,
		Name:     add.Name,
		Account:  add.Account,
		Role:     add.Role}
	judge = adminServiceImpl.AddAdmin(ad)
	if judge {
		return true, util.ADD_SUCCESS
	} else {
		return false, util.ADD_ERROR
	}
}

// 调用userservice服务层的服务
func UserList(list reqDto.UserList) interface{} {
	res := userService.UserList(list)
	return res
}

// 登出
func AdminLogout() {
	util.DelRedis(admin.AccessToken)
}

// 权限列表
func PermissionList(list reqDto.PermissionList) interface{} {
	res := permissionServiceImpl.FindPermissionList(list)
	return res
}

/*权限增加*/
func PermissionAdd(permission reqDto.PermissionAdd) bool {
	per := pojo.Permission{
		Host:            permission.Host,
		Path:            permission.Path,
		AuthorizedRoles: permission.AuthorizedRoles,
		ForbiddenRoles:  permission.ForbiddenRoles,
		Method:          permission.Method,
		AllowAnyone:     permission.AllowAnyone,
	}
	return permissionServiceImpl.AddPermission(per)
}

/*权限修改*/
func PermissionUpdate(permission reqDto.PermissionUpdate) bool {
	permissionServiceImpl.SavePermission(permission)
	return true
}

// 权限删除
func Permissiondel(id int) (bool, string) {
	err, _ := permissionServiceImpl.FindPermissionById(id)
	if err == nil {
		if o := permissionServiceImpl.DelPermission(id); o {
			return true, util.DELETE_SUCCESS
		}
		return false, util.PERMISSION_NOT_FOUND_ERROR
	}
	return false, util.PERMISSION_NOT_FOUND_ERROR
}

/*权限详情*/
func PermissionIndo(id int) (bool, interface{}) {
	err, info := permissionServiceImpl.FindPermissionById(id)
	if err != nil {
		return false, err
	}
	return true, info
}
