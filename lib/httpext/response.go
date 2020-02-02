package httpext

import (
	"github.com/gin-gonic/gin"
	"github.com/haimait/gin_web/pkg/e"
	"net/http"
)

func Error(ctx *gin.Context, errCode int) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":  errCode,
		"msg": e.GetMsg(errCode),
	})
}

func ErrorExt(ctx *gin.Context, errCode int, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":  errCode,
		"msg": e.GetMsg(errCode),
		"data":   data,
	})
}
func Success(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":  e.SUCCESS,
		"msg": e.GetMsg(e.SUCCESS),
	})
}

func SuccessExt(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{ //转换成json格式
		"code":  e.SUCCESS,
		"msg": e.GetMsg(e.SUCCESS),
		"data":   data,
	})
}

func Proxy(ctx *gin.Context, res *http.Response) {
	contentLength := res.ContentLength
	contentType := res.Header.Get("Content-Type")
	extraHeaders := map[string]string{
		"Content-Disposition": `attachment; filename="invite.png"`,
	}
	ctx.DataFromReader(http.StatusOK, contentLength, contentType, res.Body, extraHeaders)
}
