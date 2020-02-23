package cache

import (
	"fmt"
	"time"
)

// cache是实现驱动的一个规范，并且能够注册和实例化驱动
type Cache interface {
	// 通过key获取值
	Get(key string) interface{}

	// 通过key的切片获取对应的值
	GetMulti(keys []string) []interface{}

	// 通过key-value的形式创建缓存
	Put(key string, val interface{}, timeout time.Duration) error

	// 通过key自增n
	Inc(key string, num ...int) error

	// 通过key自减n
	Dec(key string, num ...int) error

	// 通过key删除缓存
	Del(key string) error

	// 通过key前缀批量删除
	Dels(key string) error

	// 查看是否存在
	IsExist(key string) bool

	// 清空缓存
	DelAll() error

	// 创建链接
	Connect(config string) error
}

// 创建一个缓存实例
type Instance func() Cache

// 构建实例map，为各种类型的驱动
var drivers = make(map[string]Instance)

// 具体驱动init的时候执行的注册方法
func Register(name string, driver Instance) {
	if driver == nil {
		panic("cache: Register driver is nil, register error, driver name is " + name)
	}

	if _, ok := drivers[name]; ok {
		panic("cache: Register called repeatedly for driver " + name)
	}
	// 将驱动加入到注册map
	drivers[name] = driver
}

// 创建cache实例
func New(driverName string, config string) (driver Cache, err error) {
	// 查看实例化的驱动是否已经注册到实例map
	instanceFunc, ok := drivers[driverName]
	if !ok {
		err = fmt.Errorf("cache: undefind driver name %q (forgot to import or not achieved it?)", driverName)
		return
	}
	// 实例化驱动 administrator
	driver = instanceFunc()
	err = driver.Connect(config)
	if err != nil {
		driver = nil
	}
	return
}
