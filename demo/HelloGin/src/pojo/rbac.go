package pojo

import "gorm.io/gorm"

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
