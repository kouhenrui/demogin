package middleWare

import (
	"HelloGin/src/global"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"runtime/debug"
)

//返回的结果：
type Result struct {
	Code int         `json:"code"` //提示代码
	Msg  string      `json:"msg"`  //提示信息
	Data interface{} `json:"data"` //数据
}

func Recover(c *gin.Context) {
	defer func() {
		r := recover()
		result := global.NewResult(c)
		if r != nil {
			log.Println("打印错误信息")
			//打印错误堆栈信息
			log.Printf("panic: %v\n", r)

			debug.PrintStack()
			result.Error(http.StatusInternalServerError, errorToString(r))
			return
			//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
			c.Abort()
		}
	}()
	//加载完 defer recover，继续后续接口调用
	c.Next()
}

// recover错误，转string
func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}
