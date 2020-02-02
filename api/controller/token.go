package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/haimait/gin_web/api/service"
	"github.com/haimait/gin_web/lib/httpext"
)

func CreateToken(ctx *gin.Context)  {
	token := service.CreateToken()

	httpext.SuccessExt(ctx,gin.H{
		"token":token,
	})
	//ctx.JSON(200,gin.H{
	//	"respeson":"test",
	//})
}
