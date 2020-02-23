package conf

import (
	"fmt"
	"github.com/go-ini/ini"
)

var (
	Cfg = new(AppConf)
)

var Server string

func init() {
	// 0. 加载配置文件
	err := ini.MapTo(Cfg, "./conf/config.ini")
	if err != nil {
		fmt.Printf("load ini failed, err:%v\n", err)
		return
	}
	serverHost := Cfg.ServerConf.HTTP_HOST + ":" + Cfg.ServerConf.HTTP_PORT
	Server = serverHost
}

type AppConf struct {
	ServerConf   `ini:"server"`
	DatabaseConf `ini:"database"`
	JwtConf      `ini:"jwt"`
	MysqlConf    `ini:"mysql"`
	RedisConf    `ini:"redis"`
}
type ServerConf struct {
	RUN_MODE      string `ini:RUN_MODE`
	ENV           string
	HTTP_HOST     string
	HTTP_PORT     string
	READ_TIMEOUT  int
	WRITE_TIMEOUT int
}

type DatabaseConf struct {
	DRIVER_TYPE string `ini:DRIVER_TYPE`
	USER        string
	PASSWORD    string
	HOST        string
	DBNAME      string
	CHARSET     string
}

type MysqlConf struct {
	TABLE_PREFIX   string `ini:Driver`
	MAX_IDLE_CONNS int
	MAX_OPEN_CONNS int
	CHAN_MAX_SIZE  int
}

type JwtConf struct {
	PAGE_SIZE  int
	JWT_SECRET string
}
type RedisConf struct {
	HOST string
}
