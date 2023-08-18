// @Author: liyongzhen
// @Description:
// @File: submit
// @Date: 2023/8/18 17:02

package service

import (
	"golang.org/x/net/context"
	"midjourney_api/internal/common"
	discord_api "midjourney_api/internal/discord-api"
	"midjourney_api/internal/moldes/req_params"
	"midjourney_api/internal/moldes/task_model"
	"midjourney_api/internal/service/store"
	"midjourney_api/utils"
	"strings"
	"time"
)

type SubmitService struct {
	Ctx context.Context
}

// SubmitImagine 提交Imagine操作
func (s *SubmitService) SubmitImagine(p *req_params.SubmitImagineParams) error {
	task := s.newTask(p)
	task.Action = common.TaskActionIMAGINE
	prompt := strings.TrimSpace(p.Prompt)
	task.Prompt = prompt
	task.PromptEn = prompt // todo 翻译英文，敏感词过滤
	task.Description = "/imagine " + prompt
	m := store.InMemoryStore{
		Ctx: s.Ctx,
	}
	// 去请求api
	if err := discord_api.SenderReq(task.PromptEn); err != nil {
		return err
	}
	return m.Save(task)
}

// 新建任务
func (s *SubmitService) newTask(p *req_params.SubmitImagineParams) *task_model.TaskInfo {
	return &task_model.TaskInfo{
		ID:         utils.RandomStr(32),
		SubmitTime: time.Now().Format(time.DateTime),
		State:      p.State,
		NotifyHook: p.NotifyHook,
	}
}
