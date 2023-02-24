package global

import "time"

//
//type MysqlConnect struct {
//	UserName string
//	PassWord string
//	HOST     string
//	DATABASE string
//	CHARSET  string
//	TimeOut  int64
//}
//type RedisConnect struct {
//	address  string
//	username string
//	password string
//	db       int
//	poolSize int
//}
//
//var RedisConfig = &RedisConnect{
//	"140.210.193.227:6379",
//	"root",
//	"123456",
//	1,
//	10,
//}
//
//var MysqlConfig = &MysqlConnect{
//	"root",
//	"123456",
//	"140.210.193.227:3306",
//	"test",
//	"utf8mb4",
//	10}
var ReuqestPaths = []string{"user/login", "user/register", "admin/login", "admin/register", "upload/file", "ws/connect", "upload/video"}

//图片格式
var PictureType = []string{"jpg", "png", "gif", "bmp", "tif",
	//不常用图片类型
	"pcx", "tga", "exif", "fpx", "svg", "psd", "cdr", "pcd", "dxf", "ufo", "eps", "ai", "raw", "WMF", "webp", "avif", "apng"}

//视频格式
var VideoType = []string{"avi", "wmv", "mpeg", "mp4", "m4v", "mov", "asf", "flv", "f4v", "rmvb", "rm", "3gp", "vob"}

const (
	UserLoginTime  = 5
	AdminLoginTime = 7
)

const (
	PORT               = ":8888"
	JWTKEY             = "jefnuhUKEWFKU@#$%^2546546"
	LANGUAGE           = "zh"
	JWTEXPTIME         = 7 * DayTime
	REDISJWTEXP        = 2
	SOCKETPORT         = 8889
	BYCTSECRET         = "iuag@%#(!)&#/$^&%@UHNVORE54"
	DayTime            = 24 * time.Hour
	FileMax     int64  = 2 << 20  //2Mb
	VideoMax    int64  = 50 << 20 //50Mb
	VideoPath   string = "dynamic"
	FilePath    string = "static"
)
