package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/haimait/gin_web/conf"
	"github.com/haimait/gin_web/lib/cache"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

var Db *gorm.DB

type Model struct {
	BaseModel
	ID         int64 `gorm:"primary_key" json:"id"`
	Status     int   `json:"status"`
	CreateTime int64 `json:"create_time"`
	UpdateTime int64 `json:"update_time"`
}

type BaseModel struct{}

var cc cache.Cache

func init() {
	var (
		err                                               error
		DRIVER_TYPE, dbName, user, password, host, tablePrefix ,charset string
		maxIdleConns, maxOpenConns                        int
	)

	dbConf := conf.Cfg.DatabaseConf
	DRIVER_TYPE = dbConf.DRIVER_TYPE
	dbName = dbConf.DBNAME
	user = dbConf.USER
	password = dbConf.PASSWORD
	host = dbConf.HOST
	charset = dbConf.CHARSET
	mysqlConf := conf.Cfg.MysqlConf
	tablePrefix = mysqlConf.TABLE_PREFIX
	maxIdleConns =  mysqlConf.MAX_IDLE_CONNS
	maxOpenConns = mysqlConf.MAX_OPEN_CONNS



	Db, err = gorm.Open(DRIVER_TYPE, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName,
		charset))

	if err != nil {
		log.Println(err)
	}
	//关闭连接
	//defer Db.Close()
	// 设置表前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	// 设置数据表关闭复数结尾
	Db.SingularTable(true)
	Db.DB().SetMaxIdleConns(maxIdleConns)
	Db.DB().SetMaxOpenConns(maxOpenConns)
	// 自动迁移
	db := Db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.AutoMigrate(&User{})

	Db.LogMode(true)




	// 加载缓存配置
	redisCfg := conf.Cfg.RedisConf.HOST
	// 初始化缓存
	cc, _ = cache.New("redis", `{"conn":"`+redisCfg+`"}`)
}

// @desc 创建一条用户数据前的钩子
func (Model *Model) BeforeCreate(scope *gorm.Scope) error {
	_ = scope.SetColumn("created_at", time.Now().Unix())
	_ = scope.SetColumn("updated_at", time.Now().Unix())
	return nil
}

// @desc 更新一条用户数据前的钩子
func (Model *Model) BeforeUpdate(scope *gorm.Scope) error {
	_ = scope.SetColumn("updated_at", time.Now().Unix())
	return nil
}


// 删单缓存
func (BaseModel *BaseModel) forgetPrimaryCache(pk string) {
	_ = cc.Del(pk)
}

// 删除块状缓存
func (BaseModel *BaseModel) forgetBlockCache(bks []string) {
	for _, ck := range bks {
		// 获取key
		_ = cc.Dels(ck)
	}
}

func (BaseModel *BaseModel) Begin() {
	Db.Begin()
}

func (BaseModel *BaseModel) RollBack() {
	Db.Rollback()
}

func (BaseModel *BaseModel) Commit() {
	Db.Commit()
}

func (BaseModel *BaseModel) DB() *gorm.DB {
	return Db
}
