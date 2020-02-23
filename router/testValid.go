package router

import (
	"github.com/gin-gonic/gin"
	"github.com/haimait/gin_web/api/controller"
)

func VestValid(r *gin.Engine)  {
	r.POST("/test/testvalid1", controller.TestValid1) //测试返参
}