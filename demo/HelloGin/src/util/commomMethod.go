package util

import (
	"fmt"
	"reflect"
)

/*
 * @MethodName ExistIn
 * @Description 判断参数是否存在
 * @Author khr
 * @Date 2023/4/14 8:52
 */

func ExistIn(param string, paths []string) bool {
	for _, v := range paths {
		if param == v {
			return true
		}
	}
	return false
}

/*
 * @MethodName DtoToStruct
 * @Description dto转结构体
 * @Author khr
 * @Date 2023/4/14 8:52
 */

func DtoToStruct(dto interface{}, entity interface{}) {
	dtoValue := reflect.ValueOf(dto)
	entityValue := reflect.ValueOf(entity).Elem()
	for i := 0; i < dtoValue.NumField(); i++ {
		fieldName := dtoValue.Type().Field(i).Name
		fieldValue := dtoValue.Field(i)
		entityFieldValue := entityValue.FieldByName(fieldName)
		if entityFieldValue.IsValid() && entityFieldValue.Type() == fieldValue.Type() {
			entityFieldValue.Set(fieldValue)
		}
	}
	fmt.Println(entity)
}
