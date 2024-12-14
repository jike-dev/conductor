package types

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
)

// ConfigManager 配置管理器
type ConfigManager struct {
	mu      sync.RWMutex
	configs map[string]interface{}
	loader  ConfigLoader
}

// ConfigLoader 配置加载器接口
type ConfigLoader interface {
	LoadConfig(ctx context.Context, activityID string) ([]byte, error)
}

func NewConfigManager(loader ConfigLoader) *ConfigManager {
	return &ConfigManager{
		configs: make(map[string]interface{}),
		loader:  loader,
	}
}

func (m *ConfigManager) GetConfig(ctx context.Context, activityID string) (interface{}, error) {
	m.mu.RLock()
	if config, ok := m.configs[activityID]; ok {
		m.mu.RUnlock()
		return config, nil
	}
	m.mu.RUnlock()

	return m.loadAndCacheConfig(ctx, activityID)
}

func (m *ConfigManager) GetCachedConfig(activityID string) (interface{}, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	cfg, ok := m.configs[activityID]
	return cfg, ok
}

func (m *ConfigManager) loadAndCacheConfig(ctx context.Context, activityID string) (interface{}, error) {
	data, err := m.loader.LoadConfig(ctx, activityID)
	if err != nil {
		return nil, fmt.Errorf("load config error: %w", err)
	}

	var cfg ActivityConfig
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("unmarshal config error: %w", err)
	}

	m.mu.Lock()
	m.configs[activityID] = &cfg
	m.mu.Unlock()

	return &cfg, nil
}
