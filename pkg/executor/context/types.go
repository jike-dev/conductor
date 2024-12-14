package context

import (
	"context"
	"sync"
	"time"

	"github.com/your-project/pkg/types"
)

// BusinessContext 业务上下文
type BusinessContext struct {
	ctx     context.Context
	mu      sync.RWMutex
	configs map[string]*types.ActivityConfig // 活动配置
	results map[string]*types.TaskRewardList // 业务结果
	data    map[string]interface{}           // 业务数据
}

// NewBusinessContext 创建业务上下文
func NewBusinessContext(ctx context.Context) *BusinessContext {
	return &BusinessContext{
		ctx:     ctx,
		configs: make(map[string]*types.ActivityConfig),
		results: make(map[string]*types.TaskRewardList),
		data:    make(map[string]interface{}),
	}
}

// SetConfig 设置活动配置
func (bc *BusinessContext) SetConfig(activityID string, cfg *types.ActivityConfig) {
	bc.mu.Lock()
	defer bc.mu.Unlock()
	bc.configs[activityID] = cfg
}

// GetConfig 获取活动配置
func (bc *BusinessContext) GetConfig(activityID string) (*types.ActivityConfig, bool) {
	bc.mu.RLock()
	defer bc.mu.RUnlock()
	cfg, ok := bc.configs[activityID]
	return cfg, ok
}

// SetResult 设置业务结果
func (bc *BusinessContext) SetResult(activityID string, result *types.TaskRewardList) {
	bc.mu.Lock()
	defer bc.mu.Unlock()
	bc.results[activityID] = result
}

// GetResult 获取业务结果
func (bc *BusinessContext) GetResult(activityID string) (*types.TaskRewardList, bool) {
	bc.mu.RLock()
	defer bc.mu.RUnlock()
	result, ok := bc.results[activityID]
	return result, ok
}

// SetData 设置业务数据
func (bc *BusinessContext) SetData(key string, value interface{}) {
	bc.mu.Lock()
	defer bc.mu.Unlock()
	bc.data[key] = value
}

// GetData 获取业务数据
func (bc *BusinessContext) GetData(key string) (interface{}, bool) {
	bc.mu.RLock()
	defer bc.mu.RUnlock()
	value, ok := bc.data[key]
	return value, ok
}

// Context 获取原始context
func (bc *BusinessContext) Context() context.Context {
	return bc.ctx
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
