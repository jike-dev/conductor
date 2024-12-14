package context

import (
	"context"
	"sync"
	"time"

	"github.com/jike-dev/conductor/pkg/types"
)

// BusinessContext 业务上下文实现
type BusinessContext struct {
	ctx     context.Context
	mu      sync.RWMutex
	configs map[string]*types.ActivityConfig
	data    map[string]interface{}
	results map[string]*types.TaskRewardList
}

// NewBusinessContext 创建业务上下文
func NewBusinessContext(ctx context.Context) *BusinessContext {
	return &BusinessContext{
		ctx:     ctx,
		configs: make(map[string]*types.ActivityConfig),
		data:    make(map[string]interface{}),
		results: make(map[string]*types.TaskRewardList),
	}
}

// 实现 context.Context 接口
func (bc *BusinessContext) Deadline() (deadline time.Time, ok bool) {
	return bc.ctx.Deadline()
}

func (bc *BusinessContext) Done() <-chan struct{} {
	return bc.ctx.Done()
}

func (bc *BusinessContext) Err() error {
	return bc.ctx.Err()
}

func (bc *BusinessContext) Value(key interface{}) interface{} {
	return bc.ctx.Value(key)
}

// 实现配置相关方法
func (bc *BusinessContext) GetActivityConfig(activityID string) (*types.ActivityConfig, bool) {
	bc.mu.RLock()
	defer bc.mu.RUnlock()
	cfg, ok := bc.configs[activityID]
	return cfg, ok
}

func (bc *BusinessContext) GetAllConfigs() map[string]*types.ActivityConfig {
	bc.mu.RLock()
	defer bc.mu.RUnlock()
	configs := make(map[string]*types.ActivityConfig)
	for k, v := range bc.configs {
		configs[k] = v
	}
	return configs
}

func (bc *BusinessContext) SetActivityConfig(activityID string, cfg *types.ActivityConfig) {
	bc.mu.Lock()
	defer bc.mu.Unlock()
	bc.configs[activityID] = cfg
}

// 实现业务数据相关方法
func (bc *BusinessContext) GetBusinessData(key string) (interface{}, bool) {
	bc.mu.RLock()
	defer bc.mu.RUnlock()
	value, ok := bc.data[key]
	return value, ok
}

func (bc *BusinessContext) SetBusinessData(key string, value interface{}) {
	bc.mu.Lock()
	defer bc.mu.Unlock()
	bc.data[key] = value
}

// 实现结果相关方法
func (bc *BusinessContext) GetActivityResult(activityID string) (*types.TaskRewardList, bool) {
	bc.mu.RLock()
	defer bc.mu.RUnlock()
	result, ok := bc.results[activityID]
	return result, ok
}

func (bc *BusinessContext) SetActivityResult(activityID string, result *types.TaskRewardList) {
	bc.mu.Lock()
	defer bc.mu.Unlock()
	bc.results[activityID] = result
}
