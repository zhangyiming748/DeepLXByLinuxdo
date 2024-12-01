package storage

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"time"
	"xorm.io/xorm"
)

var engine *xorm.Engine

func SetMysql() {
	var err error
	hostname := "127.0.0.1"
	if os.Getenv("MYSQL_HOSTNAME") != "" {
		hostname = os.Getenv("MYSQL_HOSTNAME")
	}
	port := "3306"
	if os.Getenv("MYSQL_PORT") != "" {
		port = os.Getenv("MYSQL_PORT")
	}
	passwd := "123456"
	if os.Getenv("MYSQL_PASSWD") != "" {
		passwd = os.Getenv("MYSQL_PASSWD")
	}
	dataSourceName := fmt.Sprintf("root:%s@tcp(%s:%s)/Translate?charset=utf8", passwd, hostname, port)
	engine, err = xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	engine.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
}
func GetMysql() *xorm.Engine {
	return engine
}
