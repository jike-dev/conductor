package video

import (
	"github.com/jike-dev/conductor/pkg/executor/rule"

	"github.com/jike-dev/conductor/pkg/executor"
	"github.com/jike-dev/conductor/pkg/types"
)

// VideoExecutor 视频业务执行器
type VideoExecutor struct {
	adapter    executor.Dapter // 使用基础处理器接口
	actManager types.ActivityManager
}

// NewVideoExecutor 创建视频执行器
func NewVideoExecutor(actManager types.ActivityManager) *VideoExecutor {
	// 创建视频适配器
	return &VideoExecutor{
		adapter:    NewVideoAdapter(), // 创建视频业务适配器
		actManager: actManager,
	}
}

func (e *VideoExecutor) Execute(ctx types.BusinessContext, req *types.ActivityRequest, reward *types.TaskRewardList) error {
	cfg := ctx.GetConfig()
	if cfg == nil {
		return nil
	}

	// 检查退场状态
	if e.adapter.CheckExitStatus(ctx) {
		return nil
	}

	// 获取绑定关系
	bindRelation, err := e.adapter.GetBindRelation(ctx, req)
	if err != nil {
		return err
	}

	// 使用绑定关系更新奖励信息
	if bindRelation != nil {
		reward.VideoList.IsFirstConnect = bindRelation.IsFirstConnect
	}

	// 创建规则执行器
	ruleExecutor := rule.NewRuleExecutor(e.actManager.GetOperatorManager())

	// 执行目标用户规则
	isTarget, err := ruleExecutor.ExecuteRules(ctx, cfg.TargetRules, req.BusinessParams)
	if err != nil {
		return err
	}

	if isTarget {
		// 处理目标用户逻辑
		reward.VideoList.IsTargetUser = 1
		reward.VideoList.UserActLevel = calculateUserActLevel(ctx)
		reward.VideoList.PopUp = calculatePopUp(ctx)
	}

	return nil
}

// 业务相关的辅助函数
func calculateUserActLevel(ctx types.BusinessContext) int {
	return 0
}

func calculatePopUp(ctx types.BusinessContext) int {
	return 0
}
