package storage

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"strings"
	"time"
	"xorm.io/xorm"
)

var engine *xorm.Engine

func SetMysql() {
	var err error
	sqlService := os.Getenv("SQL_SERVICE")
	if sqlService == "" {
		log.Fatalf("SQL_SERVICE environment variable is not set")
	}
	source := strings.Join([]string{"root:@tcp(", sqlService, ")/Translate?charset=utf8"}, "")
	engine, err = xorm.NewEngine("mysql", source)
	if err != nil {
		panic(err)
	}
	engine.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
}
func GetMysql() *xorm.Engine {
	return engine
}
