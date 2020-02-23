package router

import (
	"github.com/gin-gonic/gin"
	"github.com/haimait/gin_web/api/controller"
)

func User(r *gin.Engine)  {
	r.GET("/user/create_user", controller.CreateUser)
	r.GET("/user/get_user_detail", controller.GetUserDetail)
	r.GET("/user/get_user_list", controller.GetUserList)
}