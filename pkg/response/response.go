package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	ctx     *gin.Context
	Code    int         `json:"code"`
	TraceId string      `json:"trace_id"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

const (
	TraceId             = "trace_id"
	InsufficientBalance = 3002
	EDefaultCode        = -1
	EUnauthorizedCode   = 401
	SUCCESS             = 0
	ParamErr            = -2  // 参数错误
	TooManyRequest      = 429 // 请求太频繁
	Waiting             = 430 // 继续等待
)

func (r *Response) Success(d interface{}) {
	// 开始时间
	r.Msg = "success"
	if d == nil {
		d = struct{}{}
	}
	r.Data = d
	r.Code = SUCCESS
	r.ctx.JSON(http.StatusOK, r)
}

func WithCtx(ctx *gin.Context) *Response {
	r := &Response{}
	r.ctx = ctx
	r.TraceId = ctx.GetString(TraceId)
	return r
}

func (r *Response) Error(code int, msg ...string) {
	r.Code = code
	if len(msg) != 0 {
		r.Msg = msg[0]
	} else {
		r.Msg = message(code)
	}
	r.Data = struct {
	}{}
	r.ctx.JSON(http.StatusOK, r)
}
