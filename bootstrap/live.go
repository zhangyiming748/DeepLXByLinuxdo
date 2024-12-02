package bootstrap

import (
	"DeepLXByLinuxdo/controller"
	"github.com/gin-gonic/gin"
)

func InitHello(engine *gin.Engine) {
	routeGroup := engine.Group("/api/v1")
	{
		c := new(controller.LiveController)
		routeGroup.GET("/hello", c.GetHello)

	}
}
