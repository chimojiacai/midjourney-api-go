// @Author: liyongzhen
// @Description:
// @File: config
// @Date: 2023/5/18 16:29

package conf

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"time"
)

type Server interface {
	ListenAndServe() error
}

func InitServer(address string, router *gin.Engine) Server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 10 * time.Millisecond
	s.WriteTimeout = 100 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}
