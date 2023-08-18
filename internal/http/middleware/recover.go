package middleware

import (
	"github.com/gin-gonic/gin"
	"midjourney_api/pkg/response"
	"net/http"
)

// Recover 拦截全局异常
func Recover(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			response.WithCtx(ctx).Error(response.EDefaultCode, "服务器异常, 请稍后重试")
			ctx.Abort()
			return
		}
	}()
	ctx.Next()
}

func Recover404(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, map[string]any{
		"code": http.StatusNotFound,
		"msg":  "404 NOT FOUND",
	})
	ctx.Next()
}

// todo 校验API请求来源
