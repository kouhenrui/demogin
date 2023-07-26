package main

import (
	"ali/alipay"
	"ali/wxpay"
	"github.com/gin-gonic/gin"
)

var (
	//AliPayClient *alipay.Client
	//err          error
	//ali_pri_key  = "MIIEpAIBAAKCAQEAicQOVpnQoVcfozTEOr1condnrWa3cpBYQClOufSN4lGMpvbq9ecEyyUODsBJadtjgOYoxOMShxq2QRDrA69Iz1Y2ZnpzKXXo24CoEoLw7yTER0At/+I1LHafFPL3SYzagoxapN/K4D0kY+FoXyqf+2lByWmvlGtVlkaV9sjwDvnLOo3VqFHL1TXK5LcK+UpmoEqVfRD9N1Ttoev8omMe9QDS4GPwrEPuwDJLjGkQZTYM+d8UHZ0Di6giiuUnAnWl1MdBwxHsmCkUUV8fkqzSiitEjFyeKMW/UEHXCuFMPsIq5xdK2tEa76z38kgsDmDURj+sEKCDbMZvERHXHq78mwIDAQABAoIBAEcUpyk7l4+HOkWk9hIwndkdrpqjQseTflUsevgrHAHHfcCv6a8SkUCXT9eAkuBRV9er6SEc3/RhbePIbNmr2O9RViQtzbl4orqOeSmD8fgRikwQ6yr5deIJGi9e5QRH7n4pGKO07CIiqeH27Tkc7wpy1oSrSPJVJwWwSbPZHTM8IOECngC1wEWgpAVhueCG/sZiSoPPGUic1tVd3aS2FRvdlPE0ZWyvSF4W/zlRhyY59UUuWtfbVzuYS7MPLCObv2TE970yOwswo9CAZhoOOTJFgqEuQL9ioEP96KaUIfCJ8wvmaUBFOd4s1ou0O0NmmQ6jIyc8UDUQBrIp0ttXVjECgYEAx22s//lmLonowOZjxNLMOSljiPkYbZ86Amwhjiew1nUhYN4RvtzFoFFksq3WMB6OXnQBiLb0uG3vfFmqFqTwQkCnLrHQ/X6islhF6UqkyBqvGQifj0bqLRnb8gXGz/TjXkJCXMLsjOo1bfvQmU/NBlEAhfoYn1YLL8Z0dp117I0CgYEAsNh+5+s/IN1dl253YJnV3kNeDjXrKtvg1u8S7BMwrAaonH2YJO6S5/03GCH2q9PAcTDJis8Sp8LE1/PioKjOZPudo/EITWgFjMUV1LSA97tW4mfLo90fg6/c8tcwIlePKEwms5DPef5Uc0Llcuth8Ig9i5wo9P3aF7WxT5FNR8cCgYEAj5F5Vd3pnd9SXGx/rpZCx3PwYA9TcreKP2wwy/Hu8LTqDp5QECNHcp6l66wR4hpdS8ofwJhVnOAn5FF4jUy4WjnJIWiJl7Su082Qpt5BunzbSR5YIAFhXI6dNKLL+bHGbXkt5TG+scN6K295QKWeZ8mwosLlLu/2pbIs7ad12ZECgYEAgMw8qUZxRMtUpbyjnyyLUgR4lRr5+s4HVCLtAhj74t46oTbrv0Iupl2Kab4avIxNZWLl9n3YFWzKFooerWokX/HNnyAmLtIq8Jp9ytvn7gV4Qw7bhq2+jRdhcU/+U5S3w96qdS9rnGr6MLQxDmCWhSuEv5BtV/kmhQwkZlHqGfMCgYAsxw3IqGmknC09VE9laRZ+L72lXnnZqvSYtQbIVV3RKFIKc9C1lSz8R81o4By2xg72T2jGkNfJ9L/gOYrS/N0pILP0TLl+dM3jsppnyX+DorZvxrVxOvMuGunHwIozvs2j+AqgPyiYfc2a0w8IpxqGD4dK96PezwWkr8R3S5kjFQ=="
	//ali_app_id   = "2021000122694223"

	// 设置微信支付参数
	wx_appID     = "YOUR_APP_ID"                    // 微信支付的App ID
	wx_mchID     = "YOUR_MCH_ID"                    // 商户号
	wx_apiKey    = "YOUR_API_KEY"                   // API密钥
	wx_notifyURL = "http://localhost:8080/callback" // 支付回调通知URL

)

type Option func(engine *gin.Engine)

var options = []Option{}

func Include(opts ...Option) {
	options = append(options, opts...)
}
func main() {
	Include(alipay.Routers, wxpay.Routers)
	r := gin.Default()

	for _, y := range options {
		y(r)
	}

	//r.Use(aliPayMiddleware) //挂载全局支付宝

	r.Run(":8080")
}
