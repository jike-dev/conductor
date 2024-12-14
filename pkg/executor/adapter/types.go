package adapter

import (
	"github.com/your-project/pkg/types"
)

// BusinessAdapter 业务适配器接口
type BusinessAdapter interface {
	// 版本检查
	CheckVersion(ctx types.BusinessContext, version string) bool
	// 实验检查
	CheckExperiment(ctx types.BusinessContext, experiments []string) bool
	// 获取业务类型
	GetType() string
}
