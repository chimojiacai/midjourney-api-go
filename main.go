// @Author: cmjc
// @Description: 跑起来
// @File: main
// @Date: 2023/5/11 11:00

package main

import (
	"encoding/json"
	"fmt"
	"log"
	discordapi "midjourney_api/discord-api"
)

var c discordapi.Config

func init() {
	config, err := discordapi.ParseConfig()
	if err != nil {
		log.Fatal("解析配置文件失败：", err)
	}
	c = config
}

func main() {
	// 请求生成图
	err := c.SenderReq("cat")
	if err != nil {
		log.Fatal(err)
	}

	// 异步拉取聊天记录，读取历史生成的图片
	data, err := c.MidJourneyMsg()
	if err != nil {
		log.Fatal(err)
	}
	marshal, _ := json.Marshal(data)
	fmt.Println(string(marshal))
}
