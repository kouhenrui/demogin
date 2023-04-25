package global

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	"log"
	"strings"

	_ "github.com/casbin/casbin/v2"
)

/**
 * @ClassName casbin
 * @Description TODO
 * @Author khr
 * @Date 2023/4/24 14:25
 * @Version 1.0
 */
func check(e *casbin.Enforcer, sub, dom, obj, act string) {
	ok, _ := e.Enforce(sub, dom, obj, act)

	//fmt.Println(er, "err")
	if ok {
		fmt.Printf("%s CAN %s %s in %s\n", sub, act, obj, dom)
	} else {
		fmt.Printf("%s CANNOT %s %s in %s\n", sub, act, obj, dom)
	}
}

/*
 * @MethodName KeyMatchFunc
 * @Description 正则匹配
 * @Author khr
 * @Date 2023/4/24 14:26
 */
func KeyMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[2].(string)
	return KeyMatch(name1, name2), nil
}
func KeyMatch(key1, key2 string) bool {
	i := strings.Index(key2, "*")
	if i == -1 {
		return key1 == key2
	}

	if len(key1) > i {
		return key1[:i] == key2[:i]
	}

	return key1 == key2[:i]
}

func CasbinInit() {
	//var err error
	//fmt.Printf("打印权限", casbinCon)
	db := CasbinConfig.UserName + ":" + CasbinConfig.PassWord + "@tcp(" + CasbinConfig.HOST + ":" + CasbinConfig.Port + ")/"
	//fmt.Println("连接的信息：", CasbinConfig.DATABASE, db)
	a, aerr := gormadapter.NewAdapter(CasbinConfig.Type, db)
	e, eerr := casbin.NewEnforcer("auth_model.conf", a)
	//fmt.Println("e", e)
	if aerr != nil {
		fmt.Printf("权限表为创建，错误原因：%s", aerr)
	}
	if eerr != nil {
		fmt.Println("加载模型出现错误", eerr)
	}
	log.Printf("权限初始化成功")

	//创建表

	//e.AddFunction("my_func", KeyMatchFunc)
	check(e, "dajun", "root", "data1", "all")
	//check(e, "lili", "dev", "data2", "read")
	//check(e, "dajun", "tenant1", "data1", "read")
	//check(e, "dajun", "tenant2", "data2", "read")
	//check(e, "root", "", "data2", "read")
}
