package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/haimait/gin_web/conf"
	"github.com/haimait/gin_web/lib/httpext"
)


func Response(ctx *gin.Context)  {
	//time.Sleep(time.Second*5)
	fmt.Println(123)
	//httpext.SuccessExt(ctx,map[string]interface{}{
	//	"respeson":"test",
	//})
	httpext.SuccessExt(ctx,gin.H{
		"respeson":"test222",
	})
	//ctx.JSON(200,gin.H{
	//	"respeson":"test",
	//})
}

func TestLog(ctx *gin.Context)  {
	fmt.Println("接口地址11:", ctx.Request.RequestURI)
	fmt.Println("接口地址22:", ctx.Request.URL.Path)
	fmt.Println("Method:", ctx.Request.Method)
	fmt.Println("URL.User:", ctx.Request.URL.User)
	fmt.Println("header:", ctx.Request.Header)
	fmt.Println("token:", ctx.Request.Header.Get("token"))

	httpext.Success(ctx)
}

func GetConfig(ctx *gin.Context)  {
	DatabaseConf := conf.Cfg.DatabaseConf
	fmt.Printf("type：%T Value:%#v",DatabaseConf,DatabaseConf)
	httpext.SuccessExt(ctx,DatabaseConf)
	return
}