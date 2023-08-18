// @Author: cmjc
// @Description: 接收消息数据
// @File: api
// @Date: 2023/5/11 10:29

package discord_api

import (
	"fmt"
	"github.com/imroc/req"
	"midjourney_api/common"
	"net/http"
	"regexp"
	"strings"
	"time"
)

// MidJourneyMsg 筛选出来mj生成的图
func (c Config) MidJourneyMsg() ([]*Attachments, error) {
	msgList, err := c.ReceiverReq()
	if err != nil {
		return nil, err
	}
	if len(msgList) == 0 {
		fmt.Println("读取聊天记录为空！")
		return nil, nil
	}
	var resList []*Attachments
	// 开始筛选MJ数据
	for _, data := range msgList {
		// 不是mj返回的消息过滤掉
		if data.Author.Username != common.MjUsername || !strings.Contains(data.Content, "**") {
			continue
		}
		// 拿到prompt
		// 使用正则表达式提取字符串
		re := regexp.MustCompile(`\*\*(.*?)\*\*`)
		submatch := re.FindStringSubmatch(data.Content)

		if len(data.Attachments) > 0 {
			// 每次只返回一个，直接取下标0即可
			data.Attachments[0].Prompt = submatch[1]
			resList = append(resList, data.Attachments[0])
		} else {
			fmt.Println("test no attachments")
		}
	}
	return resList, nil
}

// ReceiverReq 读取频道里的所有聊天记录
func (c Config) ReceiverReq() ([]*MsgData, error) {
	h := req.Header{
		"authorization": c.Authorization,
	}
	req.SetTimeout(3 * time.Second) // 设置超时时间
	response, err := req.Get(fmt.Sprintf(common.ChannelMsgUrl, c.Channelid), h)
	if err != nil {
		return nil, err
	}
	if response.Response().StatusCode != http.StatusOK {
		return nil, fmt.Errorf("请求API失败；返回结果：%s", response.String())
	}
	var data []*MsgData
	err = response.ToJSON(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
