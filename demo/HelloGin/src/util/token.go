package util

import (
	"HelloGin/src/global"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var jwtkey = []byte(global.JWTKEY)
var str string

//var userInfo pojo.User
type UserClaims struct {
	Name    string `json:"name"`
	Role    int    `json:"role"`
	Account string `json:"account"`
	Id      uint   `json:"id"`
}

type AllClaims struct {
	jwt.StandardClaims
	User UserClaims
}

//颁发token
func SignToken(infoClaims UserClaims, day time.Duration) (string, string) {
	expireTime := time.Now().Add(day) //7天过期时间
	claims := &AllClaims{
		User: infoClaims,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),
			Issuer:    "khr",  // 签名颁发者
			Subject:   "sign", //签名主题
		},
	}
	fmt.Println(claims, "封装的信息")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println(token, "塞入的信息")
	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		fmt.Println(err, "生成token错误")
	}
	tFStr := expireTime.Format("2006-01-02 15:04:05")
	return tokenString, tFStr
}

//验证token
func AnalysyToken(c *gin.Context) interface{} {
	result := global.NewResult(c)
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		result.Success(gin.H{"code": http.StatusUnauthorized, "msg": NO_AUTHORIZATION})
		c.Abort()
		return result
	}

	user := ParseToken(tokenString)
	return user.User
}

// 解析Token
func ParseToken(tokenString string) *AllClaims {
	//claims := &Claims{}
	//解析token
	token, _ := jwt.ParseWithClaims(tokenString, &AllClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtkey, nil
	})
	user, _ := token.Claims.(*AllClaims)
	fmt.Println(user, "打印")
	return user
}
