package adapter

import (
	"github.com/your-project/pkg/types"
)

// BaseAdapter 基础适配器实现
type BaseAdapter struct {
	activityID string
}

func NewBaseAdapter(activityID string) *BaseAdapter {
	return &BaseAdapter{
		activityID: activityID,
	}
}

// CheckVersion 通用版本检查
func (a *BaseAdapter) CheckVersion(ctx types.BusinessContext, version string) bool {
	cfg, ok := ctx.GetActivityConfig(a.activityID)
	if !ok {
		return false
	}
	return version >= cfg.Version
}

// CheckExperiment 通用实验检查
func (a *BaseAdapter) CheckExperiment(ctx types.BusinessContext, experiments []string) bool {
	cfg, ok := ctx.GetActivityConfig(a.activityID)
	if !ok {
		return false
	}
	return hasMatchingExperiment(cfg.Experiments, experiments)
}

// GetBindRelation 获取绑定关系（基础实现）
func (a *BaseAdapter) GetBindRelation(ctx types.BusinessContext, req *types.ActivityRequest) (*types.BindRelation, error) {
	// 基础实现返回空，由具体业务实现覆盖
	return nil, nil
}

// CheckExitStatus 检查退场状态
func (a *BaseAdapter) CheckExitStatus(ctx types.BusinessContext) bool {
	cfg, ok := ctx.GetActivityConfig(a.activityID)
	if !ok {
		return false
	}
	// 从配置中获取退场状态
	if exitStatus, ok := cfg.BusinessRules["exit_status"].(bool); ok {
		return exitStatus
	}
	return false
}

// CheckSwitchStatus 检查开关状态
func (a *BaseAdapter) CheckSwitchStatus(ctx types.BusinessContext, switchKey string) bool {
	cfg, ok := ctx.GetActivityConfig(a.activityID)
	if !ok {
		return false
	}
	// 从配置中获取开关状态
	if switches, ok := cfg.BusinessRules["switches"].(map[string]bool); ok {
		return switches[switchKey]
	}
	return false
}

// 辅助函数
func hasMatchingExperiment(cfgExps, reqExps []string) bool {
	for _, exp := range cfgExps {
		for _, reqExp := range reqExps {
			if exp == reqExp {
				return true
			}
		}
	}
	return false
}
