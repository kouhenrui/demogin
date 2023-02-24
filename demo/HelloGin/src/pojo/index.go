package pojo

import (
	"HelloGin/src/global"
	"fmt"
	"gorm.io/gorm"
	"time"
)

//数据库生成表
var db = global.Db
var (
	user      = &User{}
	admins    = &Admin{}
	t         = &Test{}
	e         = &Examp{}
	rbac_rule = &Rule{}
	rbac_per  = &Permission{}
	//rbac_res  = &Resource{}
)

func init() {
	db.AutoMigrate(
		user,
		admins,
		rbac_rule,

		rbac_per)
	//db.AutoMigrate(e)
	fmt.Println("表创建")
}

type Base struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"ID,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
