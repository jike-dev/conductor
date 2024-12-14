package types

import "github.com/jike-dev/conductor/pkg/operator"

// ActivityManager 活动管理器接口
type ActivityManager interface {
	// 注册活动执行器
	RegisterExecutor(activityID string, executor ActivityExecutor)
	// 获取活动执行器
	GetExecutor(activityID string) (ActivityExecutor, bool)
	// 获取算子管理器
	GetOperatorManager() *operator.OperatorManager
}

// ActivityExecutor 活动执行器接口
type ActivityExecutor interface {
	Execute(ctx BusinessContext, req *ActivityRequest, reward *TaskRewardList) error
}

// ActivityResponse 统一响应结构
type ActivityResponse struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    *TaskRewardList `json:"data"` // 单个任务的奖励结构
}

// Operator 算子接口
type Operator interface {
	Execute(ctx BusinessContext, params map[string]interface{}) (bool, error)
	GetName() string
	GetType() string
}
