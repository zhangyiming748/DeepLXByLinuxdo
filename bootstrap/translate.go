package bootstrap

import (
	"DeepLXByLinuxdo/controller"
	"github.com/gin-gonic/gin"
)

func InitTranslate(engine *gin.Engine) {
	routeGroup := engine.Group("/api/v1")
	{
		c := new(controller.TranslateController)
		//routeGroup.GET("/v1/s1/gethello", c.GetHello)
		routeGroup.POST("/translate", c.TransWord)
	}
}
