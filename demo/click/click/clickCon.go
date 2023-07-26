package click

import (
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"time"
)

var (
	ClickDB driver.Conn
	err     error
)

func init() {
	var ctx = context.Background()
	ClickDB, err = clickhouse.Open(&clickhouse.Options{
		Addr: []string{"192.168.245.22:9000"},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: "default",
			Password: "",
		},
		DialTimeout: 5 * time.Second,
		//Compression: &clickhouse.Compression{
		//	Method: clickhouse.CompressionBrotli,
		//	Level:  5,
		//},
		// 必须添加协议方式
		//Protocol: clickhouse.HTTP,
	})

	if err != nil {
		fmt.Sprintf("连接错误:%s", ClickDB)
	}

	if err = ClickDB.Ping(ctx); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("Exception [%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		}
		fmt.Println("ping错误啦", err)
	}
	fmt.Println("连接成功啦")
}
