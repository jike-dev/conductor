package video

import (
	"github.com/your-project/pkg/executor/activity/video/opertator"
	"github.com/your-project/pkg/types"
)

// InitVideoActivity 初始化视频活动
func InitVideoActivity(actManager types.ActivityManager) {
	// 1. 注册执行器
	videoExecutor := NewVideoExecutor(actManager)
	actManager.RegisterExecutor("video_001", videoExecutor)

	// 2. 注册算子
	opManager := actManager.GetOperatorManager()
	opManager.RegisterOperator("video_001", opertator.NewVideoNewUserOperator())
	// 注册其他算子
}
