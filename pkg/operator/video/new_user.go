package video

import (
	"time"

	"github.com/jike-dev/conductor/pkg/types"
)

// NewUserOperator 视频新用户算子
type NewUserOperator struct{}

func NewNewUserOperator() types.Operator {
	return &NewUserOperator{}
}

func (o *NewUserOperator) Execute(ctx types.BusinessContext, params map[string]interface{}) (bool, error) {
	// 1. 获取当前活动的奖励结果
	reward, ok := ctx.GetActivityResult(ctx.GetConfig().ActivityID)
	if !ok || reward == nil {
		return false, nil
	}

	// 2. 执行业务逻辑判断是否是新用户
	isNewUser := true // 这里应该是实际的新用户判断逻辑

	// 3. 如果是新用户，直接修改结果
	if isNewUser {
		reward.VideoList.FirstActTime = time.Now().Format("2006-01-02 15:04:05")
		reward.VideoList.UserType = 1
	}

	return isNewUser, nil
}

func (o *NewUserOperator) GetName() string {
	return "video_new_user"
}

func (o *NewUserOperator) GetType() string {
	return "video"
}
