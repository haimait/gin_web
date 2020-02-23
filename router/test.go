package router

import (
	"github.com/gin-gonic/gin"
	"github.com/haimait/gin_web/api/controller"
)

func Test(r *gin.Engine)  {
	r.GET("/test/response", controller.Response) //测试返参
	r.GET("/test/testlog", controller.TestLog) //测试返参
	r.GET("/test/getconfig", controller.GetConfig) //测试返参
}