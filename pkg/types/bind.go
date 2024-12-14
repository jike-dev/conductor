package types

// BindRelation 绑定关系结构
type BindRelation struct {
	TargetUID      uint64                 `json:"target_uid"`       // 目标用户ID
	IsFirstConnect int                    `json:"is_first_connect"` // 是否首次绑定
	BindTime       string                 `json:"bind_time"`        // 绑定时间
	BindType       int                    `json:"bind_type"`        // 绑定类型
	Status         int                    `json:"status"`           // 绑定状态
	Extra          map[string]interface{} `json:"extra"`            // 扩展字段
}
