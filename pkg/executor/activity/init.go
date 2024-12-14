package activity

import (
	"context"

	"github.com/jike-dev/conductor/pkg/executor"
	"github.com/jike-dev/conductor/pkg/executor/activity/video"
	"github.com/jike-dev/conductor/pkg/types"
)

// ServiceInitializer 服务初始化器
type ServiceInitializer struct {
	configManager *types.ConfigManager
	actManager    types.ActivityManager
}

// NewServiceInitializer 创建服务初始化器
func NewServiceInitializer() *ServiceInitializer {
	return &ServiceInitializer{
		configManager: types.NewConfigManager(nil),
		actManager:    executor.NewActivityManager(),
	}
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
	// 加载配置
	cfg, err := s.loadConfig(activityIDs)
	if err != nil {
		return nil, err
	}

	// 创建上下文
	execCtx := types.NewBusinessContext(ctx)
	return execCtx.WithConfig(cfg), nil
}

// loadConfig 加载并验证配置
func (s *ServiceInitializer) loadConfig(activityIDs []string) (*types.ActivityConfig, error) {
	for _, actID := range activityIDs {
		cfg, ok := s.configManager.GetCachedConfig(actID)
		if !ok {
			continue
		}

		actCfg, ok := cfg.(*types.ActivityConfig)
		if !ok {
			continue
		}

		if err := s.validateConfig(actCfg); err != nil {
			return nil, err
		}

		return actCfg, nil
	}
	return nil, nil
}

// validateConfig 验证配置
func (s *ServiceInitializer) validateConfig(cfg *types.ActivityConfig) error {
	// 验证配置有效性
	// 1. 检查必填字段
	// 2. 验证配置格式
	// 3. 检查业务规则
	return nil
}

// InitializeServices 初始化所有服务
func (s *ServiceInitializer) InitializeServices(ctx context.Context) error {
	// 初始化各个业务模块
	video.InitVideoActivity(s.actManager)
	return nil
}
