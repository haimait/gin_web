package functionext

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/haimait/gin_web/lib/apilog"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//Md5Encode md5加密
func Md5Encode(encodeString string) string {
	h := md5.New()
	h.Write([]byte(encodeString))
	return hex.EncodeToString(h.Sum(nil)) // 输出加密结果
}

//MD5Encode md5加密后转大写
func MD5Encode(encodeString string) string {
	return strings.ToUpper(Md5Encode(encodeString))
}

// 生成随机数
func MtRand(min, max int) int {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(max-min) + min
	return randNum
}

//生成随机字符串
func GetRandomString(num int) string {
	str := "123456789abcdefghijklmnpqrstuvwxyzABCDEFGHIJKLMNPQRSTUVWXYZ"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < num; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func GetVisitCount(url string) map[string]interface{} {
	resp, err := http.Get(url)
	resu := map[string]interface{}{
		"gatewaySuccessCount": 0,
		"gatewayFailureCount": 0,
		"gatewayDayCount":     0,
		"gatewayMinuteCount":  0,
		"lastUpdateTime":      time.Now().Format("15:04:05"),
	}
	if err != nil {
		return resu
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	result := make(map[string]interface{})
	err = json.Unmarshal(body, &result)
	if err != nil {
		return resu
	}
	return result
}

// 关闭后端服务
func StopGatewayService(port string, force bool) bool {
	command := `netstat -tlnp | grep ` + port + `| awk '{print $7}'`
	cmd := exec.Command("/bin/bash", "-c", command)

	if out, err := cmd.Output(); err != nil {
		panic(err)
	} else {
		id := strings.Split(string(out), "/")[0]
		if force {
			command = `kill -9 ` + id
		} else {
			command = `kill -HUP ` + id
		}
		cmd = exec.Command("/bin/bash", "-c", command)
		if _, err := cmd.Output(); err != nil {
			fmt.Println(port)
			return false
		}
		return true
	}
}

func GetGatewayServiceStatus(port string) bool {
	command := `netstat -tlnp | grep ` + port + `| awk '{print $7}'`
	cmd := exec.Command("/bin/bash", "-c", command)

	if out, err := cmd.Output(); err != nil {
		panic(err)
	} else {
		id := strings.Split(string(out), "/")[0]
		if id != "" {
			return true
		} else {
			return false
		}
	}
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func CheckFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

// RedPackage 发红包
// count: 红包数量 整数按指定份数平均分成若干份
// money: 红包金额（单位：分) 10 5
func RedPackage(money, number int) (moneySlice []int) {
	for i := 0; i < number; i++ {
		m := randomMoney(number-i, money)
		moneySlice = append(moneySlice, m)
		money -= m
	}
	return
}

// randomMoney 随机红包
// remainCount: 剩余红包数
// remainMoney: 剩余红包金额（单位：分)
func randomMoney(remainCount, remainMoney int) int {
	if remainCount == 1 {
		return remainMoney
	}
	rand.Seed(time.Now().UnixNano())
	var min = 1
	if remainMoney == remainCount {
		return remainMoney / remainCount
	}
	max := remainMoney / remainCount
	if remainCount*2 > remainMoney {
		max = remainMoney / remainCount * 2
	}
	money := rand.Intn(max) + min
	return money
}

//BuildNick 编码昵称(编码带表情的字符串)
func BuildNick(nickName string) string {
	return base64.StdEncoding.EncodeToString([]byte(nickName))
}

//GetUserNick 获取解码昵称(解码带表情的字符串)
func GetUserNick(nickName string) string {

	unick, err := base64.StdEncoding.DecodeString(nickName)
	if err != nil {
		apilog.DebugLog("获取用户昵称失败", nickName, err)
		return nickName
	}
	//if len(unick)%4 != 0 {
	//	return nickName
	//}
	//fmt.Println(66666666)
	//fmt.Printf("%s", string(unick))
	return fmt.Sprintf("%s", string(unick))
}

//Contain 判断obj是否在target中，target支持的类型array,slice,map
func Contain(obj interface{}, target interface{}) bool {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true
		}
	}

	return false
}

//RutureWhitePhone 手机号白名单
func RutureWhitePhone() (phones []string) {
	phones = []string{
		"13811111111",
		"13811111112",
	}
	return
}

//判断是不是手机格式
func ValidMobile(mobile string) bool {
	reg := `^1[3456789]\d{9}$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(mobile)
}

//PayType 支付类型: 1 金币支付 2:微信支付 3:支付宝支付 4:apple_pay
func PayType() []int {
	return []int{1, 2, 3, 4}
}

//MobileReplce 隐形手机号中间四位为*  phone 13811111112    mobile 138****1112
func MobileReplce(phone string) (mobile string) {
	str := []byte(phone)
	if len(str) == 11 {
		str[3], str[4], str[5], str[6] = 42, 42, 42, 42
		mobile = string(str)
	}
	return
}

//IsNil 判断是否为空接口
func IsNil(i interface{}) bool {
	defer func() {
		recover()
	}()
	vi := reflect.ValueOf(i)
	return vi.IsNil()
}

//根据出生日期 计算年龄
func GetAgeByBirthday(birthday string) (int, error) {
	//birthday 1995-01-01
	if birthday != "" {
		birthday := strings.Split(birthday, "-")

		if len(birthday) < 3 {
			return 0, errors.New("出生日期格式解析错误")
		}

		birYear, _ := strconv.Atoi(birthday[0])
		//birMonth, _ := strconv.Atoi(birthday[1])

		age := time.Now().Year() - birYear

		//if int(time.Now().Month()) < birMonth {
		//	age--
		//}

		return age, nil
	} else {
		return 0, errors.New("出生日期为空")
	}
}
