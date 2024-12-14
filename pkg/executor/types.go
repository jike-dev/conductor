package executor

import (
	"context"

	"github.com/jike-dev/conductor/pkg/types"
)

// Dapter 基础处理接口
type Dapter interface {
	// 版本检查
	CheckVersion(ctx types.BusinessContext, version string) bool
	// 实验检查
	CheckExperiment(ctx types.BusinessContext, experiments []string) bool
	// 获取绑定关系
	GetBindRelation(ctx types.BusinessContext, req *types.ActivityRequest) (*types.BindRelation, error)
	// 检查退场状态
	CheckExitStatus(ctx types.BusinessContext) bool
	// 检查开关状态
	CheckSwitchStatus(ctx types.BusinessContext, switchKey string) bool
}

// ActivityExecutor 活动执行器接口
type ActivityExecutor interface {
	Execute(ctx types.BusinessContext, req *types.ActivityRequest, reward *types.TaskRewardList) error
}

// ActivityResponse 统一响应结构
type ActivityResponse struct {
	Code    int                   `json:"code"`
	Message string                `json:"message"`
	Data    *types.TaskRewardList `json:"data"` // 单个任务的奖励结构
}

// IExecutorManager 执行器管理器接口
type IExecutorManager interface {
	RegisterExecutor(actName string, executor ActivityExecutor)
	Execute(ctx context.Context, req *types.ActivityRequest) (*ActivityResponse, error)
}
