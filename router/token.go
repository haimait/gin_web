package router

import (
	"github.com/gin-gonic/gin"
	"github.com/haimait/gin_web/api/controller"
)

func Token(r *gin.Engine)  {
	r.GET("/token/create", controller.CreateToken)
}