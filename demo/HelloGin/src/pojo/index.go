package pojo

import (
	"HelloGin/src/dto/resDto"
	"HelloGin/src/global"
	"fmt"
	"gorm.io/gorm"
	"time"
)

// 数据库生成表
var db = global.Db
var reslist = resDto.CommonList{}
var count int64
var (
	userpojo  = &User{}
	adminpojo = &Admin{}
	rbac_rule = &Rule{}
	rbac_per  = &Permission{}
	group     = &Group{}
)

func init() {
	db.AutoMigrate(
		user,
		adminpojo,
		rbac_rule,
		rbac_per,
		group,
	)
	fmt.Println("表创建")
}

type Base struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"ID,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
