package main

import (
	"HelloGin/src/controller/admin"
	"HelloGin/src/controller/com"
	"HelloGin/src/controller/user"
	"HelloGin/src/controller/ws"
	"HelloGin/src/global"
	"HelloGin/src/routers"
	"fmt"
	"log"
)

func main() {
	//挂载路由
	routers.Include(admin.Routers, ws.Routers, user.Routers, com.Routers)

	//初始化路由器,加载中间件等
	r := routers.InitRoute()
	log.Printf("程序配置文件加载无误,开始运行")

	var err error
	if global.HttpVersion {
		//http服务
		if err = r.Run(global.Port); err != nil {
			fmt.Errorf("端口占用,err:%v\n", err)
		}
	} else {
		//https服务
		if err = r.RunTLS(global.Port, "https/certificate.crt", "https/private.key"); err != nil {
			fmt.Errorf("端口占用,err:%v\n", err)
		}
	}

}
