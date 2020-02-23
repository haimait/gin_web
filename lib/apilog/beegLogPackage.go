package apilog

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"os"
	"time"
)

func convertLogLevel(level string) int {
	switch level {
	case "debug":
		return logs.LevelDebug
	case "warn":
		return logs.LevelWarn
	case "info":
		return logs.LevelInfo
	case "trace":
		return logs.LevelTrace
	}
	return logs.LevelDebug
}

func InitLogger() (err error) {
	var floderLog = "logs/beelog"
	logName := floderLog + "/debug_log." + time.Now().Format("2006-01-02") + ".log"
	if _, err := os.Stat(floderLog); err != nil {
		if !os.IsExist(err) {
			os.MkdirAll(floderLog, os.ModePerm)
		}
	}

	config := make(map[string]interface{})
	config["filename"] = logName
	config["level"] = convertLogLevel("debug")

	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("initLogger failed,marshal err:", err)
		return
	}
	logs.SetLogger(logs.AdapterFile, string(configStr))
	return
}

