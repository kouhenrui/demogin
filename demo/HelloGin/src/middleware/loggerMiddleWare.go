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

type BodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w BodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
func (w BodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}
func LoggerMiddleWare() gin.HandlerFunc {

	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next() // 调用该请求的剩余处理程序
		stopTime := time.Since(startTime)
		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds()/1000000))))
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		param := c.Params
		query := c.Request.URL.RawQuery
		//errorMessage := c.Errors
		//dataSize := c.Writer.Size()
		//if dataSize < 0 {
		//	dataSize = 0
		//}
		requestBody, _ := c.GetRawData()
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))
		bodyLogWriter := &BodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		responseBody := bodyLogWriter.body.String()
		//bodyLogWriter := &BodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		//c.Writer = bodyLogWriter
		//fmt.Println("---body/--- \r\n" + string(body))
		//
		//fmt.Println(c.Request.Proto, "HTTP请求版本")

		//fmt.Println("请求提", string(requestBody))
		method := c.Request.Method
		url := c.Request.RequestURI
		//resBody := c.Writer
		//fmt.Println("body", buf.Bytes())
		Log := logger.Logger.WithFields(
			logrus.Fields{
				"SpendTime":   spendTime,
				"path":        url,
				"Method":      method,
				"status":      statusCode,
				"proto":       c.Request.Proto,
				"Ip":          clientIP,
				"body":        string(requestBody),
				"param":       param,
				"query":       query,
				"resBody":     responseBody,
				"erroMessage": c.Errors.Last(),
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
