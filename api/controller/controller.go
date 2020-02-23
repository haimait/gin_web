package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/haimait/gin_web/lib/apilog"
)

func getRequestBody(context *gin.Context, key string,s interface{}) error { //获取request的body
	body, _ := context.Get(key) //转换成json格式
	reqBody, _ := body.(string)
	apilog.DebugLog("--------------", reqBody)
	decoder := json.NewDecoder(bytes.NewReader([]byte(reqBody)))
	decoder.UseNumber() //作为数字而不是float64
	err := decoder.Decode(&s)
	fmt.Println(err)
	return err
}

// 获取post接口参数
func GetPostParams(ctx *gin.Context) (map[string]interface{}, error) {
	params := make(map[string]interface{})
	err := getRequestBody(ctx, "json",&params)
	return params, err
}

// 获取post接口参数
func GetPostParamsStruct(ctx *gin.Context, structName interface{}) error {
	err := getRequestBody(ctx, "json",&structName)
	return err
}

// 获取中间件里设置的结构体
func GetThisUserInfo(ctx *gin.Context,structName interface{}) error {
	body, _ := ctx.Get("userInfo") //转换成json格式
	err := json.Unmarshal(body.([]byte),&structName)
	//str := fmt.Sprintf("%s", body)
	//err := json.Unmarshal([]byte(str),&structName)
	//err := json.Unmarshal(body.([]byte),&structName)
	return err
}
