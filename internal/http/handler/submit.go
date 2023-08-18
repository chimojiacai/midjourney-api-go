// @Author: liyongzhen
// @Description:
// @File: submit
// @Date: 2023/8/18 14:58

package handler

import (
	"github.com/gin-gonic/gin"
	"midjourney_api/internal/moldes/req_params"
	"midjourney_api/pkg/response"
)

// Imagine 文生图+图生图
func Imagine(ctx *gin.Context) {
	p := &req_params.SubmitImagineParams{}
	if err := ctx.ShouldBindJSON(p); err != nil {
		response.WithCtx(ctx).Error(response.ParamErr)
		return
	}
	if p.Prompt == "" {
		response.WithCtx(ctx).Error(response.ParamErr, "prompt不能为空")
		return
	}

}

// SimpleChange 绘图变化-simple
func SimpleChange(ctx *gin.Context) {

}

// Change 绘图变化
func Change(ctx *gin.Context) {

}

// Describe 绘图变化
func Describe(ctx *gin.Context) {

}

// Blend 提交Blend任务
func Blend(ctx *gin.Context) {

}
