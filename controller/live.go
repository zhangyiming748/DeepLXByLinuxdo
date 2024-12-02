package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type LiveController struct{}

// 定义一个结构体，用于作为返回的JSON数据结构
type User struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdmin bool   `json:"is_admin"`
}

func (l LiveController) GetHello(c *gin.Context) {
	user := User{
		Name:    "John Doe",
		Age:     30,
		IsAdmin: false,
	}

	// 将User结构体转换为JSON格式并返回给客户端，状态码设置为200
	c.JSON(http.StatusOK, user)
}
