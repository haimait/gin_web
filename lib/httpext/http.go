package httpext

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/haimait/gin_web/lib/apilog"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// @desc 通过上下文获取header中指定key的内容
func GetHeaderByName(ctx *gin.Context, key string) string {
	return ctx.Request.Header.Get(key)
}

// @desc 通过上下文获取body内容并将内容写到指定key中
func SetBodyJson(context *gin.Context, key string) {
	body, _ := ioutil.ReadAll(context.Request.Body)
	apilog.DebugLog("当前请求数据内容:", string(body[:]))
	fmt.Println("当前请求的数据:", string(body[:]))
	context.Set(key, string(body[:]))
}

// Get ... 发送请求 ...
// url：         请求地址
// response：    请求返回的内容
func SetGet(url string) string {

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	return result.String()
}

//GetWithHeard ... 带header头的GET请求　参数先拼到地址里
/*
	url要访问的url　例：https://test.ibanana.club/task/p_task_list2?user_id=100000072&linkbook_id=bo10g2f1qr3u2r41mp40&page=1&row=10
	hearderParam hearder头信息
*/
func SetGetWithHeard(url string, hearderParam map[string]string) string {

	client := &http.Client{}
	//生成要访问的url
	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("err11->", err)
	}
	//增加header选项 添加请求头
	if len(hearderParam) > 0 {
		for k, v := range hearderParam {
			reqest.Header.Add(k, v)
		}
	}
	if err != nil {
		panic(err)
	}
	//处理返回结果
	response, _ := client.Do(reqest)
	defer response.Body.Close()

	// 添加请求头

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("err22->", err)
	}
	// fmt.Println(string(data))
	return string(data)
}

//禁用KeepAlive的client  //适用于请求特别频繁每几秒请求一次
// 超时时间：10秒
//var (
//	client = http.Client{
//		Transport:     &http.Transport{
//			DisableKeepAlives:false,
//			},
//		Timeout: 10 * time.Second,
//	}
//)
//GetWithHeard ... 带header头的GET请求　带hearder头　参数分开传过来，函数里自己拼接
/*
	apiUrl 要访问的url 　例：https://test.ibanana.club/task/p_task_list2
	setParam　要带的参数
	hearderParam hearder头信息
*/
func GetWithHeard2(apiUrl string, setParam map[string]string, hearderParam map[string]string) string {

	//生成要访问的url

	//增加header选项 添加请求头
	if len(setParam) > 0 {
		// URL param
		data := url.Values{}
		for k, v := range setParam {
			data.Set(k, v)
		}
		u, err := url.ParseRequestURI(apiUrl)
		if err != nil {
			fmt.Printf("parse url requestUrl failed,err:%v\n", err)
		}
		u.RawQuery = data.Encode() // URL encode
		apiUrl = u.String()
		fmt.Println(apiUrl)
	}

	//提交请求
	request, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		fmt.Println("GET query err 11->", err)
	}
	//增加header选项 添加请求头
	if len(hearderParam) > 0 {
		for k, v := range hearderParam {
			request.Header.Add(k, v)
		}
	}

	//禁用KeepAlive的client  //适用于请求不是特别频繁，用完就关闭了，不然连接多了就会假死　如：每天请求２次
	tr := &http.Transport{
		DisableKeepAlives: false,
	}
	// 超时时间：10秒
	client := http.Client{
		Transport: tr,
		Timeout:   10 * time.Second,
	}

	//处理返回结果
	response, _ := client.Do(request)
	defer response.Body.Close()

	// 添加请求头

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("ReadAll response.Body err22->", err)
	}
	// fmt.Println(string(data))
	return string(data)
}

// Post ... 发送POST请求 ...
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// content：     请求放回的内容
func SetPost(url string, data interface{}, contentType string) string {

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result)
}

// SendPostWithHeard ...   发送POST请求 ...
// url：         请求地址
// data：        POST请求提交的数据
// hearderParam  设置header头信息
// contentType： 请求体格式，如：json
// content：     请求放回的内容
func SendPostWithHeard(url string, data interface{}, hearderParam map[string]string, contentType string) string {
	bytesData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}

	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}
	//增加header选项 添加请求头
	if len(hearderParam) > 0 {
		for k, v := range hearderParam {
			request.Header.Set(k, v)
		}
	}

	var contentTypeParam = map[string]string{
		"json": "application/json;charset=UTF-8",
	}
	request.Header.Set("Content-Type", contentTypeParam[contentType])
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}

	return string(respBytes)
}
