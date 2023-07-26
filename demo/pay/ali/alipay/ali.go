package alipay

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/smartwalle/alipay/v3"
	"net/http"
	"net/url"
	"time"
)

/**
 * @ClassName ali
 * @Description TODO
 * @Author khr
 * @Date 2023/5/23 10:25
 * @Version 1.0
 */

var (
	AliPayClient *alipay.Client
	err          error
	ali_pri_key  = "MIIEpAIBAAKCAQEAicQOVpnQoVcfozTEOr1condnrWa3cpBYQClOufSN4lGMpvbq9ecEyyUODsBJadtjgOYoxOMShxq2QRDrA69Iz1Y2ZnpzKXXo24CoEoLw7yTER0At/+I1LHafFPL3SYzagoxapN/K4D0kY+FoXyqf+2lByWmvlGtVlkaV9sjwDvnLOo3VqFHL1TXK5LcK+UpmoEqVfRD9N1Ttoev8omMe9QDS4GPwrEPuwDJLjGkQZTYM+d8UHZ0Di6giiuUnAnWl1MdBwxHsmCkUUV8fkqzSiitEjFyeKMW/UEHXCuFMPsIq5xdK2tEa76z38kgsDmDURj+sEKCDbMZvERHXHq78mwIDAQABAoIBAEcUpyk7l4+HOkWk9hIwndkdrpqjQseTflUsevgrHAHHfcCv6a8SkUCXT9eAkuBRV9er6SEc3/RhbePIbNmr2O9RViQtzbl4orqOeSmD8fgRikwQ6yr5deIJGi9e5QRH7n4pGKO07CIiqeH27Tkc7wpy1oSrSPJVJwWwSbPZHTM8IOECngC1wEWgpAVhueCG/sZiSoPPGUic1tVd3aS2FRvdlPE0ZWyvSF4W/zlRhyY59UUuWtfbVzuYS7MPLCObv2TE970yOwswo9CAZhoOOTJFgqEuQL9ioEP96KaUIfCJ8wvmaUBFOd4s1ou0O0NmmQ6jIyc8UDUQBrIp0ttXVjECgYEAx22s//lmLonowOZjxNLMOSljiPkYbZ86Amwhjiew1nUhYN4RvtzFoFFksq3WMB6OXnQBiLb0uG3vfFmqFqTwQkCnLrHQ/X6islhF6UqkyBqvGQifj0bqLRnb8gXGz/TjXkJCXMLsjOo1bfvQmU/NBlEAhfoYn1YLL8Z0dp117I0CgYEAsNh+5+s/IN1dl253YJnV3kNeDjXrKtvg1u8S7BMwrAaonH2YJO6S5/03GCH2q9PAcTDJis8Sp8LE1/PioKjOZPudo/EITWgFjMUV1LSA97tW4mfLo90fg6/c8tcwIlePKEwms5DPef5Uc0Llcuth8Ig9i5wo9P3aF7WxT5FNR8cCgYEAj5F5Vd3pnd9SXGx/rpZCx3PwYA9TcreKP2wwy/Hu8LTqDp5QECNHcp6l66wR4hpdS8ofwJhVnOAn5FF4jUy4WjnJIWiJl7Su082Qpt5BunzbSR5YIAFhXI6dNKLL+bHGbXkt5TG+scN6K295QKWeZ8mwosLlLu/2pbIs7ad12ZECgYEAgMw8qUZxRMtUpbyjnyyLUgR4lRr5+s4HVCLtAhj74t46oTbrv0Iupl2Kab4avIxNZWLl9n3YFWzKFooerWokX/HNnyAmLtIq8Jp9ytvn7gV4Qw7bhq2+jRdhcU/+U5S3w96qdS9rnGr6MLQxDmCWhSuEv5BtV/kmhQwkZlHqGfMCgYAsxw3IqGmknC09VE9laRZ+L72lXnnZqvSYtQbIVV3RKFIKc9C1lSz8R81o4By2xg72T2jGkNfJ9L/gOYrS/N0pILP0TLl+dM3jsppnyX+DorZvxrVxOvMuGunHwIozvs2j+AqgPyiYfc2a0w8IpxqGD4dK96PezwWkr8R3S5kjFQ=="
	ali_app_id   = "2021000122694223"
)

func aliPayMiddleware(c *gin.Context) {
	c.Set("alipayClient", AliPayClient)
	c.Next()
}
func Routers(e *gin.Engine) {
	aliGroup := e.Group("/ali")
	aliGroup.Use(aliPayMiddleware)
	aliGroup.GET("/alipay", HandleAliPayment) // 支付请求
	aliGroup.GET("/alinotify", AliPayNotify)  // 生成支付链接并跳转至支付页面
}
func init() {
	// 创建支付宝客户端
	AliPayClient, err = alipay.New(ali_app_id, ali_pri_key, false)
	if err != nil {
		fmt.Printf("错误出现了：%s", err)
		errors.New("接入支付宝失败")
	}
	AliPayClient.LoadAppPublicCertFromFile("conf/appPublicCert.crt")       // 加载应用公钥证书
	AliPayClient.LoadAliPayRootCertFromFile("conf/alipayRootCert.crt")     // 加载支付宝根证书
	AliPayClient.LoadAliPayPublicCertFromFile("conf/alipayPublicCert.crt") // 加载支付宝公钥证书
}

// 处理支付请求
func HandleAliPayment(c *gin.Context) {
	// 创建支付宝客户端
	AliPayClient, err = alipay.New(ali_app_id, ali_pri_key, false)
	if err != nil {
		c.String(http.StatusInternalServerError, "支付请求失败")
		return
	}
	AliPayClient.LoadAppPublicCertFromFile("conf/appPublicCert.crt")       // 加载应用公钥证书
	AliPayClient.LoadAliPayRootCertFromFile("conf/alipayRootCert.crt")     // 加载支付宝根证书
	AliPayClient.LoadAliPayPublicCertFromFile("conf/alipayPublicCert.crt") // 加载支付宝公钥证书
	// 构建支付请求参数
	param := alipay.TradePagePay{}
	param.NotifyURL = "https://home.firefoxchina.cn/"
	param.Subject = "测试 公钥证书模式-这是一个gin订单"
	template := "2006-01-02 15:04:05"
	param.OutTradeNo = time.Now().Format(template)
	param.TotalAmount = "0.1"
	param.ProductCode = "FAST_INSTANT_TRADE_PAY"

	var redirectUrl *url.URL
	// 生成支付链接
	redirectUrl, err = AliPayClient.TradePagePay(param)
	if err != nil {
		// 处理错误
		c.String(500, "支付请求失败")
		return
	}
	fmt.Println(redirectUrl)
	// 重定向到支付宝支付页面
	//c.Redirect(http.StatusOK, redirectUrl.String())
	c.JSON(http.StatusOK, redirectUrl)

}
func AliPayNotify(c *gin.Context) {
	req := c.Request
	req.ParseForm()
	ok, _ := AliPayClient.VerifySign(req.Form)

	fmt.Println(ok)
	//处理订单逻辑关系
	fmt.Println(req.Form)

	c.String(http.StatusOK, "支付成功")
}
