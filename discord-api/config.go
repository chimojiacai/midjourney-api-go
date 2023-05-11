// @Author: cmjc
// @Description:
// @File: config
// @Date: 2023/5/11 10:56

package discord_api

import (
	"github.com/spf13/viper"
)

type Config struct {
	Channelid      string `json:"channelid"`
	Authorization  string `json:"authorization"`
	Application_Id string `json:"application_id"`
	Guild_Id       string `json:"guild_id"`
	Session_Id     string `json:"session_id"`
	Version        string `json:"version"`
	Id             string `json:"id"`
	Flags          string `json:"flags"`
}

// ParseConfig 解析配置文件
func ParseConfig() (Config, error) {
	var c Config
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig() // 会查找和读取配置文件
	if err != nil {             // Handle errors reading the config file
		return c, err
	}
	err = viper.Unmarshal(&c)
	if err != nil {
		return c, err
	}
	return c, err
}
