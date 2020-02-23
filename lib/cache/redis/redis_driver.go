package redis

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/haimait/gin_web/lib/apilog"
	"github.com/haimait/gin_web/lib/cache"

	//"linkbook.com/LinkBookGo/lib/apilog"
	//"linkbook.com/LinkBookGo/lib/cache"
	"strconv"
	"strings"
	"time"
)

// 声明redis的cache结构，用来去实现cache的接口
type Cache struct {
	po       *redis.Pool
	connInfo string
	dbNum    int
	key      string
	password string
	maxIdle  int
}

// 定义默认的cache_key前缀，用来区分不同的服务
var (
	DefaultKey = "cacheRedis"
)

// 创建cache实例s
func NewCacheRedis() cache.Cache {
	return &Cache{key: DefaultKey}
}

// redis底层调用各种操作的方法，不对外暴露 *!! args[0]必须为key
func (rds *Cache) do(method string, args ...interface{}) (ret interface{}, err error) {
	if len(args) < 1 {
		return nil, errors.New("cache:redis missing required arguments")
	}
	args[0] = rds.join(args[0])
	conn := rds.po.Get()
	defer conn.Close()

	return conn.Do(method, args...)
}

// 将调用的键名拼接前缀名 比如cacheRedis:go_lib
func (rds *Cache) join(originKey interface{}) string {
	return fmt.Sprintf("%s:%s", rds.key, originKey)
}

// 通过key获取value
func (rds *Cache) Get(key string) interface{} {
	if v, err := rds.do("GET", key); err == nil {
		return v
	}
	return nil
}


//spush
func (rds *Cache) Lpush(key string, val string) error {
	_, err := rds.do("LPUSH", key, val)
	return err
}

func (rds *Cache) LRange(key string, start, end int) ([]string, error) {
	strs := make([]string, 0)
	inters, err := redis.Values(rds.do("LRange", key, start, end))
	if err != nil {
		return strs, err
	}
	err = redis.ScanSlice(inters, &strs)
	return strs, err
}

func (rds *Cache) LLen(key string) (interface{}, error) {
	return rds.do("LLEN", key)
}

func (rds *Cache) Lpop(key string) (interface{}, error) {
	return rds.do("LPOP", key)
}

func (rds *Cache) Rpop(key string) (interface{}, error) {
	return rds.do("RPOP", key)
}

func (rds *Cache) Sadd(key string, arg ...interface{}) (interface{}, error) {
	return rds.do("SADD", key, arg)
}

func (rds *Cache) SaddInt64(key string, arg int64) (interface{}, error) {
	return rds.do("SADD", key, arg)
}

func (rds *Cache) SisMembersInt64(key string, arg int64) (int64, error) {
	is, err := rds.do("SISMEMBER", key, arg)
	if err != nil {
		apilog.DebugLog("sismember失败", err)
	}
	return strconv.ParseInt(fmt.Sprintf("%d", is), 10, 64)
}

func (rds *Cache) SaddString(key string, arg string) (interface{}, error) {
	return rds.do("SADD", key, arg)
}

func (rds *Cache) Srem(key string, arg interface{}) (interface{}, error) {

	return rds.do("SREM", key, arg)
}

func (rds *Cache) SmemberInt64(key string) ([]int64, error) {
	Ids := make([]int64, 0)
	inter, err := redis.Values(rds.do("SMEMBERS", key))
	if err != nil {
		apilog.DebugLog("redis获取集合失败(smembers)", err)
	}
	redis.ScanSlice(inter, &Ids)
	return Ids, err
}

func (rds *Cache) SmemberString(key string) ([]string, error) {
	strs := make([]string, 0)
	inter, err := redis.Values(rds.do("SMEMBERS", key))
	if err != nil {
		apilog.DebugLog("redis获取集合失败(smembers)", err)
		return strs, err
	}
	redis.ScanSlice(inter, &strs)
	return strs, err
}

func (rds *Cache) HKeys(key string) ([]string, error) {
	strs := make([]string, 0)
	inter, err := redis.Values(rds.do("HKEYS", key))
	if err != nil {
		apilog.DebugLog("redis获取集合失败(smembers)", err)
		return strs, err
	}
	redis.ScanSlice(inter, &strs)
	return strs, err
}

func (rds *Cache) Hset(key, ukey, value string) (interface{}, error) {
	return rds.do("HSET", key, ukey, value)
}

func (rds *Cache) HGet(key, ukey string) (interface{}, error) {
	inter, err := rds.do("HGET", key, ukey)
	if err != nil {
		apilog.DebugLog("redis获hash合失败(hash)", err)
		return inter, err
	}
	return inter, err
}

func (rds *Cache) HGetAll(key string) (interface{}, error) {
	var val []string
	inter, err := redis.Values(rds.do("HGETALL", key))
	if err != nil {
		apilog.DebugLog("redis获hash合失败(hgetall)", err)
		return val, err
	}
	err = redis.ScanSlice(inter, &val)
	return val, err
}

func (rds *Cache) HDel(key, ukey string) (interface{}, error) {
	inter, err := rds.do("HDEL", key, ukey)
	if err != nil {
		apilog.DebugLog("redis删除hash合失败(hash)", err, ukey, key)
		return inter, err
	}
	return inter, err
}

func (rds *Cache) HLen(key string) (interface{}, error) {
	inter, err := rds.do("HLEN", key)
	if err != nil {
		apilog.DebugLog("redis获取hlen合失败(hash)", err, key)
		return inter, err
	}
	return inter, err
}

func (rds *Cache) HExists(key, ukey string) (interface{}, error) {
	inter, err := rds.do("HExists", key, ukey)
	if err != nil {
		apilog.DebugLog("redis获取hexists合失败(hash)", err, key)
		return inter, err
	}
	return inter, err
}

// 通过keys的切片去获取对应的值
func (rds *Cache) GetMulti(keys []string) []interface{} {
	var args []interface{}
	c := rds.po.Get()
	defer c.Close()
	// 批量添加前缀
	for _, key := range keys {
		args = append(args, rds.join(key))
	}
	// 获取结果
	if values, err := redis.Values(c.Do("MGET", args...)); err == nil {
		return values
	}
	return nil
}

// 将数据根据key推入redis
func (rds *Cache) Put(key string, val interface{}, timeout time.Duration) error {
	_, err := rds.do("SETEX", key, int64(timeout/time.Second), val)
	return err
}

// 自增，如果不传增长的值，默认自增1
func (rds *Cache) Inc(key string, num ...int) error {
	var inc int
	if num == nil {
		inc = 1
	} else {
		inc = num[0]
	}
	_, err := redis.Bool(rds.do("INCRBY", key, inc))
	return err
}

// 自减， 如果不传减小的值，默认自减1
func (rds *Cache) Dec(key string, num ...int) error {
	var inc int
	if num == nil {
		inc = -1
	} else {
		inc = -1 * num[0]
	}
	_, err := redis.Bool(rds.do("INCRBY", key, inc))
	return err
}

// 删除某个key
func (rds *Cache) Del(key string) error {
	_, err := rds.do("DEL", key)
	return err
}

// 删除某个key通配删除
func (rds *Cache) Dels(key string) error {
	c := rds.po.Get()
	defer c.Close()
	// 现将所有的key取出
	cachedKeys, err := redis.Strings(c.Do("KEYS", rds.key+":"+key+"*"))
	if err != nil {
		return err
	}
	// 遍历删除
	for _, str := range cachedKeys {
		if _, err = c.Do("DEL", str); err != nil {
			return err
		}
	}
	return err
}

// 判断是否存在
func (rds *Cache) IsExist(key string) bool {
	v, err := redis.Bool(rds.do("EXISTS", key))
	if err != nil {
		return false
	}
	return v
}

// 清空整个数据仓库
func (rds *Cache) DelAll() error {
	// 原本的设计想法是用flushdb命令去执行，后来想了想，不能这么做
	// 原因是这么做的话会将非本系统的其他值删除掉(共用一个redis数据仓库)
	// 因为会循环执行，如果
	c := rds.po.Get()
	defer c.Close()
	// 现将所有的key取出
	cachedKeys, err := redis.Strings(c.Do("KEYS", rds.key+":*"))
	if err != nil {
		return err
	}
	// 遍历删除
	for _, str := range cachedKeys {
		if _, err = c.Do("DEL", str); err != nil {
			return err
		}
	}
	return err
}

// 注入驱动配置，生成连接
func (rds *Cache) Connect(config string) error {
	var cf map[string]string
	_ = json.Unmarshal([]byte(config), &cf)

	if _, ok := cf["key"]; !ok {
		cf["key"] = DefaultKey
	}
	if _, ok := cf["conn"]; !ok {
		return errors.New("cache: redis config has no conn key")
	}

	// redis连接配置格式 redis://<password>@<host>:<port>
	cf["conn"] = strings.Replace(cf["conn"], "redis://", "", 1)
	if i := strings.Index(cf["conn"], "@"); i > -1 {
		cf["password"] = cf["conn"][0:i]
		cf["conn"] = cf["conn"][i+1:]
	}

	if _, ok := cf["dbNum"]; !ok {
		cf["dbNum"] = "0"
	}
	if _, ok := cf["password"]; !ok {
		cf["password"] = ""
	}
	if _, ok := cf["maxIdle"]; !ok {
		cf["maxIdle"] = "3"
	}
	rds.key = cf["key"]
	rds.connInfo = cf["conn"]
	rds.dbNum, _ = strconv.Atoi(cf["dbNum"])
	rds.password = cf["password"]
	rds.maxIdle, _ = strconv.Atoi(cf["maxIdle"])

	rds.connectInit()

	c := rds.po.Get()
	defer c.Close()

	return c.Err()
}

// 连接到redis
func (rds *Cache) connectInit() {
	dialFunc := func() (c redis.Conn, err error) {
		c, err = redis.Dial("tcp", rds.connInfo)
		if err != nil {
			return nil, err
		}

		if rds.password != "" {
			if _, err := c.Do("AUTH", rds.password); err != nil {
				c.Close()
				return nil, err
			}
		}

		_, selectErr := c.Do("SELECT", rds.dbNum)
		if selectErr != nil {
			c.Close()
			return nil, selectErr
		}
		return
	}
	// 初始化一个新的pool
	rds.po = &redis.Pool{
		MaxIdle:     rds.maxIdle,
		IdleTimeout: 180 * time.Second,
		Dial:        dialFunc,
	}
}

// 初始化，将redis的Cache结构的指针返回给cache的注册中心
func init() {
	cache.Register("redis", NewCacheRedis)
}
