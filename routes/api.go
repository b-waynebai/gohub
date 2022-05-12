package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterApiRoutes(r *gin.Engine) {

	// 测试一个v1的路由
	v1 := r.Group("/v1")
	{
		// 注册一个路由
		v1.GET("/", func(ctx *gin.Context) {
			// 以JSON格式响应
			ctx.JSON(http.StatusOK, gin.H{
				"Hello": "World!",
			})
		})
	}
}
