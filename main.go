package main

import (
	"DeepLXByLinuxdo/bootstrap"
	"DeepLXByLinuxdo/model"
	"DeepLXByLinuxdo/storage"
	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func testResponse(c *gin.Context) {
	c.JSON(http.StatusGatewayTimeout, gin.H{
		"code": http.StatusGatewayTimeout,
		"msg":  "timeout",
	})
}

func timeoutMiddleware() gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(3000*time.Millisecond),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(testResponse),
	)
}
func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	storage.SetMysql()
	if err := storage.GetMysql().Ping(); err != nil {
		log.Fatalf("mysql ping err:%v", err)
	}
	if err := storage.GetMysql().Sync2(new(model.TranslateCache)); err != nil {
		log.Fatalf("sync translate cache err:%v", err)
	}
}
func main() {
	// gin服务
	gin.SetMode(gin.DebugMode)
	engine := gin.New()
	engine.Use(timeoutMiddleware())
	bootstrap.InitTranslate(engine)
	// 启动http服务
	engine.Run(":2147")
}
