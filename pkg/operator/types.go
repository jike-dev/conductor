package operator

import "context"

// Operator 算子基础接口
type Operator interface {
	Execute(ctx context.Context, params map[string]interface{}) (bool, error)
	GetName() string
	GetType() string // 返回算子类型：video/board/item
}

// OperatorFactory 算子工厂接口
type OperatorFactory interface {
	CreateOperator(name string) (Operator, error)
}
