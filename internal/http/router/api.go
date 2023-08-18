// @Author: liyongzhen
// @Description:
// @File: api
// @Date: 2023/8/18 15:01

package router

import (
	"github.com/gin-gonic/gin"
	"midjourney_api/internal/http/handler"
	"midjourney_api/internal/http/middleware"
	"midjourney_api/pkg/response"
)

// Router 路由
func Router(r *gin.Engine) {
	r.Use(middleware.Recover)        // 全局异常
	r.NoRoute(middleware.Recover404) // 找不到路由
	r.GET("/ping", func(c *gin.Context) {
		response.WithCtx(c).Success("pong")
	}) // 测试服务是否启动

	// api路由
	apiRouter(r)
}

func apiRouter(r *gin.Engine) {
	mj := r.Group("mj")
	{
		mj.POST("imagine", handler.Imagine)
	}

	// task
}
