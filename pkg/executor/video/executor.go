package video

import (
	"github.com/jike-dev/conductor/pkg/types"
)

// Executor 视频业务执行器
type Executor struct {
	opManager types.OperatorManager
}

// NewExecutor 创建视频执行器
func NewExecutor(opManager types.OperatorManager) *Executor {
	return &Executor{
		opManager: opManager,
	}
}

func (e *Executor) Execute(ctx types.BusinessContext, req *types.ActivityRequest, reward *types.TaskRewardList) error {
	cfg := ctx.GetConfig()
	if cfg == nil {
		return nil
	}

	// 执行目标用户规则
	op, ok := e.opManager.GetOperator(cfg.ActivityID, "video_new_user")
	if !ok {
		return nil
	}

	// 执行算子
	isTarget, err := op.Execute(ctx, req.BusinessParams)
	if err != nil {
		return err
	}

	// 算子已经直接修改了reward中的数据，这里不需要额外处理
	_ = isTarget

	return nil
}
