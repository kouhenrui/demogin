package pojo

import (
	"HelloGin/src/util"
	"gorm.io/gorm"
	"log"
	"reflect"
)

type Test struct {
	gorm.Model `json:"gorm.Model"`
	Name       string `json:"name,omitempty"`
	Password   string `json:"password,omitempty"`
}
type TestInt interface {
	Write(interface{}) (bool, interface{})
}

func (i Test) Write(p interface{}) (bool, interface{}) {
	log.Println(p, "传输数据")
	user := reflect.ValueOf(p)
	log.Println(user, "反射接口类型数据")
	i.Name = user.FieldByName("name").String()
	i.Password = user.FieldByName("password").String()
	log.Println(i.Name, "i.name is ")

	log.Println(i.Password, "i.Password is ")
	res := db.Create(i)
	if res.RowsAffected == 0 {
		return false, util.INSET_USER_ERROR
	}
	return true, util.SUCCESS
}
