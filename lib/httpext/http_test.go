package httpext

import (
	"encoding/json"
	"fmt"
)

type resData struct {
	Data   map[int]string `json:"data"`
	Errmsg string         `json:"errmsg"`
	Errno  int            `json:"errno"`
}
type resData2 struct {
	Data []interface{} //方法一
	// Data   []Datalist2　//方法二
	Errmsg string `json:"errmsg"`
	Errno  int    `json:"errno"`
}

// Datalist2 ...
type Datalist2 struct {
	ID       int64  `json:"id"`
	Position string `json:"position"`
}
type PostresData2 struct {
	Data   interface{} //方法一
	Errmsg string      `json:"errmsg"`
	Errno  int         `json:"errno"`
}

func GetTest() {
	var d resData2
	url := "https://api.ibanana.club/select/major/list_by_key?page=1&row=10&major_name=工程"
	res := SetGet(url)
	_ = json.Unmarshal([]byte(res), &d)
	fmt.Println(1111111111)
	fmt.Printf("type:%T value:%#v \n", d, d)
}

func SetGetWithHeardTest() {
	var d resData2
	url := "https://test.ibanana.club/task/p_task_list2?user_id=100000072&linkbook_id=bo10g2f1qr3u2r41mp40&page=1&row=10"
	var hearderParam = make(map[string]string)
	hearderParam["token"] = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkZXZpY2VfdXVpZCI6IjAwMDAwMDAwLTE5ZjgtZTVkOS0wMDAwLTAwMDA2NmRiNTNlMyIsInVzZXJfaWQiOjEwMDAwMDA3MiwibGlua19ib29rX2lkIjoiYm8xMGcyZjFxcjN1MnI0MW1wNDAiLCJwYXNzd29yZCI6IiIsImVuZF90aW1lIjoxNTgxNzM4MTc4LCJpc191cGRhdGVfdXNlcl9pbmZvIjowLCJwbGF0X2Zvcm0iOiJhcHAifQ.jYAI8axSHOE4Gh8yqeEsIGNjjMbM0uuI0LA1ROeAio8"
	res := SetGetWithHeard(url, hearderParam)
	_ = json.Unmarshal([]byte(res), &d)
	fmt.Println(1111111111)
	fmt.Printf("type:%T value:%#v \n", d, d)
}

func GetWithHeard2Test() {
	var d resData2
	url := "https://test.ibanana.club/task/p_task_list2"
	var hearderParam = make(map[string]string)
	hearderParam["token"] = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkZXZpY2VfdXVpZCI6IjAwMDAwMDAwLTE5ZjgtZTVkOS0wMDAwLTAwMDA2NmRiNTNlMyIsInVzZXJfaWQiOjEwMDAwMDA3MiwibGlua19ib29rX2lkIjoiYm8xMGcyZjFxcjN1MnI0MW1wNDAiLCJwYXNzd29yZCI6IiIsImVuZF90aW1lIjoxNTgxNzM4MTc4LCJpc191cGRhdGVfdXNlcl9pbmZvIjowLCJwbGF0X2Zvcm0iOiJhcHAifQ.jYAI8axSHOE4Gh8yqeEsIGNjjMbM0uuI0LA1ROeAio8"
	var setParam = make(map[string]string)
	setParam["user_id"] = "100000072"
	setParam["linkbook_id"] = "bo10g2f1qr3u2r41mp40"
	setParam["page"] = "1"
	setParam["row"] = "10"
	fmt.Println(setParam)
	res := GetWithHeard2(url, setParam, hearderParam)
	_ = json.Unmarshal([]byte(res), &d)
	fmt.Println(222222222)
	fmt.Printf("type:%T value:%#v \n", d, d)
}

func SendPostWithHeardTest() {
	var d PostresData2
	url := "https://test.ibanana.club/ad/get_detail"
	var hearderParam = make(map[string]string)
	hearderParam["token"] = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkZXZpY2VfdXVpZCI6IjAwMDAwMDAwLTE5ZjgtZTVkOS0wMDAwLTAwMDA2NmRiNTNlMyIsInVzZXJfaWQiOjEwMDAwMDA3MiwibGlua19ib29rX2lkIjoiYm8xMGcyZjFxcjN1MnI0MW1wNDAiLCJwYXNzd29yZCI6IiIsImVuZF90aW1lIjoxNTgxNzM4MTc4LCJpc191cGRhdGVfdXNlcl9pbmZvIjowLCJwbGF0X2Zvcm0iOiJhcHAifQ.jYAI8axSHOE4Gh8yqeEsIGNjjMbM0uuI0LA1ROeAio8"

	var dataParam = make(map[string]interface{})
	dataParam["ads_id"] = 38
	dataParam["ad_position_id"] = 1
	dataParam["ad_type"] = 1

	res := SendPostWithHeard(url, dataParam, hearderParam, "json")
	_ = json.Unmarshal([]byte(res), &d)
	fmt.Println(777777777777)
	fmt.Printf("type:%T value:%#v \n", d, d)
}
