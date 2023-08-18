// @Author: liyongzhen
// @Description:
// @File: task
// @Date: 2023/8/18 17:12

package task_model

type TaskInfo struct {
	ID          string `json:"id"`          // 任务id
	Action      string `json:"action"`      // 动作
	Prompt      string `json:"prompt"`      // 原生咒语
	PromptEn    string `json:"prompt_en"`   // 英文咒语
	Description string `json:"description"` // 任务描述
	SubmitTime  string `json:"submit_time"` // 任务提交时间
	StartTime   string `json:"start_time"`  // 任务开始时间
	FinishTime  string `json:"finish_time"` // 任务结束时间
	ImgUrl      string `json:"img_url"`     // 图片链接
	Status      string `json:"status"`      // 任务状态
	Progress    string `json:"progress"`    // 任务进度
	FailReason  string `json:"fail_reason"` // 失败原因
	State       string `json:"state"`       // 自定义参数
	NotifyHook  string `json:"notify_hook"` // 回调地址
}
