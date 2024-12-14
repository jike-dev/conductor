package activity

import (
	"fmt"

	"github.com/jike-dev/conductor/pkg/types"
)

// ActivityService 活动服务
type ActivityService struct {
	configManager *types.ConfigManager
	actManager    types.ActivityManager
}

// activityInfo 活动信息，用于排序
type activityInfo struct {
	id     string
	config *types.ActivityConfig
}

func NewActivityService(
	configManager *types.ConfigManager,
	actManager types.ActivityManager,
) *ActivityService {
	return &ActivityService{
		configManager: configManager,
		actManager:    actManager,
	}
}

func (s *ActivityService) Execute(ctx types.BusinessContext, req *types.ActivityRequest) (*types.ActivityResponse, error) {
	// 1. 参数校验
	if err := s.validateRequest(req); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	// 2. 创建响应结构
	resp := &types.ActivityResponse{
		Data: &types.TaskRewardList{},
	}

	// 3. 获取并排序活动
	activities, err := s.getSortedActivities(ctx, req.ActNames)
	if err != nil {
		return nil, fmt.Errorf("get sorted activities error: %w", err)
	}

	// 4. 按顺序执行活动
	for _, act := range activities {
		// 获取执行器
		executor, ok := s.actManager.GetExecutor(act.id)
		if !ok {
			continue
		}

		// 创建活动上下文
		actCtx := ctx.WithConfig(act.config)

		// 执行活动
		if err := executor.Execute(actCtx, req, resp.Data); err != nil {
			resp.Message = fmt.Sprintf("execute activity %s error: %v", act.id, err)
			continue
		}
	}

	// 5. 处理响应
	if resp.Code == 0 {
		resp.Message = "success"
	}
	return resp, nil
}

// getSortedActivities 获取排序后的活动列表
func (s *ActivityService) getSortedActivities(ctx types.BusinessContext, activityIDs []string) ([]activityInfo, error) {
	// 构建依赖图
	graph := make(map[string][]string)
	inDegree := make(map[string]int)
	activities := make(map[string]*types.ActivityConfig)

	// 收集活动信息和构建依赖关系
	for _, actID := range activityIDs {
		cfg, ok := ctx.GetActivityConfig(actID)
		if !ok {
			continue
		}
		activities[actID] = cfg
		inDegree[actID] = 0
		graph[actID] = cfg.Dependencies
	}

	// 计算入度
	for _, deps := range graph {
		for _, dep := range deps {
			if _, exists := activities[dep]; !exists {
				return nil, fmt.Errorf("dependency activity %s not found", dep)
			}
			inDegree[dep]++
		}
	}

	// 拓扑排序
	var sorted []activityInfo
	var queue []string

	// 找出入度为0的节点
	for actID := range activities {
		if inDegree[actID] == 0 {
			queue = append(queue, actID)
		}
	}

	// 执行拓扑排序
	for len(queue) > 0 {
		actID := queue[0]
		queue = queue[1:]

		sorted = append(sorted, activityInfo{
			id:     actID,
			config: activities[actID],
		})

		// 更新依赖此活动的其他活动的入度
		for _, dep := range graph[actID] {
			inDegree[dep]--
			if inDegree[dep] == 0 {
				queue = append(queue, dep)
			}
		}
	}

	// 检查是否存在循环依赖
	if len(sorted) != len(activities) {
		return nil, fmt.Errorf("circular dependency detected")
	}

	return sorted, nil
}

// validateRequest 验证请求参数
func (s *ActivityService) validateRequest(req *types.ActivityRequest) error {
	if req == nil {
		return fmt.Errorf("request is nil")
	}
	if len(req.ActNames) == 0 {
		return fmt.Errorf("activity names is empty")
	}
	if req.UID == 0 {
		return fmt.Errorf("uid is required")
	}
	return nil
}
