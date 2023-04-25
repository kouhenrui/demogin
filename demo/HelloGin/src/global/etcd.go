package global

import (
	c "context"
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"golang.org/x/net/context"
	"log"
	"time"
)

/**
 * @ClassName etcd
 * @Description TODO
 * @Author khr
 * @Date 2023/4/25 13:56
 * @Version 1.0
 */

var EtcdConn *clientv3.Client

// var RabbitChannel *clientv3.
func EtcdInit() {
	var err error
	//url := "http://" + EtcdConfig.Host + ":" + EtcdConfig.Port
	EtcdConn, err = clientv3.New(clientv3.Config{
		Endpoints:   EtcdArry,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	log.Printf("etcd连接成功")
	defer EtcdConn.Close()
}

func EtcdPut() {
	var err error
	ctx, cancel := c.WithTimeout(c.Background(), time.Second)
	_, err = EtcdConn.Put(ctx, "lmh", "lmh")
	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return
	}
}
func EtcdGet() {
	ctx, cancel := c.WithTimeout(c.Background(), time.Second)
	resp, errWatch := EtcdConn.Get(ctx, "lmh")
	cancel()
	if errWatch != nil {
		fmt.Printf("get from etcd failed, err:%v\n", errWatch)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	}
}

// watch key:lmh change
func EtcdWatch() {

	rch := EtcdConn.Watch(context.Background(), "lmh") // <-chan WatchResponse
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("Type: %s Key:%s Value:%s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}
