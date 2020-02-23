package main

import (
	"github.com/gin-gonic/gin"
	"github.com/haimait/gin_web/conf"
	"github.com/haimait/gin_web/router"
)


func main() {
	r := gin.Default()
	////初始化日志库
	//err := apilog.InitLogger()
	//if err!=nil{
	//	logs.Error("init logger fialed, err:%v",err)
	//	panic(fmt.Sprintf("init logger failed, err:%v",err))
	//}
	//logs.Debug("init logger success.")

	//hsot := conf.Cfg.ServerConf.HTTP_HOST
	//port := conf.Cfg.ServerConf.HTTP_PORT
	//fmt.Println(hsot)
	//fmt.Println(port)
	router.Init(r)
	_ = r.Run(conf.Server)

}
