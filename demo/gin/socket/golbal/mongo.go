package golbal

import (
	"context"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/mongodb-adapter/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

/**
 * @ClassName mongo
 * @Description TODO
 * @Author khr
 * @Date 2023/6/13 18:23
 * @Version 1.0
 */
var (
	Enforcer *casbin.Enforcer
	ctx      = context.Background()
)

//// CasbinRule 结构体
//type CasbinRule struct {
//	Ptype string `bson:"ptype"`
//	V0    string `bson:"v0"`
//	V1    string `bson:"v1"`
//	V2    string `bson:"v2"`
//	V3    string `bson:"v3"`
//	V4    string `bson:"v4"`
//}

func init() {
	// 创建 MongoDB 客户端连接
	clientOptions := options.Client().ApplyURI("mongodb://192.168.245.22:27017").SetAuth(options.Credential{
		Username: "admin",
		Password: "123456",
	})
	// 创建 Casbin 的 MongoDB 适配器
	adapter, err := mongodbadapter.NewAdapterWithCollectionName(clientOptions, "cabin", "policy", 30*time.Second)
	if err != nil {
		panic(err)
	}
	// 加载Casbin模型文件
	model, err := model.NewModelFromFile("auth_model.conf")
	if err != nil {
		fmt.Printf("Failed to load Casbin model: %v", err)
		panic(err)
	}
	// 创建Casbin Enforcer
	Enforcer, err = casbin.NewEnforcer(model, adapter)

	if err != nil {
		fmt.Printf("Failed to create Casbin enforcer: %v", err)
		panic(err)
	}
	fmt.Println("cabin加载成功")
	//_, err = Enforcer.AddPolicy("admin", "all", "write")
	//_, err = Enforcer.AddGroupingPolicy("root", "admin")
	//if err != nil {
	//	fmt.Printf("写入错误", err)
	//}
	//fmt.Println("写入数据状态", ok, err)

	name := Enforcer.GetPolicy()
	for _, p := range name {
		fmt.Println(p)
	}
	g := Enforcer.GetGroupingPolicy()
	for _, y := range g {
		fmt.Println(y)
	}
	fmt.Println("写入数据成")
}
