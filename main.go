package main

import (
	"fmt"
	"gohub/bootstrap"

	"github.com/gin-gonic/gin"
)

func main() {

	// new一个 gin engine实例
	router := gin.New()

	// 初始化路由绑定
	bootstrap.SetupRoute(router)

	// 运行服务
	err := router.Run(":3000")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

}
