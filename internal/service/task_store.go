// @Author: liyongzhen
// @Description:
// @File: task_store
// @Date: 2023/8/18 17:06

package service

import "midjourney_api/internal/moldes/task_model"

type TaskStore interface {
	Save(task *task_model.TaskInfo) error                        // 保存任务
	Delete(taskId string) error                                  // 删除任务
	Get(taskId string) (*task_model.TaskInfo, error)             // 获取单个任务
	List() ([]*task_model.TaskInfo, error)                       // 获取所有任务
	ListFilter(condition string) ([]*task_model.TaskInfo, error) // 根据条件筛选任务列表
	FindOne(condition string) (*task_model.TaskInfo, error)      // 根据条件筛选某个任务
}
