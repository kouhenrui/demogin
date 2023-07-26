package golbal

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 统一返回格式
func FormatResponse() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// 判断是否有错误信息
		if len(c.Errors) > 0 {
			fmt.Println("出现错误", c.Errors.Last().Error())
			// 返回错误信息
			c.JSON(http.StatusOK, gin.H{

				"code": http.StatusInternalServerError,
				"msg":  c.Errors.Last().Error(),
				"data": "",
			})
			return
		}

		// 判断是否有返回数据
		if c.Writer.Status() == http.StatusOK {
			// 获取返回数据
			data, ok := c.Get("res")
			if ok {
				// 格式化返回数据
				c.JSON(http.StatusOK, gin.H{
					"code": http.StatusOK,
					"msg":  "",
					"data": data,
				})
				return
			}
		}

		// 返回空数据
		c.JSON(http.StatusOK, gin.H{})
	}
}

// 全局日志中间件
//func LoggerMiddleWare() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		startTime := time.Now()
//		requestBody, _ := c.GetRawData()
//		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))
//		rbody := string(requestBody)
//		query := c.Request.URL.RawQuery
//		c.Next() // 调用该请求的剩余处理程序
//		stopTime := time.Since(startTime)
//		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds()/1000000))))
//		statusCode := c.Writer.Status()
//		clientIP := c.ClientIP()
//		method := c.Request.Method
//		url := c.Request.RequestURI
//		Log := Logger.WithFields(
//			logrus.Fields{
//				"SpendTime": spendTime,       //接口花费时间
//				"path":      url,             //请求路径
//				"Method":    method,          //请求方法
//				"status":    statusCode,      //接口返回状态
//				"proto":     c.Request.Proto, //http请求版本
//				"Ip":        clientIP,        //IP地址
//				"body":      rbody,           //请求体
//				"query":     query,           //请求query
//				"message":   c.Errors,        //返回错误信息
//			})
//		if len(c.Errors) > 0 { // 矿建内部错误
//			Log.Error(c.Errors.ByType(gin.ErrorTypePrivate))
//		}
//		if statusCode > 200 {
//			Log.Error()
//		} else {
//			Log.Info()
//		}
//	}
//}

//var visitorMap = make(map[string]*rate.Limiter) // 存储IP地址和速率限制器的映射
//var mu sync.Mutex                               // 互斥锁，保证并发安全
//// IP请求次数拦截中间件
//// 定义一个中间件函数，用于拦截请求并进行计数
//func IPInterceptor() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		ip := c.ClientIP()
//		if ip == "" {
//			ip = c.Request.RemoteAddr
//		}
//		if util.ExistIn(ip, IpAccess) {
//			c.Next()
//		}
//		path := c.Request.URL.Path
//		//fmt.Println(ip, path)
//
//		// 组合出 key
//		key := fmt.Sprintf("request:%s:%s", ip, path)
//		//fmt.Print("key", key)
//		// 将请求次数 +1，并设置过期时间
//		_, err := Redis.Incr(c, key).Result()
//
//		if err != nil {
//			// 记录日志
//			fmt.Println("incr error:", err)
//			c.AbortWithStatus(http.StatusInternalServerError)
//			return
//		}
//		Redis.Expire(c, key, time.Hour)
//
//		// 获取当前IP在 path 上的请求次数
//		accessTime, err := Redis.Get(c, key).Int()
//
//		if err != nil {
//			// 记录日志
//			fmt.Println("get error:", err)
//			c.AbortWithStatus(http.StatusInternalServerError)
//			return
//		}
//		//ip一小时内访问路径超过次数限制，拒绝访问
//		if accessTime > 60 {
//			requestLimit := fmt.Sprintf("request:%s:%s", ip, path)
//			Redis.RPush(c, RedisReqLimitUrl, requestLimit)
//			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "too many requests"})
//			return
//		}
//		mu.Lock()
//		_, ok := visitorMap[ip]
//		var limiter = rate.NewLimiter(1, 10) // 设置限制为1个请求/秒，最多允许10个并发请求
//		// 如果该IP地址不存在，则创建一个速率限制器
//		if !ok {
//			visitorMap[ip] = limiter
//		}
//		mu.Unlock()
//		// 尝试获取令牌，如果没有可用的令牌则阻塞
//		if !limiter.Allow() {
//			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "too many requests"})
//			return
//		}
//		c.Next()
//	}
//}
