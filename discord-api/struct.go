// @Author: cmjc
// @Description:
// @File: params
// @Date: 2023/5/11 11:36

package discord_api

import "time"

// SenderParams sender的请求参数
type SenderParams struct {
	Type          int    `json:"type"`
	ApplicationID string `json:"application_id"`
	GuildID       string `json:"guild_id"`
	ChannelID     string `json:"channel_id"`
	SessionID     string `json:"session_id"`
	Data          Data   `json:"data"`
	Nonce         string `json:"nonce"`
}

type Options struct {
	Type  int    `json:"type"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Options0 struct {
	Type        int    `json:"type"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Required    bool   `json:"required"`
}

type ApplicationCommand struct {
	ID                       string     `json:"id"`
	ApplicationID            string     `json:"application_id"`
	Version                  string     `json:"version"`
	DefaultMemberPermissions any        `json:"default_member_permissions"`
	Type                     int        `json:"type"`
	Nsfw                     bool       `json:"nsfw"`
	Name                     string     `json:"name"`
	Description              string     `json:"description"`
	DmPermission             bool       `json:"dm_permission"`
	Contexts                 any        `json:"contexts"`
	Options                  []Options0 `json:"options"`
}

type Data struct {
	Version            string             `json:"version"`
	ID                 string             `json:"id"`
	Name               string             `json:"name"`
	Type               int                `json:"type"`
	Options            []Options          `json:"options"`
	ApplicationCommand ApplicationCommand `json:"application_command"`
	Attachments        []any              `json:"attachments"`
}

// MsgData msg的结构体
type MsgData struct {
	ID               string           `json:"id"`
	Type             int              `json:"type"`
	Content          string           `json:"content"`
	ChannelID        string           `json:"channel_id"`
	Author           Author           `json:"author"`
	Attachments      []*Attachments   `json:"attachments"`
	Embeds           []Embeds         `json:"embeds"`
	Mentions         []any            `json:"mentions"`
	MentionRoles     []any            `json:"mention_roles"`
	Pinned           bool             `json:"pinned"`
	MentionEveryone  bool             `json:"mention_everyone"`
	Tts              bool             `json:"tts"`
	Timestamp        time.Time        `json:"timestamp"`
	EditedTimestamp  *time.Time       `json:"edited_timestamp"`
	Flags            int              `json:"flags"`
	Components       []any            `json:"components"`
	WebhookID        string           `json:"webhook_id,omitempty"`
	MessageReference MessageReference `json:"message_reference,omitempty"`
	ApplicationID    string           `json:"application_id,omitempty"`
	Interaction      Interaction      `json:"interaction,omitempty"`
}

type Attachments struct {
	ContentType string `json:"content_type"`
	Filename    string `json:"filename"`
	Height      int    `json:"height"`
	ID          string `json:"id"`
	ProxyURL    string `json:"proxy_url"`
	Size        int    `json:"size"`
	URL         string `json:"url"`
	Width       int    `json:"width"`
	Prompt      string `json:"prompt"`
}

type Author struct {
	ID               string `json:"id"`
	Username         string `json:"username"`
	GlobalName       string `json:"global_name"`
	DisplayName      string `json:"display_name"`
	Avatar           string `json:"avatar"`
	Discriminator    string `json:"discriminator"`
	PublicFlags      int    `json:"public_flags"`
	Bot              bool   `json:"bot"`
	AvatarDecoration any    `json:"avatar_decoration"`
}

type Thumbnail struct {
	URL      string `json:"url"`
	ProxyURL string `json:"proxy_url"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
}

type Embeds struct {
	Type      string    `json:"type"`
	URL       string    `json:"url"`
	Thumbnail Thumbnail `json:"thumbnail"`
}

type MessageReference struct {
	ChannelID string `json:"channel_id"`
	GuildID   string `json:"guild_id"`
	MessageID string `json:"message_id"`
}

type User struct {
	ID               string `json:"id"`
	Username         string `json:"username"`
	GlobalName       string `json:"global_name"`
	DisplayName      string `json:"display_name"`
	Avatar           string `json:"avatar"`
	Discriminator    string `json:"discriminator"`
	PublicFlags      int    `json:"public_flags"`
	AvatarDecoration any    `json:"avatar_decoration"`
}

type Interaction struct {
	ID   string `json:"id"`
	Type int    `json:"type"`
	Name string `json:"name"`
	User User   `json:"user"`
}

//// MjMsgData 筛选出来mj返回的数据
//type MjMsgData struct {
//	MsgId         string `json:"msg_id"`         // 消息id
//	AttachmentsId string `json:"attachments_id"` // 附件消息id
//	FileName      string `json:"file_name"`      // 文件名字
//	Height        int    `json:"height"`         // 图片高
//	Width         int    `json:"width"`          // 图片宽
//	Url           string `json:"url"`            // 图片访问url
//	ProxyURL      string `json:"proxy_url"`      // 图片访问代理url
//}
