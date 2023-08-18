// @Author: liyongzhen
// @Description:
// @File: memory
// @Date: 2023/8/18 17:06

package store

import (
	"context"
	"midjourney_api/internal/moldes/task_model"
)

var TaskMap = map[string]*task_model.TaskInfo{} // 缓存全局变量

type InMemoryStore struct {
	Ctx context.Context
}

func (i InMemoryStore) Save(task *task_model.TaskInfo) error {
	if TaskMap == nil {
		TaskMap = make(map[string]*task_model.TaskInfo)
	}
	TaskMap[task.ID] = task
	return nil
}

func (i InMemoryStore) Delete(taskId string) error {
	if TaskMap == nil {
		return nil
	}
	delete(TaskMap, taskId)
	return nil
}

func (i InMemoryStore) Get(taskId string) (*task_model.TaskInfo, error) {
	if TaskMap == nil {
		return nil, nil
	}
	if v, ok := TaskMap[taskId]; ok {
		return v, nil
	}
	return nil, nil
}

func (i InMemoryStore) List() ([]*task_model.TaskInfo, error) {
	if TaskMap == nil {
		return nil, nil
	}
	// map转list
	var list []*task_model.TaskInfo
	for _, v := range TaskMap {
		list = append(list, v)
	}
	return list, nil
}

func (i InMemoryStore) ListFilter(condition string) ([]*task_model.TaskInfo, error) {
	if TaskMap == nil {
		return nil, nil
	}
	// map转list
	var list []*task_model.TaskInfo
	for _, v := range TaskMap {
		if v.PromptEn == condition {
			list = append(list, v)
		}
	}
	return list, nil
}

func (i InMemoryStore) FindOne(condition string) (*task_model.TaskInfo, error) {
	if TaskMap == nil {
		return nil, nil
	}
	for _, v := range TaskMap {
		if v.PromptEn == condition {
			return v, nil
		}
	}
	return nil, nil
}
