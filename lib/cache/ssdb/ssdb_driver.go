package ssdb

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ssdb/gossdb/ssdb"
	"github.com/haimait/gin_web/lib/cache"
	"strconv"
	"strings"
	"time"
)

type Cache struct {
	conn     *ssdb.Client
	connInfo []string
	key      string
}

var (
	DefaultKey = "cacheRedis"
)

//创建cache实例
func NewCacheSsdb() cache.Cache {
	return &Cache{key: DefaultKey}
}

// 拼接前缀,可能传入字符串或字符串切片，返回也得返回字符串或字符串切片
func (ss *Cache) join(originKey string) string {
	return fmt.Sprintf("%s:%s", ss.key, originKey)
}

// 批量拼接
func (ss *Cache) sliceJoin(originKeys []string) []string {
	var retKeys []string
	for i := 0; i < len(originKeys); i++ {
		retKeys = append(retKeys, ss.join(originKeys[i]))
	}
	return retKeys
}

// 获取单条
func (ss *Cache) Get(key string) interface{} {
	key = ss.join(key)
	v, err := ss.conn.Get(key)
	if err != nil {
		return nil
	}
	return v
}

// 获取多条
func (ss *Cache) GetMulti(keys []string) []interface{} {
	keys = ss.sliceJoin(keys)
	var values []interface{}
	size := len(keys)
	if ss.conn == nil {
		// 异常填充
		if err := ss.connectInit(); err != nil {
			for i := 0; i < size; i++ {
				values = append(values, err)
			}
			return values
		}
	}

	// res的返回结果格式为[ok key1 value1 key2 value2 ....]，例：[ok wang 112358 xian 223469]
	res, err := ss.conn.Do("multi_get", keys)
	resSize := len(res)
	// 需取下标1之后的偶数下标的值
	if err == nil {
		for i := 1; i < resSize; i += 2 {
			values = append(values, res[i+1])
		}
		return values
	}

	// 异常填充
	for i := 0; i < size; i++ {
		values = append(values, err)
	}
	return values

}

// 存入一条数据
func (ss *Cache) Put(key string, val interface{}, timeout time.Duration) error {
	key = ss.join(key)
	v, ok := val.(string)
	if !ok {
		return errors.New("cache:ssdb put value need string")
	}
	var (
		res []string
		err error
	)
	if int64(timeout/time.Second) < 0 {
		res, err = ss.conn.Do("set", key, v)
	} else {
		res, err = ss.conn.Do("setx", key, v, int64(timeout/time.Second))
	}
	if err != nil {
		return err
	}
	if len(res) == 2 && res[0] == "ok" {
		return nil
	}
	return errors.New("cache:ssdb response isn't ok")
}

// 删除多个
func (ss *Cache) DelMulti(keys []string) error {
	keys = ss.sliceJoin(keys)
	_, err := ss.conn.Do("multi_del", keys)
	return err
}

// 自增
func (ss *Cache) Inc(key string, num ...int) error {
	key = ss.join(key)
	var inc int
	if num == nil {
		inc = 1
	} else {
		inc = num[0]
	}

	_, err := ss.conn.Do("incr", key, inc)
	return err
}

//自减
func (ss *Cache) Dec(key string, num ...int) error {
	key = ss.join(key)
	var dec int
	if num == nil {
		dec = -1
	} else {
		dec = -1 * num[0]
	}

	_, err := ss.conn.Do("incr", key, dec)
	return err
}

// 根据key删除值
func (ss *Cache) Del(key string) error {
	key = ss.join(key)
	_, err := ss.conn.Del(key)
	return err
}

// 判断是否存在
func (ss *Cache) IsExist(key string) bool {
	key = ss.join(key)
	resp, err := ss.conn.Do("exists", key)
	if err != nil {
		return false
	}
	if len(resp) == 2 && resp[1] == "1" {
		return true
	}
	return false
}

// 通配删除, 数量多可能存在性能问题
func (ss *Cache) Dels(key string) error {
	// 判断字符串是否被包含
	keyStart, keyEnd, limit := "", "", 50
	res, err := ss.Scan(keyStart, keyEnd, limit)
	for err == nil {
		size := len(res)
		if size == 1 {
			return nil
		}
		var keys []string
		for i := 1; i < size; i += 2 {
			if strings.Contains(res[i], key) {
				keys = append(keys, res[i])
			}
		}
		_, e := ss.conn.Do("multi_del", keys)
		if e != nil {
			return e
		}
		keyStart = res[size-2]
		res, err = ss.Scan(keyStart, keyEnd, limit)
	}
	return err
}

// 删除所有
func (ss *Cache) DelAll() error {
	keyStart, keyEnd, limit := "", "", 50
	res, err := ss.Scan(keyStart, keyEnd, limit)
	for err == nil {
		size := len(res)
		if size == 1 {
			return nil
		}
		var keys []string
		for i := 1; i < size; i += 2 {
			keys = append(keys, res[i])
		}
		_, e := ss.conn.Do("multi_del", keys)
		if e != nil {
			return e
		}
		keyStart = res[size-2]
		res, err = ss.Scan(keyStart, keyEnd, limit)
	}
	return err
}

// 查看ssdb上的所有key
func (ss *Cache) Scan(keyStart string, keyEnd string, limit int) ([]string, error) {
	res, err := ss.conn.Do("scan", keyStart, keyEnd, limit)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// 注入驱动配置，生成连接
func (ss *Cache) Connect(config string) error {
	var conf map[string]string
	_ = json.Unmarshal([]byte(config), &conf)
	if _, ok := conf["conn"]; !ok {
		return errors.New("cache: ssdb config need conn key")
	}
	// 注入配置
	ss.connInfo = strings.Split(conf["conn"], ";")
	// 单例
	if ss.conn == nil {
		if err := ss.connectInit(); err != nil {
			return err
		}
	}
	return nil
}

// 初始化ssdb连接
func (ss *Cache) connectInit() error {
	var err error
	connInfoArr := strings.Split(ss.connInfo[0], ":")
	host := connInfoArr[0]
	if port, e := strconv.Atoi(connInfoArr[1]); e != nil {
		return e
	} else {
		ss.conn, err = ssdb.Connect(host, port)
	}
	return err
}

// 初始化，将ssdb的Cache结构的指针返回给cache的注册中心
func init() {
	cache.Register("ssdb", NewCacheSsdb)
}
