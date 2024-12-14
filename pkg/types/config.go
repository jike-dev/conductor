package types

// ActivityConfig 统一的活动配置结构
type ActivityConfig struct {
	ActivityID    string                 `json:"activity_id"`
	Type          string                 `json:"type"`           // 活动类型：video/board/item
	Status        int                    `json:"status"`         // 活动状态
	Version       string                 `json:"version"`        // 版本号
	Dependencies  []string               `json:"dependencies"`   // 依赖的活动ID列表
	Experiments   []string               `json:"experiments"`    // 实验列表
	TargetRules   []OperatorRule         `json:"target_rules"`   // 目标用户规则
	BusinessRules map[string]interface{} `json:"business_rules"` // 业务规则
}

// OperatorRule 算子规则
type OperatorRule struct {
	Type     string                 `json:"type"`     // 规则类型：operator/group
	Operator string                 `json:"operator"` // 算子名称（当type=operator时）
	Params   map[string]interface{} `json:"params"`   // 算子参数
	Rules    []OperatorRule         `json:"rules"`    // 子规则（当type=group时）
	Logic    string                 `json:"logic"`    // 组合逻辑：and/or（当type=group时）
}
