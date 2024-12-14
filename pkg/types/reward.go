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
	ListRation     string `json:"lost_ratio"`
	UserType       int    `json:"user_type"`
	PopUp          int    `json:"popup"`
	Timer          int    `json:"timer"`
	IsTargetUser   int    `json:"istargetuser"`
	Channel        string `json:"channel"`
	IsNeedRecoder  int    `json:"isneedrecoder"`
	Sid            string `json:"sid"`
	Ctr            string `json:"ctr"`
	IsFirstConnect int    `json:"isfirstconnect"`
	IsTimerOn      int    `json:"istimeron"`
	IsOut          string `json:"isout"`
	OutNum         int    `json:"outnum"`
	Source         string `json:"source"`
	TargetUID      uint64 `json:"targetuid"`
}

// VideoControl 视频活动控制
type VideoControl struct {
	TargetFirstAct       string `json:"target_first_act"`       // 首次成为目标用户
	FirstActTime         string `json:"first_act_time"`         // 端激活时间
	UserType             int    `json:"user_type"`              // 返回用户类型
	UserActLevel         int    `json:"user_act_level"`         // 返回用户活跃等级
	UserPenetrationLevel int    `json:"user_penetration_level"` // 返回用户活跃等级(渗透)
	UserActLevelJump     int    `json:"user_act_level_jump"`    // 返回用户二跳场景等级
	PopUp                int    `json:"popup"`                  // 是否红包弹框
	IsTargetUser         int    `json:"istargetuser"`           // 是否目标用户
	Timer                int    `json:"timer"`                  // 是否显示计时器
	IsNeedRecoder        int    `json:"isneedrecoder"`          // 是否记录日志
	Sid                  string `json:"sid"`                    // 最终命中的sid
	IsFirstConnect       int    `json:"isfirstconnect"`         // 是否首次绑定目标用户关系
	IsInActivate         int    `json:"isinactivate"`           // 是否失活
	IsFreeze             int    `json:"isfreeze"`               // 是否退场
	NoLoginReward        int    `json:"nologinreward"`          // 是否未登录发放奖励
	Force                string `json:"force"`                  // 是否强制成为目标用户
	TargetFilterReason   int    `json:"target_filter_reason"`   // 目标用户过滤的原因
	IsGuide              int    `json:"isGuide"`                // 是否在首页进行导流
}

// ItemIncentiveControl 实物激励控制
type ItemIncentiveControl struct {
	TargetFirstAct     string `json:"target_first_act"`     // 首次成为目标用户
	FirstActTime       string `json:"first_act_time"`       // 端激活时间
	UserType           int    `json:"user_type"`            // 返回用户类型
	IsTargetUser       int    `json:"istargetuser"`         // 是否目标用户
	IsNeedRecoder      int    `json:"isneedrecoder"`        // 是否记录日志
	Sid                string `json:"sid"`                  // 最终命中的sid
	IsFirstConnect     int    `json:"isfirstconnect"`       // 是否首次绑定目标用户关系
	IsOut              string `json:"isout"`                // 是否退场
	OutNum             int    `json:"outnum"`               // 退场次数
	NoLoginReward      int    `json:"nologinreward"`        // 是否未登录发放奖励
	Force              string `json:"force"`                // 是否强制成为目标用户
	TargetFilterReason int    `json:"target_filter_reason"` // 目标用户过滤的原因
}
