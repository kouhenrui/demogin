package wxpay

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay"
	"net/http"
	"time"

	"github.com/go-pay/gopay/wechat/v3"
	//"github.com/go-pay/gopay/wechat"
)

/**
 * @ClassName wxpay
 * @Description TODO
 * @Author khr
 * @Date 2023/5/23 10:52
 * @Version 1.0
 */
var (
	WxPayClient *wechat.ClientV3
	err         error
	// 设置微信支付参数
	serialNo = "YOUR_APP_ID"  // 微信支付的App ID
	mchID    = "YOUR_MCH_ID"  // 商户号
	apiKey   = "YOUR_API_KEY" // API密钥

	priKey    = "YOU_PRIVATE_KEY"
	notifyURL = "http://localhost:8080/callback" // 支付回调通知URL
)

func wxPayMiddleware(c *gin.Context) {
	c.Set("wxpayClient", WxPayClient)
	c.Next()
}
func Routers(e *gin.Engine) {
	wxGroup := e.Group("/wx")

	wxGroup.Use(wxPayMiddleware)
	wxGroup.GET("/wxpay", HandleWXPayment) // 支付请求
	wxGroup.GET("/wxnotify", WXPayNotify)  // 生成支付链接并跳转至支付页面
}
func init() {
	// 创建微信支付客户端
	WxPayClient, err = wechat.NewClientV3(mchID, serialNo, apiKey, priKey)
	if err != nil {
		errors.New("微信支付链接失败")
	}
	// 打开Debug开关，输出请求日志，默认关闭
	//WxPayClient.DebugSwitch = gopay.DebugOn
	//WxPayClient.AutoVerifySign()
	// 自定义配置http请求接收返回结果body大小，默认 10MB
	//WxPayClient.SetBodySize() // 没有特殊需求，可忽略此配置

	// 设置国家：不设置默认 中国国内
	//    wechat.China：中国国内
	//    wechat.China2：中国国内备用
	//    wechat.SoutheastAsia：东南亚
	//    wechat.Other：其他国家
	//WxPayClient.SetCountry(wechat.China)

	//// 添加微信pem证书
	//WxPayClient.AddCertPemFilePath()
	//WxPayClient.AddCertPemFileContent()
	////或
	//// 添加微信pkcs12证书
	//WxPayClient.AddCertPkcs12FilePath()
	//WxPayClient.AddCertPkcs12FileContent()
}

func HandleWXPayment(c *gin.Context) {
	expire := time.Now().Add(10 * time.Minute).Format(time.RFC3339)
	// 初始化 BodyMap
	bm := make(gopay.BodyMap)
	bm.Set("sp_appid", "sp_appid").
		Set("sp_mchid", "sp_mchid").
		Set("sub_mchid", "sub_mchid").
		Set("description", "测试Jsapi支付商品").
		Set("out_trade_no", "ORDER12345").
		Set("time_expire", expire).
		Set("notify_url", "https://www.fmm.ink").
		SetBodyMap("amount", func(bm gopay.BodyMap) {
			bm.Set("total", 1).
				Set("currency", "CNY")
		}).
		SetBodyMap("payer", func(bm gopay.BodyMap) {
			bm.Set("sp_openid", "asdas")
		})

	fmt.Println("要发起支付")
	// 发起支付请求
	res, err := WxPayClient.V3TransactionH5(c, bm)
	if err != nil {
		fmt.Println("发起支付请求失败：", err)
		c.String(http.StatusInternalServerError, "failed")
		return
	}
	fmt.Println(res)
	// 处理支付结果
	if res.Code == http.StatusOK {
		// 获取预支付信息
		url := res.Response

		// 返回给客户端的数据
		data := gin.H{
			"prepay_id": url,
		}

		c.JSON(http.StatusOK, data)
	} else {
		fmt.Println("发起支付请求失败：", res.Error)
		c.String(http.StatusInternalServerError, "failed")
	}

}

func WXPayNotify(c *gin.Context) {
	// 解析支付回调参数
	//notifyData, err := WxPayClient.V3DecryptText(c.Request)
	//if err != nil {
	//	fmt.Println("解析支付回调通知失败：", err)
	//	c.String(http.StatusBadRequest, "failed")
	//	return
	//}
	//fmt.Printf("返回结果：%s", notifyData)
	//// 处理支付结果
	//if notifyData.GetString("return_code") == "SUCCESS" && notifyData.GetString("result_code") == "SUCCESS" {
	//	// 支付成功逻辑
	//	fmt.Println("支付成功")
	//	c.String(http.StatusOK, "success")
	//} else {
	//	// 支付失败逻辑
	//	fmt.Println("支付失败")
	//	c.String(http.StatusOK, "failed")
	//}
}
