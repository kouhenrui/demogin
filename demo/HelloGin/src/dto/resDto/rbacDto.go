package resDto

type RoleList struct {
	Id   uint   `json:"id"`
	Name string `json:"name" `
}

type GroupList struct {
	Id           uint   `json:"id"`
	Name         string `json:"name" `
	RoleId       string `json:"role_id"`
	PermissionId string `json:"permission_id"`
}

type PermissonList struct {
	Id              uint   `json:"id"`
	Host            string `json:"host"`
	Path            string `json:"path"`
	Method          string `json:"method"`
	AuthorizedRoles string `json:"authorized_roles"`
	ForbiddenRoles  string `json:"forbidden_roles"`
	AllowAnyone     bool   `json:"allow_anyone"`
}

type PermissionInfo struct {
	Id              uint   `json:"id"`
	Host            string `json:"host"`
	Path            string `json:"path"`
	Method          string `json:"method"`
	AuthorizedRoles string `json:"authorized_roles"`
	ForbiddenRoles  string `json:"forbidden_roles"`
	AllowAnyone     bool   `json:"allow_anyone"`
}
