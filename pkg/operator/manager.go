package operator

import (
	"sync"

	"github.com/jike-dev/conductor/pkg/types"
)

// Manager 算子管理器实现
type Manager struct {
	mu        sync.RWMutex
	operators map[string]map[string]types.Operator // activityID -> operatorName -> operator
}

func NewManager() *Manager {
	return &Manager{
		operators: make(map[string]map[string]types.Operator),
	}
}

// RegisterOperator 注册算子
func (m *Manager) RegisterOperator(activityID string, operator types.Operator) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.operators[activityID]; !ok {
		m.operators[activityID] = make(map[string]types.Operator)
	}
	m.operators[activityID][operator.GetName()] = operator
}

// GetOperator 获取算子
func (m *Manager) GetOperator(activityID, operatorName string) (types.Operator, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if ops, ok := m.operators[activityID]; ok {
		op, ok := ops[operatorName]
		return op, ok
	}
	return nil, false
}
