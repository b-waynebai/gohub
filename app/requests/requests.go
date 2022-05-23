// Package requests 处理请求数据和表单验证
package requests

import (
	"fmt"
	"gohub/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

// ValidatorFunc 验证函数
type ValidatorFunc func(interface{}, *gin.Context) map[string][]string

func Validate(c *gin.Context, obj interface{}, handler ValidatorFunc) bool {

	// 1 解析请求 支持 JSON 数据、表单请求和 URL Query
	if err := c.ShouldBind(obj); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。",
			"error":   err,
		})
		fmt.Println(err.Error())
		return false
	}

	// 2 表单验证
	errs := handler(obj, c)

	// 3 判断验证是否通过
	if len(errs) > 0 {
		// c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
		// 	"message": "请求验证不通过，具体查看 errors",
		// 	"errors":  errs,
		// })
		response.ValidationError(c, errs)
		return false
	}

	return true
}

func validate(date interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {
	opts := govalidator.Options{
		Data:          date,
		Rules:         rules,
		TagIdentifier: "valid",
		Messages:      messages,
	}

	return govalidator.New(opts).ValidateStruct()
}
