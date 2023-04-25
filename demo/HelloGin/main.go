package main

import (
	"HelloGin/src/controller/admin"
	"HelloGin/src/controller/async"
	"HelloGin/src/controller/upload"
	"HelloGin/src/controller/user"
	"HelloGin/src/controller/ws"
	"HelloGin/src/global"
	"HelloGin/src/routers"
	"fmt"
	"log"
)

func main() {
	routers.Include(upload.Routers, async.Routers, admin.Routers, ws.Routers, user.Routers)
	r := routers.InitRoute()
	////读取ini配置文件
	//Cfg, inierr := ini.Load("conf.ini")
	//if inierr != nil {
	//	fmt.Printf("Fail to read file: %v", inierr)
	//	os.Exit(1)
	//}
	//fmt.Println(cfg.EdctConf.Address)
	//port := Cfg.Section("server").Key("http_port").String()
	//log.Printf("程序开始运行", port)
	//serve:=r.
	log.Printf("程序配置文件加载无误,开始运行")
	var err error
	if global.HttpVersion {
		//http服务
		if err = r.Run(global.Port); err != nil {
			fmt.Errorf("端口占用,err:%v\n", err)
		}
	} else {
		//https服务
		if err = r.RunTLS(global.Port, "https/server.crt", "https/server.key"); err != nil {
			fmt.Errorf("端口占用,err:%v\n", err)
		}
	}

}
