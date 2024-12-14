package video

import (
	"github.com/jike-dev/conductor/pkg/dmp"
	"github.com/jike-dev/conductor/pkg/executor"
	"github.com/jike-dev/conductor/pkg/executor/adapter"
	"github.com/jike-dev/conductor/pkg/types"
)

// VideoAdapter 视频业务适配器，实现 executor.Dapter 接口
type VideoAdapter struct {
	*adapter.BaseAdapter
	dmpClient dmp.DmpClient
}

// NewVideoAdapter 创建视频适配器
func NewVideoAdapter() executor.Dapter {
	return &VideoAdapter{
		BaseAdapter: adapter.NewBaseAdapter("act_id"),
	}
}

// GetBindRelation 重写基础适配器的绑定关系获取方法，实现视频业务特有逻辑
func (a *VideoAdapter) GetBindRelation(ctx types.BusinessContext, req *types.ActivityRequest) (*types.BindRelation, error) {
	// 构建视频业务特有的DMP请求参数
	dmpParams := map[string]interface{}{
		"uid":        req.UID,
		"scene":      req.BusinessParams["scene"],
		"video_type": req.BusinessParams["video_type"],
	}
	return a.dmpClient.GetBindRelation(ctx, dmpParams)
}
