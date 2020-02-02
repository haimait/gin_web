package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
