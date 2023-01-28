package global

import (
	_ "database/sql"
	"fmt"
	"gorm.io/driver/mysql"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// 定义db全局变量
var Db *gorm.DB

// 初始化链接
func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local", MysqlConfig.UserName, MysqlConfig.PassWord, MysqlConfig.HOST, MysqlConfig.DATABASE, MysqlConfig.CHARSET) //&timeout=%s , MysqlConfig.TimeOut

	var err error
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true, //跳过默认事务
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //false 复数形式
			//TablePrefix:   "",    //表名前缀 User的表名应该是t_users
		},
		DisableForeignKeyConstraintWhenMigrating: true, //设置成为逻辑外键(在物理数据库上没有外键，仅体现在代码上)

	})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	sqlDB, _ := Db.DB()
	//设置数据库连接池参数
	sqlDB.SetMaxOpenConns(100) //设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(20)  //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。

	Db.AutoMigrate()

}

//GetDB
func GetDB() *gorm.DB {
	return Db
}
