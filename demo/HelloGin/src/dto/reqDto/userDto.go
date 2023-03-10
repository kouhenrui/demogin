package reqDto

type UserLogin struct {
	Account  string `json:"account" `
	Name     string `json:"name"  `
	Password string `json:"password" binding:"required" `
	Method   string `json:"method" binding:"required" gorm:"default:false;one of account,name"`
	Revoke   bool   `json:"revoke" validate:"required"`
}
type UpdateUser struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Account  string `json:"account"`
}
type AddUser struct {
	Name     string `json:"name,omitempty" binding:"required" validate:"required"`
	Password string `json:"password,omitempty"  binding:"required" validate:"required"`
	Account  string `json:"account,omitempty"  binding:"required" validate:"omitempty"`
	Salt     string `json:"salt,omitempty"`
	Role     int    `json:"role,omitempty"`
}
type UserList struct {
	Take int    `json:"take,omitempty" binding:"required"`
	Skip uint   `json:"skip,omitempty"`
	Name string `json:"name,omitempty"`
}
