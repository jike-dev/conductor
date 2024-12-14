package dmp

import (
	"context"

	"github.com/jike-dev/conductor/pkg/types"
)

// DmpClient DMP客户端接口
type DmpClient interface {
	GetBindRelation(ctx context.Context, params map[string]interface{}) (*types.BindRelation, error)
}
