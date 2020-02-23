package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/haimait/gin_web/api/model"
	"github.com/haimait/gin_web/lib/apilog"
	"github.com/haimait/gin_web/lib/httpext"
	"github.com/unknwon/com"
	"net/http"
)


func CreateUser(ctx *gin.Context)  {
	var userMod = model.User{
		Age:0,
	}
	res := userMod.CreateUser()
	fmt.Println("userMod333:",userMod)
	if(userMod.ID == 0 && !res){
		httpext.ErrorExt(ctx,http.StatusBadRequest,"添加失败")
	}
	httpext.Success(ctx)
}

func GetUserDetail(ctx *gin.Context)  {
	id := com.StrTo(ctx.Query("id")).MustInt64()
	var userMod = model.User{ID:id}
	userMod.GetUserDetailById()
	fmt.Println("userMod3334")
	fmt.Println(userMod)
	httpext.SuccessExt(ctx,userMod)
}

func GetUserList(ctx *gin.Context)  {
	var userMod = model.User{}
	userlist := userMod.GetUserGetUserList()
	apilog.Log("user","getuserlist",userlist)
	fmt.Println("userlist1")
	fmt.Printf("%#v \n",userlist)
	httpext.SuccessExt(ctx,userlist)
}

