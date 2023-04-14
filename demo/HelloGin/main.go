package main

import (
	"HelloGin/src/controller/admin"
	"HelloGin/src/controller/async"
	"HelloGin/src/controller/upload"
	"HelloGin/src/controller/user"
	"HelloGin/src/controller/ws"
	"HelloGin/src/routers"
	"fmt"
	"gopkg.in/ini.v1"
	"log"
	"os"
)

func main() {
	routers.Include(upload.Routers, async.Routers, admin.Routers, ws.Routers, user.Routers)
	r := routers.InitRoute()
	//读取配置文件
	Cfg, inierr := ini.Load("conf.ini")
	if inierr != nil {
		fmt.Printf("Fail to read file: %v", inierr)
		os.Exit(1)
	}
	//fmt.Println(cfg.EdctConf.Address)
	port := Cfg.Section("server").Key("http_port").String()
	log.Printf("程序开始运行")
	if err := r.Run(port); err != nil {
		fmt.Errorf("端口占用,err:%v\n", err)
	}
}
