package httpext

import (
	"sync"
)

var (
	secLimitMgr = &SecLimitMap{
		UserLimitMap: make(map[int]*SecLimit, 0),
	}
)

/**
*　所有用户访问次数的信息
*  author haima
 */
type SecLimitMap struct {
	UserLimitMap map[int]*SecLimit
	lock         sync.Mutex
}

// func antiSpam(req *SecRequest) (err error) {
// 	secLimitMgr.lock.Lock()
// 	secLimit, ok := secLimitMgr.UserLimitMap[req.UserId]
// 	if !ok {
// 		secLimit = &SecLimit{}
// 		secLimitMgr.UserLimitMap[req.UserId] = secLimit
// 	}
// 	count := secLimit.Count(req.AccessTime.Unix())
// 	secLimitMgr.lock.Unlock()
// 	if count > 5 { //每秒访问超过５次报错　开发环境时次数限制数写在配置里
// 		err = fmt.Errorf("invalid request")
// 		return
// 	}
// 	return
// }

type SecLimit struct {
	count   int
	curTime int64
}

/**
*　返回访问次数
*  nowTime 用户当前请求接口的时间戳
*  author haima
 */
func (p *SecLimit) Count(nowTime int64) (curCount int) {
	if p.curTime != nowTime {
		p.count = 1
		p.curTime = nowTime
		curCount = p.count
		return
	}
	p.count++
	curCount = p.count
	return
}

/**
*　检测访问次数　返回访问次数
*  nowTime 用户当前请求接口的时间戳
*  author haima
 */
func (p *SecLimit) Check(nowTime int64) int {
	if p.curTime != nowTime {
		return 0
	}
	return p.count
}
