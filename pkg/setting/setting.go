package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg *ini.File

	RunMode      string
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	PageSize     int
	JwtSecret    string
)

func init() {
	var err error
	Cfg, err = ini.Load("config/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'config/app.ini': %s", err)
	}
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
	LoadServer()
	LoadApp()
}

func LoadServer() {
	sec, err := GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %s", err)
	}

	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
	sec, err := GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %s", err)
	}

	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}

func env(section string) string {
	return section + "_" + Cfg.Section("").Key("Env").MustString("test")
}

func GetSection(section string) (*ini.Section, error) {
	return Cfg.GetSection(env(section))
}

func GetInitValueByKey(key, valuekey string) (keyExsts int, value string) {
	cronCfg, err := GetSection(key)
	keyExsts = 1
	if err != nil {
		keyExsts = 0
		log.Fatal(2, "Fail to get section 'oss': %v", err)
	}
	value = cronCfg.Key(valuekey).String()
	return keyExsts, value
}
