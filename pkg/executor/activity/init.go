package activity

import (
	"context"

	"github.com/jike-dev/conductor/pkg/executor"
	"github.com/jike-dev/conductor/pkg/executor/video"
	"github.com/jike-dev/conductor/pkg/operator"
	"github.com/jike-dev/conductor/pkg/types"
)

// ServiceInitializer 服务初始化器
type ServiceInitializer struct {
	configManager *types.ConfigManager
	actManager    types.ActivityManager
	opManager     types.OperatorManager
}

// NewServiceInitializer 创建服务初始化器
func NewServiceInitializer() *ServiceInitializer {
	return &ServiceInitializer{
		configManager: types.NewConfigManager(nil),
		actManager:    executor.NewActivityManager(),
		opManager:     operator.NewManager(),
	}
}

// InitializeServices 初始化所有服务
func (s *ServiceInitializer) InitializeServices(ctx context.Context) error {
	// 初始化视频模块
	videoExecutor := video.InitModule(s.opManager)
	s.actManager.RegisterExecutor("video_001", videoExecutor)

	return nil
}

// Execute 执行活动
func (s *ServiceInitializer) Execute(ctx context.Context, req *types.ActivityRequest) (*types.ActivityResponse, error) {
	// 1. 准备执行上下文
	execCtx, err := s.prepareContext(ctx, req.ActNames)
	if err != nil {
		return nil, err
	}

	// 2. 执行活动
	svc := NewActivityService(s.configManager, s.actManager)
	return svc.Execute(execCtx, req)
}

// prepareContext 准备执行上下文
func (s *ServiceInitializer) prepareContext(ctx context.Context, activityIDs []string) (types.BusinessContext, error) {
	execCtx := types.NewBusinessContext(ctx)

	// 加载每个活动的配置
	for _, actID := range activityIDs {
		cfg, err := s.configManager.GetConfig(ctx, actID)
		if err != nil {
			continue
		}
		if actCfg, ok := cfg.(*types.ActivityConfig); ok {
			execCtx.SetActivityConfig(actID, actCfg)
		}
	}

	return execCtx, nil
}
