package opertator

import "github.com/jike-dev/conductor/pkg/types"

type VideoNewUserOperator struct{}

func NewVideoNewUserOperator() types.Operator {
	return &VideoNewUserOperator{}
}

func (o *VideoNewUserOperator) Execute(ctx types.BusinessContext, params map[string]interface{}) (bool, error) {
	// 实现新用户判断逻辑
	return true, nil
}

func (o *VideoNewUserOperator) GetName() string {
	return "video_new_user"
}

func (o *VideoNewUserOperator) GetType() string {
	return "video"
}
