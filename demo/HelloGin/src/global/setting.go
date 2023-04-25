package global

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"log"
	"time"
)

type mysqlConnect struct {
	UserName string `json:"username,omitempty" yaml:"username" mapstructure:"username"`
	PassWord string `json:"password,omitempty" yaml:"password" mapstructure:"password"`
	HOST     string `json:"host,omitempty" yaml:"host" mapstructure:"host"`
	Port     string `json:"port,omitempty" yaml:"port" mapstructure:"port"`
	DATABASE string `json:"database,omitempty" yaml:"database" mapstructure:"database"`
	CHARSET  string `json:"charset,omitempty" yaml:"charset" mapstructure:"charset"`
	TimeOut  int64  `json:"timeout,omitempty" yaml:"timeout" mapstructure:"timeout"`
}
type redisConnect struct {
	UserName   string `json:"username,omitempty" yaml:"username"`
	PassWord   string `json:"password,omitempty" yaml:"password"`
	Host       string `json:"host,omitempty" yaml:"host"`
	Port       string `json:"port,omitempty" yaml:"port"`
	Db         int    `json:"db,omitempty" yaml:"db"`
	PoolSize   int    `json:"poolsize,omitempty" yaml:"poolsize"`
	MaxRetries int    `json:"maxRetries" yaml:"maxRetries"`
}
type rabbitmqConnect struct {
	Url      string `json:"url,omitempty" yaml:"url" `
	UserName string `json:"username" yaml:"username"`
	PassWord string `json:"password" yaml:"password"`
	Host     string `json:"host" yaml:"host"`
	Port     string `json:"port" yaml:"port"`
}
type casbinConnect struct {
	Type     string `json:"type" yaml:"type" mapstructure:"type"`
	UserName string `json:"username,omitempty" yaml:"username" mapstructure:"username"`
	PassWord string `json:"password,omitempty" yaml:"password" mapstructure:"password"`
	HOST     string `json:"host,omitempty" yaml:"host" mapstructure:"host"`
	Port     string `json:"port,omitempty" yaml:"port" mapstructure:"port"`
	DATABASE string `json:"database,omitempty" yaml:"database" mapstructure:"database"`
	Exist    bool   `json:"exist,omitempty" yaml:"exist" mapstructure:"exist"`
}

//	type etcdConnect struct {
//		Host string `json:"host,omitempty" yaml:"host"`
//		Port string `json:"port,omitempty" yaml:"port"`
//	}
type LogCof struct {
	LogPath  string `json:"logPath,omitempty" yaml:"logPath"`
	LinkName string `json:"linkName,omitempty" yaml:"linkName"`
}

var (
	Port           string
	MysqlConfig    mysqlConnect
	RedisConfig    redisConnect
	RabbitMQConfig rabbitmqConnect
	LogConf        LogCof
	CasbinConfig   casbinConnect
	//EtcdConfig     etcdConnect
	ReuqestPaths = []string{"user/login", "user/register", "admin/login", "admin/register", "upload/file", "ws/connect", "upload/video", "mq/post", "ync/test"}
	PictureType  = []string{"jpg", "png", "gif", "bmp", "tif", "pcx", "tga", "exif", "fpx", "svg", "psd", "cdr", "pcd", "dxf", "ufo", "eps", "ai", "raw", "WMF", "webp", "avif", "apng"} //图片类型
	VideoType    = []string{"avi", "wmv", "mpeg", "mp4", "m4v", "mov", "asf", "flv", "f4v", "rmvb", "rm", "3gp", "vob"}                                                                  // 视频格式
	RoleName     = []string{}
	EtcdArry     = []string{}
	KafkaArry    = []string{}
	HttpVersion  bool
)

const (
	UserLoginTime         = 5
	AdminLoginTime        = 7
	JWTKEY                = "jefnuhUKEWFKU@#$%^2546546"
	LANGUAGE              = "zh"
	DayTime               = 24 * time.Hour
	FileMax        int64  = 2 << 20  //2Mb
	VideoMax       int64  = 50 << 20 //50Mb
	VideoPath      string = "dynamic"
	FilePath       string = "static"
	//JWTEXPTIME            = 7 * DayTime
	//REDISJWTEXP           = 2
	//SOCKETPORT            = 8889
	//BYCTSECRET            = "iuag@%#(!)&#/$^&%@UHNVORE54"
)

var v *viper.Viper

func init() {
	var err error
	// 构建 Viper 实例
	v = viper.New()
	v.SetConfigFile("conf.yaml") // 指定配置文件路径
	v.SetConfigName("conf")      // 配置文件名称(无扩展名)
	v.SetConfigType("yaml")      // 如果配置文件的名称中没有扩展名，则需要配置此项
	//viper.AddConfigPath("/etc/appname/")   // 查找配置文件所在的路径
	//viper.AddConfigPath("$HOME/.appname")  // 多次调用以添加多个搜索路径
	v.AddConfigPath(".") // 还可以在工作目录中查找配置
	// 查找并读取配置文件
	if err = v.ReadInConfig(); err != nil { // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.OnConfigChange(func(e fsnotify.Event) {
		log.Print("Config file updated.")
		viperLoadConf() // 加载配置的方法
	})
	v.WatchConfig() //开启监听
	viperLoadConf()

}
func viperLoadConf() {
	//读取单条配置文件
	Port = ":" + v.GetString("port")

	//设置http1.0还是2.0
	HttpVersion = v.GetBool("protocol")

	//日志路径及名称设置
	logConfig := v.GetStringMap("log")

	//读取mysql,redis,rabbitmq,casbin
	mysql := v.GetStringMap("mysql")
	redis := v.GetStringMap("redis")
	mq := v.GetStringMap("rabbitmq")
	cn := v.GetStringMap("casbin")

	//map转struct
	mapstructure.Decode(mysql, &MysqlConfig)
	mapstructure.Decode(redis, &RedisConfig)
	mapstructure.Decode(mq, &RabbitMQConfig)
	mapstructure.Decode(logConfig, &LogConf)
	mapstructure.Decode(cn, &CasbinConfig)
	etcd := v.GetString("etcd")
	kafka := v.GetString("kafka")
	EtcdArry = append(EtcdArry, etcd)
	KafkaArry = append(KafkaArry, kafka)
	//设置超管
	role := v.GetString("superadmin")
	RoleName = append(RoleName, role)
	log.Printf("全局文件读取无误,开始载入")
	Mqinit()
	Dbinit()
	Redisinit()
	Loginit()
	CasbinInit()
	//EtcdInit()
}
