package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/haimait/gin_web/lib/apilog"
	"github.com/haimait/gin_web/lib/httpext"
	"github.com/haimait/gin_web/lib/jwt"
	"github.com/haimait/gin_web/pkg/e"
)

// @author yuanbuyu
// @desc 验证中间件，校验客户端令牌是否生效

func ValidToken(ctx *gin.Context) {
	// 读取请求中header的认证信息
	token := httpext.GetHeaderByName(ctx, "token")
	//fmt.Println("token:",token)
	//apilog.DebugLog("当前获取的ｔｏｋｅｎ信息：" , token)
	jwtObj := jwt.NewJWT()
	c, err := jwtObj.ParseToken(token)
	if err != nil {
		httpext.Error(ctx, e.ERROR_AUTH)
		ctx.Abort()
	}
	if token == ""{
		apilog.DebugLog(e.ERROR_AUTH, token)
		ctx.Abort()
	}
	if token != "" {
		apilog.DebugLog("当前获取的ｔｏｋｅｎ信息：", token)
		//fmt.Printf("当前登录用户%#v\n", c)
		//apilog.DebugLog("当前用户访问接口地址:", ctx.Request.RequestURI)
		apilog.DebugLog("当前登录用户c", fmt.Sprintf("当前登录用户%#v", c))
		apilog.DebugLog("当前登录用户c", c.UserId)
		apilog.DebugLog("当前用户登录平台", c.PlateForm)

		//var u model.User
		//u.ID = c.UserId
		//resu , _ := json.Marshal(u)
		//ctx.Set("userInfo",resu)

	}
	ctx.Next()

}
