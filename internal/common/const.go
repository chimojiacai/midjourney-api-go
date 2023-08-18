// @Author: cmjc
// @Description:
// @File: const
// @Date: 2023/5/11 11:20

package common

const (
	CmdSendUrl    = "https://discord.com/api/v9/interactions"                   // 触发指令的api url
	CmdName       = "imagine"                                                   // 指令名字
	ChannelMsgUrl = "https://discord.com/api/v9/channels/%s/messages?limit=100" // 拉取群里的聊天记录 url, 每次读取100条
	MjUsername    = "Midjourney Bot"
)

const (
	TaskActionIMAGINE   = "IMAGINE"   // 生成图片
	TaskActionUPSCALE   = "UPSCALE"   // 选中放大
	TaskActionVARIATION = "VARIATION" // 选中其中的一张图，生成四张相似的
	TaskActionREROLL    = "REROLL"    // 重新执行
	TaskActionDESCRIBE  = "DESCRIBE"  // 图转prompt
	TaskActionBLEND     = "BLEND"     // 多图混合
)
