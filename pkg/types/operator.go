package types

// Operator 算子基础接口
type Operator interface {
	Execute(ctx BusinessContext, params map[string]interface{}) (bool, error)
	GetName() string
	GetType() string
}

// OperatorManager 算子管理器接口
type OperatorManager interface {
	RegisterOperator(activityID string, operator Operator)
	GetOperator(activityID, operatorName string) (Operator, bool)
}
