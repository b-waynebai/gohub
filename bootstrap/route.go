// package bootstrap 处理程序初始化逻辑
package bootstrap

import (
	"gohub/app/http/middlewares"
	"gohub/routes"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// SetupRoute 路由初始化
func SetupRoute(router *gin.Engine) {

	// 注册全局中间件
	registerGlobalMiddleWare(router)

	// 注册 Api 路由
	routes.RegisterApiRoutes(router)

	// 配置404 请求
	setup404Handler(router)
}

// 全局中间件
func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(
		// gin.Logger(),
		middlewares.Logger(),
		gin.Recovery(),
	)
}

// 处理404 请求
func setup404Handler(router *gin.Engine) {

	router.NoRoute(func(c *gin.Context) {
		// 获取标头信息的 Accept 信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			c.String(http.StatusNotFound, "页面返回 404")
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "页面不存在，请确认请求路径",
			})
		}
	})
}
