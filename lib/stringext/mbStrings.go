package stringext

import (
	"fmt"
	"strings"
	"unsafe"
)

// 字符串转成字符数组，通过修改结构实现
func Str2Bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// 字符数组转成字符串，通过修改结构实现
func Bytes2Str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// 统计字符串中各字符出现的次数
func WordStrCount(s string) map[string]int {
	var word string
	m := make(map[string]int)
	for i := 0; i < len(s); {
		word = s[i : i+1]
		v, ok := m[word]
		if ok != false {
			m[word] = v + 1
		} else {
			m[word] = 1
		}
		i += 1
	}
	return m
}

// 将字符串数组按照分隔符拼接成字符串
func Implode(arr []string, split string) string {
	return strings.Replace(strings.Trim(fmt.Sprint(arr), "[]"), " ", split, -1)
}

func ImplodeInt64(arr []int64, split string) string {
	return strings.Replace(strings.Trim(fmt.Sprint(arr), "[]"), " ", split, -1)
}

// 将字符串按照分隔符切割成数组
func Explode(str string, split string) []string {
	return strings.Split(str, split)
}
