package text

import (
	"HelloGin/src/global"
	"HelloGin/src/pojo"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"time"
)

type Login struct {
	Name     string `from:"name" binding:"required"`
	Password string `from:"password" binding:"required"`
}

func Routers(e *gin.Engine) {

	testGroup := e.Group("/test")
	//{
	testGroup.GET("/get", getText)
	testGroup.POST("/post", postText)
	testGroup.POST("/chaos", chaos)
	//}

}
func getText(c *gin.Context) {
	res := global.NewResult(c)
	name := "我疯狂v表空间"

	copyContext := c.Copy()
	go func() {
		time.Sleep(3 * time.Second)

		log.Println("异步执行" + copyContext.Request.URL.Path)

	}()
	//return {
	//	n := name
	//}
	res.Success(name)
	return
	//c.JSON(http.StatusBadRequest, gin.H{
	//	"data": name,
	//})

}
func postText(c *gin.Context) {
	res := global.NewResult(c)
	var js Login
	if err := c.BindJSON(&js); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			res.Error(http.StatusBadRequest, err.Error())
			return
		}
		res.DiyErr(http.StatusBadRequest, global.Translate(errs))
		return
	}
	var test pojo.TestInt
	//go test.Write(js)
	f := pojo.Test{
		Name:     js.Name,
		Password: js.Password,
	}
	log.Println(f)
	test = &f
	b, t := test.Write(js)
	if !b {
		res.DiyErr(600, t)
		return
	}

	res.Success(t)
	return
}

func chaos(c *gin.Context) {
	req, _ := c.Get("request")
	fmt.Println("request", req)
	c.JSON(http.StatusOK, gin.H{
		"message": "hello gin",
		"request": req,
	})

}
