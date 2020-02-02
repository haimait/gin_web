package router

import (
	"github.com/gin-gonic/gin"
	"github.com/haimait/gin_web/api/controller"
)

func Test(r *gin.Engine)  {
	r.GET("/test/response", controller.Response) //测试返参
}