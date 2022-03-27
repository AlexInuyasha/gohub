package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {

	// 初始化Gin实例
	// r := gin.Default()
	r := gin.New()

	//注册中间件
	r.Use(gin.Logger(), gin.Recovery())

	// 注册一个路由
	r.GET("/", func(c *gin.Context) {

		// 以JSON格式响应
		c.JSON(http.StatusOK, gin.H{
			"Hello": "World!",
		})
	})

	//处理404请求
	r.NoRoute(func(c *gin.Context) {
		//获取标头信息的Aeccpt信息
		acceptString := c.Request.Header.Get("Accept")

		// 如果是HTML的话
		if strings.Contains(acceptString, "text/html") {
			c.JSON(http.StatusNotFound, "页面返回404")
		} else {

			// 默认返回JSON
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})

	// 运行服务
	r.Run(":8080")
}
