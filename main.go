package main

import (
	"DeepLXByLinuxdo/bootstrap"
	"DeepLXByLinuxdo/storage"
	"log"
	"net/http"
	"runtime"
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
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
	if runtime.GOOS == "windows" {
		log.Fatalln("程序不可以在Windows上运行")
	}
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	storage.SetDatabase()
}
func main() {
	// gin服务
	gin.SetMode(gin.DebugMode)
	engine := gin.New()
	engine.Use(timeoutMiddleware())
	bootstrap.InitTranslate(engine)
	//bootstrap.InitHello(engine)
	// 启动http服务
	engine.Run(":8192")
}
