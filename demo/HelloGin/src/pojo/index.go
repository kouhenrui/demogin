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
	user   = &User{}
	admins = &Admin{}
	t      = &Test{}
	e      = &Examp{}
)

func init() {
	//db.AutoMigrate(user, admins, t)
	//db.AutoMigrate(e)
	fmt.Println("数据库创建")
}

type Base struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"ID,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
