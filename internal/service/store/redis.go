// @Author: liyongzhen
// @Description:
// @File: redis
// @Date: 2023/8/18 17:06

package store

import (
	"context"
	"midjourney_api/internal/moldes/task_model"
)

const (
	REDIS_KEY_PREFIX = "mj-task-store::"
)

type RedisStore struct {
	Ctx context.Context
}

func (r RedisStore) Save(task *task_model.TaskInfo) error {
	//TODO implement me
	panic("implement me")
}

func (r RedisStore) Delete(taskId string) error {
	//TODO implement me
	panic("implement me")
}

func (r RedisStore) Get(taskId string) (*task_model.TaskInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (r RedisStore) List() ([]*task_model.TaskInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (r RedisStore) ListFilter(condition string) ([]*task_model.TaskInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (r RedisStore) FindOne(condition string) (*task_model.TaskInfo, error) {
	//TODO implement me
	panic("implement me")
}
