package apilog

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"time"
)

/*自定义日志文件*/
// func DebugLog(args ...interface{}) error {
// 	var floderLog = "logs"
// 	logName := floderLog + "/debug_log." + time.Now().Format("2006-01-02") + ".log"
// 	if _, err := os.Stat(logName); err != nil {
// 		if os.IsExist(err) {
// 			os.Create(logName)
// 		}
// 	}

// 	logFile, err := os.OpenFile(logName, os.O_APPEND|os.O_RDWR|os.O_CREATE, os.ModePerm)
// 	if err != nil {
// 		fmt.Print("打开文件失败")
// 		return err
// 	}
// 	defer logFile.Close()
// 	tm := time.Now().Format("2006-01-02 15:04:05")
// 	logBody := fmt.Sprintf("%s ====>   %+v\r\n", tm, args)
// 	logFile.WriteString(logBody)
// 	return nil
// }
/*自定义日志文件*/
func DebugLog(args ...interface{}) error {
	var floderLog = "logs"
	logName := floderLog + "/debug_log." + time.Now().Format("2006-01-02") + ".log"
	if _, err := os.Stat(floderLog); err != nil {
		if !os.IsExist(err) {
			os.MkdirAll(floderLog, os.ModePerm)
		}
	}

	logFile, err := os.OpenFile(logName, os.O_APPEND|os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Print("打开文件失败")
		return err
	}
	defer logFile.Close()
	tm := time.Now().Format("2006-01-02 15:04:05")
	logBody := fmt.Sprintf("%s star============================================>\r\n", tm)
	logFile.WriteString(logBody)
	logBody = fmt.Sprintf("%+v\r\n", args)
	logFile.WriteString(logBody)
	logBody = fmt.Sprintf("\r\n")
	logFile.WriteString(logBody)
	return nil
}

//DebugLog2 自定义日志文件
/*
*　dir　文件夹名字
*　fileName　文件名字
*　args　要写入的数据
 */
func Log(dir string, fileName string, args ...interface{}) error {
	var floderLog = "logs"
	if fileName != "" {
		floderLog = floderLog + "/" + dir
	}
	logName := floderLog + "/" + fileName + time.Now().Format("2006-01-02") + ".log"
	if _, err := os.Stat(floderLog); err != nil {
		if !os.IsExist(err) {
			os.MkdirAll(floderLog, os.ModePerm)
		}
	}

	logFile, err := os.OpenFile(logName, os.O_APPEND|os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("打开文件失败 err:", err)
		return err
	}
	defer logFile.Close()
	tm := time.Now().Format("2006-01-02 15:04:05")

	//方法一：os包
	// logBody := fmt.Sprintf("%s star============================================>\r\n", tm)
	// logFile.WriteString(logBody)
	// logBody = fmt.Sprintf("%+v\r\n", args)
	// logFile.WriteString(logBody)
	// logBody = fmt.Sprintf("\r\n")
	// logFile.WriteString(logBody)

	//方法二：bufio包
	writer := bufio.NewWriter(logFile)
	logBody := fmt.Sprintf("%s star============================================>\r\n", tm)
	writer.WriteString(logBody) //写入缓存
	logBody = fmt.Sprintf("%+v\r\n", args)
	writer.WriteString(logBody)
	logBody = fmt.Sprintf("\r\n")
	writer.WriteString(logBody)

	writer.Flush() //从缓存写入文件
	return nil
}

//GetFileInfo 获取文件路径，方法名　行号
func GetFileInfo() {
	funName, fileName, lineNum, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("get GetFileInfo info fail")
	}
	str := fmt.Sprintf("funName:%s  fileName:%s  line:%d", fileName, runtime.FuncForPC(funName).Name(), lineNum)
	Log("api", "apislice", str)
}
