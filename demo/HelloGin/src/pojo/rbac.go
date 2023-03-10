package pojo

import (
	"HelloGin/src/dto/reqDto"
	"HelloGin/src/dto/resDto"
	"fmt"
	"gorm.io/gorm"
)

/**
* @program: work_space
*
* @description:rbac结构体
*
* @author: khr
*
* @create: 2023-02-21 09:27
**/
// Rule即规则，用于定义Resource和Permission之间的关系
type Rule struct {
	gorm.Model
	// ID决定了Rule的优先级。
	// ID值越大意味着Rule的优先级越高。
	// 当请求被多个规则同时匹配时，grbac将仅使用具有最高ID值的规则。
	// 如果有多个规则同时具有最大的ID，则将随机使用其中一个规则。
	Name string `json:"name"`
}

// Permission用于定义权限控制信息
type Permission struct {
	gorm.Model

	// Host 定义资源的Host，允许使用增强的通配符。
	Host string `json:"host" gorm:""`
	// Path 定义资源的Path，允许使用增强的通配符。
	Path string `json:"path"`
	// Method 定义资源的Method，允许使用增强的通配符。
	Method string `json:"method"`

	// AuthorizedRoles定义允许访问资源的角色
	// 支持的类型: 非空字符串，*
	//      *: 意味着任何角色，但访问者应该至少有一个角色，
	//      非空字符串：指定的角色
	AuthorizedRoles string `json:"authorized_roles"`
	// ForbiddenRoles 定义不允许访问指定资源的角色
	// ForbiddenRoles 优先级高于AuthorizedRoles
	// 支持的类型：非空字符串，*
	//      *: 意味着任何角色，但访问者应该至少有一个角色，
	//      非空字符串：指定的角色
	//
	ForbiddenRoles string `json:"forbidden_roles"`
	// AllowAnyone的优先级高于 ForbiddenRoles、AuthorizedRoles
	// 如果设置为true，任何人都可以通过验证。
	// 请注意，这将包括“没有角色的人”。
	AllowAnyone bool `json:"allow_anyone"gorm:"default:false;not bull"`
}

type Group struct {
	gorm.Model
	Name         string `json:"name"`
	RoleId       string `json:"role_id"`
	PermissionId string `json:"permission_id"`
}

var (
	groups          = []Group{}
	rules           = []Rule{}
	permissions     = []Permission{}
	rolesList       = []resDto.RoleList{}
	groupsList      = []resDto.GroupList{}
	permissionsList = []resDto.PermissonList{}
)

func RbacRule() Rule {
	return Rule{}
}

func RbacGroup() Group {
	return Group{}
}
func RbacPermission() Permission {
	return Permission{}
}

// 角色查询
func (r *Rule) FindRoleName(id uint) (bool, *Rule) {
	r.ID = id
	res := db.Find(&r)
	if res.RowsAffected != 1 {
		return false, r
	}
	return true, r
}

// 增加，修改角色
func (r *Rule) AddRole(rule Rule) bool {
	res := db.Save(&rule)
	if res.RowsAffected != 1 {
		return false
	}
	return true
}

// 角色列表
func (r *Rule) FindRoleList(list reqDto.RbacList) resDto.CommonList {
	query := db.Model(&r)
	if list.Name != "" {
		query.Where("name like ?", "%"+list.Name+"%")
	}
	query.Limit(list.Take).Offset(int(list.Skip)).Find(&rolesList)
	reslist.Count = uint(query.RowsAffected)
	reslist.List = rolesList
	return reslist
}

//// 修改角色
//func (r *Rule) SaveRole(rule Rule) bool {
//	res := db.Save(&rule)
//	if res.RowsAffected != 1 {
//		return false
//	}
//	return true
//}

// 删除角色
func (r *Rule) DelRole(id int) bool {
	r.ID = uint(id)
	res := db.Delete(&r)
	if res.RowsAffected != 1 {
		return false
	}
	return true
}

// 查看组
func (g *Group) FindGroupName(id uint) (bool, *Group) {
	g.ID = id
	res := db.Find(&g)
	if res.RowsAffected != 1 {
		return false, g
	}
	return true, g
}

// 增加，修改组
func (g *Group) AddGroup(group Group) bool {
	res := db.Save(&group)
	if res.RowsAffected != 1 {
		return false
	}
	return true
}

// 组列表
func (g *Group) FindGroupList(list reqDto.RbacList) resDto.CommonList {
	query := db.Model(&g)
	if list.Name != "" {
		query.Where("name like ?", "%"+list.Name+"%")
	}
	query.Limit(list.Take).Offset(int(list.Skip)).Find(&groupsList)
	reslist.Count = uint(query.RowsAffected)
	reslist.List = groupsList
	return reslist
}

// 修改组
func (g *Group) SaveGroup(grop Group) bool {
	res := db.Updates(grop)
	if res.RowsAffected != 1 {
		return false
	}
	return true
}

// 删除组
func (g *Group) DelGroup(id int) bool {
	g.ID = uint(id)
	res := db.Delete(&g)
	if res.RowsAffected != 1 {
		return false
	}
	return true
}

// 通过id查看权限
func (p *Permission) FindPermissionName(id uint) (bool, *Permission) {
	p.ID = id
	res := db.Find(&p)
	if res.RowsAffected != 1 {
		return false, p
	}
	return true, p
}

// 通过请求路径查看权限
func (p *Permission) FindPermissionByPath(path string) (bool, *Permission) {
	p.Path = path
	res := db.Find(&p).Where("path like ?", "%"+path+"%")
	if res.RowsAffected != 1 {
		return false, p
	}
	return true, p
}

// 增加,修改权限
func (p *Permission) AddPermission(permission Permission) bool {
	fmt.Println(permission)
	res := db.Save(&permission)
	if res.RowsAffected != 1 {
		return false
	}
	return true
}

// 权限列表
func (p *Permission) FindPermissionList(list reqDto.PermissionList) resDto.CommonList {
	query := db.Model(&p)
	if list.Path != "" {
		query.Where("path like ?", "%"+list.Path+"%")
	}
	query.Limit(list.Take).Offset(int(list.Skip)).Find(&permissionsList)
	reslist.Count = uint(query.RowsAffected)
	reslist.List = permissionsList
	return reslist
}

// 修改权限
func (p *Permission) SavePermission(permission Permission) bool {
	res := db.Save(&permission)
	if res.RowsAffected != 1 {
		return false
	}
	return true
}

// 删除权限
func (p *Permission) DelPermission(id int) bool {
	p.ID = uint(id)
	res := db.Delete(&p)
	if res.RowsAffected != 1 {
		return false
	}
	return true
}
