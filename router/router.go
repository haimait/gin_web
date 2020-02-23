package router

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"github.com/haimait/gin_web/api/middleware"
	"github.com/haimait/gin_web/lib/apilog"
	"github.com/haimait/gin_web/lib/httpext"
	"time"
)


func Init(r *gin.Engine)  {
	r.Use(func(c *gin.Context) {
		fmt.Println("request star==============================================================>")
		fmt.Printf("当前请求的接口地址 Method:%s   %s\n", c.Request.Method ,c.Request.RequestURI)
		apilog.DebugLog(fmt.Sprintf("当前请求的接口地址 Method:%s   %s", c.Request.Method ,c.Request.RequestURI))
	})
	r.Use(middleware.Cors()) //解决跨域问题
	Global(r)
	r.Use(myTime)            //统计时间
	initLog()
	User(r)
	Token(r)
	Test(r)
	GlobalValid(r)
	VestValid(r)
}

//统计时间
func myTime(r *gin.Context){
	start :=time.Now()
	r.Next()
	since:=time.Since(start)
	fmt.Println()
	fmt.Println("程序执行时间：",since)
}
//初始化日志库
func initLog(){
	err := apilog.InitLogger()
	if err!=nil{
		logs.Error("init logger fialed, err:%v",err)
		panic(fmt.Sprintf("init logger failed, err:%v",err))
	}
	logs.Debug("init logger success.")
}


// @desc 全局中间件,存放中间件
func Global(router *gin.Engine) {
	// 路由切分中间件
	router.Use(func(ctx *gin.Context) {
		httpext.SetBodyJson(ctx, "json")
	})
	router.Use(middleware.Logger())

}





// 全局非白名单路由验证
func GlobalValid(router *gin.Engine) {
	// token验证中间件
	router.Use(middleware.ValidToken)
}