package video

import (
	"github.com/jike-dev/conductor/pkg/operator/video"
	"github.com/jike-dev/conductor/pkg/types"
)

// InitModule 初始化视频模块
func InitModule(opManager types.OperatorManager) types.ActivityExecutor {
	// 1. 注册算子
	opManager.RegisterOperator("video_new_user", video.NewNewUserOperator())

	// 2. 创建并返回执行器
	return NewExecutor(opManager)
}
