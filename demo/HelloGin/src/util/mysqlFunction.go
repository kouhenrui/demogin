package util

import (
	"HelloGin/src/global"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"

	"time"
)

var db = global.Db

type TimeModel struct {
	uuid      uuid.UUID `json:"id" gorm:"type:char(36);primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

//GetDB
func GetDB() *gorm.DB {
	return db
}

//判断是否存在
func Exist(condition []string, table string) bool {
	//res := db.Model().Where(condition).Take(&table)
	//if len(res.RowsAffected) {
	//
	//}

	return true
}

//分页
func GetTableAll(take int, skip int, query []string) *gorm.DB {
	return db.Offset(skip).Limit(take)
}
