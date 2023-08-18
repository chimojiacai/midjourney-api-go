// @Author: cmjc
// @Description: 生成图片请求
// @File: api
// @Date: 2023/5/11 10:29

package discord_api

import (
	"fmt"
	"github.com/imroc/req"
	"midjourney_api/internal/common"
	"midjourney_api/internal/conf"
	"midjourney_api/utils"
	"net/http"
	"time"
)

// SenderReq 开始画图
func SenderReq(prompt string) error {
	header := req.Header{
		"authorization": conf.Conf.Authorization,
	}

	p, err := senderParams(prompt)
	if err != nil {
		return err
	}

	req.SetTimeout(3 * time.Second) // 设置超时时间
	response, err := req.Post(common.CmdSendUrl, header, req.BodyJSON(p))
	if err != nil {
		return err
	}
	if response.Response().StatusCode != http.StatusNoContent {
		return fmt.Errorf("请求API失败；返回结果：%s", response.String())
	}
	return nil
}

// 拼接请求参数
func senderParams(prompt string) (*SenderParams, error) {
	c := conf.Conf
	// 生成nonce值
	node, err := utils.NewNode(1)
	if err != nil {
		return nil, err
	}
	return &SenderParams{
		Type:          2,
		ApplicationID: c.Application_Id,
		GuildID:       c.Guild_Id,
		ChannelID:     c.Channelid,
		SessionID:     c.Session_Id,
		Data: Data{
			Version: c.Version,
			ID:      c.Id,
			Name:    common.CmdName,
			Type:    1,
			Options: []Options{{Type: 3, Name: "prompt", Value: prompt}},
			ApplicationCommand: ApplicationCommand{
				ID:                       c.Id,
				ApplicationID:            c.Application_Id,
				Version:                  c.Version,
				DefaultMemberPermissions: nil,
				Type:                     1,
				Nsfw:                     false,
				Name:                     common.CmdName,
				Description:              "Create images with Midjourney",
				DmPermission:             true,
				Contexts:                 nil,
				Options: []Options0{{
					Type:        3,
					Name:        "prompt",
					Description: "The prompt to imagine",
					Required:    true,
				}},
			},
			Attachments: nil,
		},
		Nonce: node.Generate().String(),
	}, nil
}
