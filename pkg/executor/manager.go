package executor

import (
	"sync"

	"github.com/your-project/pkg/operator"
	"github.com/your-project/pkg/types"
)

// ActivityManager 活动管理器实现
type ActivityManager struct {
	mu        sync.RWMutex
	executors map[string]types.ActivityExecutor
	opManager *operator.OperatorManager
}

func NewActivityManager() *ActivityManager {
	return &ActivityManager{
		executors: make(map[string]types.ActivityExecutor),
		opManager: operator.NewOperatorManager(),
	}
}

// RegisterExecutor 注册执行器
func (m *ActivityManager) RegisterExecutor(activityID string, executor types.ActivityExecutor) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.executors[activityID] = executor
}

// GetExecutor 获取执行器
func (m *ActivityManager) GetExecutor(activityID string) (types.ActivityExecutor, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	exec, ok := m.executors[activityID]
	return exec, ok
}

// GetOperatorManager 获取算子管理器
func (m *ActivityManager) GetOperatorManager() *operator.OperatorManager {
	return m.opManager
}
