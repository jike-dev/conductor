package types

// ActivityRequest 统一请求结构
type ActivityRequest struct {
	ActNames       []string               `json:"act_names"`       // 活动名称列表
	UID            uint64                 `json:"uid"`             // 用户ID
	BusinessParams map[string]interface{} `json:"business_params"` // 业务参数
}
