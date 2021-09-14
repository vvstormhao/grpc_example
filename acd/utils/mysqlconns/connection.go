package mysqlconns

import (
	"fmt"
	"io"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var conn *Connection
var once sync.Once

// MySQLConfig 数据库配置
type MySQLConfig struct {
	Endpoint     string
	User         string
	Passwd       string
	DBName       string
	MaxIdleConns int
	MaxOpenConns int
	OutPut       io.Writer
}

// NewConfig 创建mysql 配置
func NewConfig(endpoint, user, passwd, dbName string) *MySQLConfig {
	return &MySQLConfig{
		Endpoint:     endpoint,
		User:         user,
		Passwd:       passwd,
		DBName:       dbName,
		MaxIdleConns: 10,
		MaxOpenConns: 30,
	}
}

// Connection MySQL连接
type Connection struct {
	gormDB   *gorm.DB
	endpoint string
}

// Instance 获取Mysql连接
func Instance() *Connection {
	once.Do(func() {
		conn = &Connection{}
	})
	return conn
}

// Init 初始化数据库连接
func (c *Connection) Init(cfg *MySQLConfig) error {
	var err error
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local&readTimeout=1s&writeTimeout=1s",
		cfg.User,
		cfg.Passwd,
		cfg.Endpoint,
		cfg.DBName,
	)

	fmt.Println(dsn)
	c.gormDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if nil == err {
		if db, err := c.gormDB.DB(); nil == err {
			db.SetMaxIdleConns(cfg.MaxIdleConns)
			db.SetMaxOpenConns(cfg.MaxOpenConns)
			db.Ping()
			return nil
		} else {
			return fmt.Errorf("Init failed, %v", err)
		}
	}
	c.endpoint = cfg.Endpoint

	return fmt.Errorf("Init failed %v", err)
}

// DB 返回gormDB
func (c *Connection) DB() *gorm.DB {
	return c.gormDB
}

func (c *Connection) Endpoint() string {
	return c.endpoint
}
