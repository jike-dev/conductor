package types

// TaskRewardList 统一返回结构
type TaskRewardList struct {
	BoardRwxtList        BoardRwxtControl     `json:"rwxtcontrol"`
	VideoList            VideoControl         `json:"videocontrol"`
	ItemIncentiveControl ItemIncentiveControl `json:"itemincentivecontrol"`
	Ext                  map[string]string    `json:"ext"`
}

// BoardRwxtControl 红包弹框控制
type BoardRwxtControl struct {
	AppFirstActDmp string `json:"bdapp_first_act_dmp"`
	AppFirstAct    string `json:"bdapp_first_act"`
	AppLastAct     string `json:"bdapp_last_act"`
}

// VideoControl 视频活动控制
type VideoControl struct {
	TargetFirstAct string `json:"target_first_act"` // 首次成为目标用户
	FirstActTime   string `json:"first_act_time"`   // 端激活时间
	UserType       int    `json:"user_type"`        // 返回用户类型
}

// ItemIncentiveControl 实物激励控制
type ItemIncentiveControl struct {
	TargetFirstAct string `json:"target_first_act"` // 首次成为目标用户
	FirstActTime   string `json:"first_act_time"`   // 端激活时间
	UserType       int    `json:"user_type"`        // 返回用户类型
}
