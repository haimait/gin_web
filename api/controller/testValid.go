package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/haimait/gin_web/api/model"
	"github.com/haimait/gin_web/lib/httpext"
	"log"
)


func TestValid1(ctx *gin.Context)  {
	log.Println("TestValid1")

	//获取当前用户
	user := model.User{}
	err := GetThisUserInfo(ctx,&user)
	if err!=nil {
		fmt.Println("get middleware err:",err)
		httpext.ErrorExt(ctx,400,"get middleware err")
	}
	fmt.Println(user.ID)
	httpext.Success(ctx)
}

