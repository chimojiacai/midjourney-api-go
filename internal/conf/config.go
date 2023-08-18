// @Author: liyongzhen
// @Description:
// @File: config
// @Date: 2023/8/18 15:11

package conf

import (
	"github.com/spf13/viper"
)

var Conf = new(Config)

func NewConfig() {
	config, err := ParseConfig()
	if err != nil {
		panic("解析配置文件失败：" + err.Error())
	}
	Conf = config
}

// ParseConfig 解析配置文件
func ParseConfig() (*Config, error) {
	c := new(Config)
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig() // 会查找和读取配置文件
	if err != nil {             // Handle errors reading the config file
		return c, err
	}
	viper.WatchConfig()
	err = viper.Unmarshal(&c)
	if err != nil {
		return c, err
	}
	if c.HttpPort == 0 {
		c.HttpPort = 1234
	}
	return c, err
}

type Config struct {
	HttpPort       int    `json:"http_port"`
	Channelid      string `json:"channelid"`
	Authorization  string `json:"authorization"`
	Application_Id string `json:"application_id"`
	Guild_Id       string `json:"guild_id"`
	Session_Id     string `json:"session_id"`
	Version        string `json:"version"`
	Id             string `json:"id"`
	Flags          string `json:"flags"`
}
