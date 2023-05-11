# midjourney-api-go
由于midjourney官方未提供API，这里使用discord的API，模拟请求生成midjourney的图；并通过api请求获取聊天记录筛选出MJ返回的数据。

### 准备工作
本次操作需要在浏览器登陆discord。并且需要创建一个服务器，把Midjourney的bot邀请到服务器里；至此需要打开开发者工具(按F12也可)，找到Network（网络）选项卡，你将看到所有discord的api请求；
### 配置文件
接以上，回头服务器的聊天框，聊天框里输入 “/” 随便选择一个指令，并发送到聊天框里；此时我们可以在Network里看到“interaction”请求；点击它，找到以下字段(在Payload选项里):
+ channelid
+ application_id
+ guild_id
+ session_id
+ version
+ id
+ authorization(这个需要在Header选项卡里找到request headers里找到)
### 克隆这个仓库
```azure
git clone https://github.com/chimojiacai/midjourney-api-go.git
```
### 配置文件
打开config.json把上面的值赋值到字段上，切记不要更改字段名字

### 运行
#### go环境
    >=go1.19
#### 下载依赖
```azure
1. export GOPROXY=https://goproxy.io,direct
2. go mod tidy 
```

### 基础代码
找到main文件分别运行以下代码
```azure
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
```
#### 发送生成图片请求
```azure
err := c.SenderReq("cat") // cat 改成自己的咒语
if err != nil {
    log.Fatal(err)
}
```

#### 读取服务器里的历史消息
```azure
// 异步拉取聊天记录，读取历史生成的图片
data, err := c.MidJourneyMsg()
if err != nil {
    log.Fatal(err)
}
marshal, _ := json.Marshal(data)
fmt.Println(string(marshal))
```
### 注意
这里只是实现了一个简单的demo。由于请求生成图片的接口有速率限制，所有需要开发者自己实现队列取请求api；否则可能请求不成功。

### TODO 
+ 图生图api请求开发
+ 图生图后的数据抓取
+ 其他

### 感谢
George-iam
