// @Author: liyongzhen
// @Description:
// @File: req
// @Date: 2023/8/18 16:52

package req_params

type SubmitImagineParams struct {
	Prompt     string   `json:"prompt"`      // 咒语 required=true
	Urls       []string `json:"urls"`        // 图片链接，用户已有链接，无链接的话传base64
	Base64s    []string `json:"base64_s"`    // 图片base64数据，前端传过来bas64数据
	State      string   `json:"state"`       // 自定义参数
	NotifyHook string   `json:"notify_hook"` // 回调地址, 为空时使用全局notifyHook，全局为空时则不回调,这里的优先级高
}
