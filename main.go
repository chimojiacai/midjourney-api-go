// @Author: cmjc
// @Description: 跑起来
// @File: main
// @Date: 2023/5/11 11:00

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"midjourney_api/internal/conf"
	"midjourney_api/internal/http/router"
)

func main() {
	// 加载配置文件
	conf.NewConfig() // 加载配置文件，找不到会panic

	r := gin.New()
	router.Router(r) // 初始化路由

	s := conf.InitServer(fmt.Sprintf(":%d", conf.Conf.HttpPort), r) // 启动服务+优雅关闭
	log.Println(fmt.Sprintf("server run success：0.0.0.0:%d", +conf.Conf.HttpPort))
	log.Fatalf(s.ListenAndServe().Error())
}
