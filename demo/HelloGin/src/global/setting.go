package global

import "time"

type MysqlConnect struct {
	UserName string
	PassWord string
	HOST     string
	DATABASE string
	CHARSET  string
	TimeOut  int64
}
type RedisConnect struct {
	address  string
	username string
	password string
	db       int
	poolSize int
}

var ReuqestPaths = []string{"user/login", "user/register", "admin/login", "admin/register", "upload/file", "ws/connect"}
var PictureType = []string{"jpg", "png", "gif", "bmp", "tif",
	//不常用图片类型
	"pcx", "tga", "exif", "fpx", "svg", "psd", "cdr", "pcd", "dxf", "ufo", "eps", "ai", "raw", "WMF", "webp", "avif", "apng"}

const (
	UserLoginTime  = 5
	AdminLoginTime = 7
)

const (
	PORT        = ":8888"
	WSADDRESS   = "8889"
	JWTKEY      = "jefnuhUKEWFKU@#$%^2546546"
	LANGUAGE    = "zh"
	JWTEXPTIME  = 7 * DayTime
	REDISJWTEXP = 2
	SOCKETPORT  = 8889
	BYCTSECRET  = "iuag@%#(!)&#/$^&%@UHNVORE54"
	DayTime     = 24 * time.Hour
	HourTime    = 1 * time.Hour
	//AuthPath    = []string{"user/login", "user/regist", "admin/login"}
)

var RedisConfig = &RedisConnect{
	"140.210.193.227:6379",
	"root",
	"123456",
	1,
	10,
}

var MysqlConfig = &MysqlConnect{
	"root",
	"123456",
	"140.210.193.227:3306",
	"golang",
	"utf8mb4",
	10}
