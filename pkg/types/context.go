package types

import "context"

// BusinessContext 业务上下文接口
type BusinessContext interface {
	context.Context
	// 配置相关
	GetActivityConfig(activityID string) (*ActivityConfig, bool)
	SetActivityConfig(activityID string, cfg *ActivityConfig)
	WithConfig(cfg *ActivityConfig) BusinessContext // 返回带有新配置的上下文
	GetConfig() *ActivityConfig                     // 获取当前配置
	// 业务数据相关
	GetBusinessData(key string) (interface{}, bool)
	SetBusinessData(key string, value interface{})
	// 结果相关
	GetActivityResult(activityID string) (*TaskRewardList, bool)
	SetActivityResult(activityID string, result *TaskRewardList)
}

// businessContext 业务上下文实现
type businessContext struct {
	context.Context
	config  *ActivityConfig
	configs map[string]*ActivityConfig
	data    map[string]interface{}
	results map[string]*TaskRewardList
}

func NewBusinessContext(ctx context.Context) BusinessContext {
	return &businessContext{
		Context: ctx,
		configs: make(map[string]*ActivityConfig),
		data:    make(map[string]interface{}),
		results: make(map[string]*TaskRewardList),
	}
}

func (c *businessContext) WithConfig(cfg *ActivityConfig) BusinessContext {
	return &businessContext{
		Context: c.Context,
		config:  cfg,
		data:    c.data,
		results: c.results,
	}
}

func (c *businessContext) GetConfig() *ActivityConfig {
	return c.config
}

func (c *businessContext) GetBusinessData(key string) (interface{}, bool) {
	value, ok := c.data[key]
	return value, ok
}

func (c *businessContext) SetBusinessData(key string, value interface{}) {
	c.data[key] = value
}

func (c *businessContext) GetActivityResult(activityID string) (*TaskRewardList, bool) {
	result, ok := c.results[activityID]
	return result, ok
}

func (c *businessContext) SetActivityResult(activityID string, result *TaskRewardList) {
	c.results[activityID] = result
}

func (c *businessContext) GetActivityConfig(activityID string) (*ActivityConfig, bool) {
	cfg, ok := c.configs[activityID]
	return cfg, ok
}

func (c *businessContext) SetActivityConfig(activityID string, cfg *ActivityConfig) {
	c.configs[activityID] = cfg
}
