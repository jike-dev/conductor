package rule

import (
	"github.com/jike-dev/conductor/pkg/types"
)

// RuleExecutor 规则执行器
type RuleExecutor struct {
	opManager types.OperatorManager
}

func NewRuleExecutor(opManager types.OperatorManager) *RuleExecutor {
	return &RuleExecutor{
		opManager: opManager,
	}
}

// ExecuteRules 执行规则组
func (e *RuleExecutor) ExecuteRules(ctx types.BusinessContext, rules []types.OperatorRule, params map[string]interface{}) (bool, error) {
	results := make([]bool, 0)

	for _, rule := range rules {
		result, err := e.executeRule(ctx, rule, params)
		if err != nil {
			return false, err
		}
		results = append(results, result)

		// 短路逻辑
		if rule.Logic == "or" && result {
			return true, nil
		} else if rule.Logic == "and" && !result {
			return false, nil
		}
	}

	if len(rules) > 0 && rules[0].Logic == "and" {
		return allTrue(results), nil
	}
	return anyTrue(results), nil
}

// executeRule 执行单个规则
func (e *RuleExecutor) executeRule(ctx types.BusinessContext, rule types.OperatorRule, params map[string]interface{}) (bool, error) {
	switch rule.Type {
	case "operator":
		return e.executeOperator(ctx, rule.Operator, rule.Params)
	case "group":
		return e.ExecuteRules(ctx, rule.Rules, params)
	default:
		return false, nil
	}
}

// executeOperator 执行单个算子
func (e *RuleExecutor) executeOperator(ctx types.BusinessContext, operatorName string, params map[string]interface{}) (bool, error) {
	cfg := ctx.GetConfig()
	if cfg == nil {
		return false, nil
	}

	op, ok := e.opManager.GetOperator(cfg.ActivityID, operatorName)
	if !ok {
		return false, nil
	}
	return op.Execute(ctx, params)
}

// 辅助函数
func allTrue(results []bool) bool {
	for _, r := range results {
		if !r {
			return false
		}
	}
	return true
}

func anyTrue(results []bool) bool {
	for _, r := range results {
		if r {
			return true
		}
	}
	return false
}
