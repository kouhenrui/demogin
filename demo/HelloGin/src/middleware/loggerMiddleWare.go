package middleWare

import (
	"HelloGin/src/logger"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"math"
	"time"
)

func LoggerMiddleWare() gin.HandlerFunc {

	return func(c *gin.Context) {
		//bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		//c.Writer = bodyLogWriter
		startTime := time.Now()
		c.Next() // 调用该请求的剩余处理程序
		stopTime := time.Since(startTime)
		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds()/1000000))))
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		//reqbody := c.Request.Body
		param := c.Params
		query := c.Request.URL.RawQuery
		//errorMessage := c.Errors
		dataSize := c.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		//buf := n
		////d :=ew(bytes.Buffer)
		//		//buf.ReadFrom(reqbody)

		body, _ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		fmt.Println("---body/--- \r\n" + string(body))

		fmt.Println(c.Request.Proto, "HTTP请求版本")
		//for i:=0;i<reqbody{
		//
		//}
		//log.Println(b, "请求参数")
		//log.Println(*reqbody, "请求参数")
		//log.Println(c.Request.Response.Body, "返回参数")
		method := c.Request.Method
		url := c.Request.RequestURI
		resBody := c.Writer
		//fmt.Println("body", buf.Bytes())
		Log := logger.Logger.WithFields(
			logrus.Fields{
				//">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>": "\n",
				"SpendTime": spendTime,
				"path":      url,
				"Method":    method,
				"status":    statusCode,
				"Ip":        clientIP,
				"body":      string(body),
				"param":     param,
				"query":     query,
				//"erroMessage": errorMessage,

				"resBody": resBody,

				//"\n<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<": ,
			})
		if len(c.Errors) > 0 { // 矿建内部错误
			Log.Error(c.Errors.ByType(gin.ErrorTypePrivate))
		}
		if statusCode >= 500 {
			Log.Error()
		} else if statusCode >= 400 {
			Log.Warn()
		} else {
			Log.Info()
		}
	}
}
