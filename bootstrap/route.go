// Package bootstrap 处理程序初始化逻辑
package bootstrap

import (
	"gohub/routes"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

//SetupRoute 路由初始化

func SetupRoute(router *gin.Engine) {

	// 注册全局中间件
	registerGlobleMiddleWare(router)

	// 注册 API 路由
	routes.RegisterAPIRouters(router)
	// 配置 404 路由
	setup404Handler(router)
}

func registerGlobleMiddleWare(router *gin.Engine) {
	router.Use(
		gin.Logger(),
		gin.Recovery(),
	)
}

func setup404Handler(router *gin.Engine) {
	//处理 404 请求
	router.NoRoute(func(ctx *gin.Context) {

		// 获取标头信息的 Accept 信息
		acceptString := ctx.Request.Header.Get("Accept")
		// 判断是不是HTML
		// 如果是HTML,返回页面
		if strings.Contains(acceptString, "text/html") {
			ctx.JSON(http.StatusOK, "页面返回404")
		} else {
			// 如果不是HTML则返回默认JSON
			ctx.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})
}
