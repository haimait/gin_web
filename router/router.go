package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/haimait/gin_web/api/middleware"
	"github.com/haimait/gin_web/lib/apilog"
	"time"
)

func Init(r *gin.Engine)  {
	r.Use(func(c *gin.Context) {
		apilog.DebugLog("接口地址:", c.Request.RequestURI)
	})
	r.Use(middleware.Cors()) //解决跨域问题
	r.Use(myTime)            //统计时间
	Token(r)
	GlobalValid(r)
	Test(r)
}

//统计时间
func myTime(r *gin.Context){
	start :=time.Now()
	r.Next()
	since:=time.Since(start)
	fmt.Println("程序执行时间：",since)
}
// 全局非白名单路由验证
func GlobalValid(router *gin.Engine) {
	// token验证中间件
	router.Use(middleware.ValidToken)
}