
# Actflow

Actflow 是一个灵活的业务活动编排和执行框架，支持复杂的业务规则处理、活动依赖管理和插件化的算子系统。

## 特性

- 活动编排：支持多活动的依赖关系管理和顺序执行
- 规则引擎：基于配置的规则执行系统，支持复杂的组合逻辑
- 插件化算子：可扩展的算子系统，支持业务自定义算子
- 业务适配：统一的适配器接口，方便不同业务接入
- 配置驱动：通过配置控制活动行为，支持热更新

## 架构设计

### 核心组件

- ActivityService: 活动服务，负责活动的编排和执行
- ActivityManager: 活动管理器，管理执行器和算子
- BusinessContext: 业务上下文，管理配置和数据流转
- RuleExecutor: 规则执行器，处理复杂的业务规则
- OperatorManager: 算子管理器，管理业务算子

### 目录结构
├── cmd/ # 命令行入口
├── pkg/ # 包目录
│ ├── conf/ # 配置管理
│ ├── dmp/ # DMP客户端
│ ├── executor/ # 执行器相关
│ │ ├── activity/ # 活动执行
│ │ ├── adapter/ # 适配器
│ │ └── rule/ # 规则执行
│ ├── operator/ # 算子相关
│ └── types/ # 公共类型定义



## 快速开始

### 安装

1. 创建活动配置：

{
"activity_id": "video_001",
"type": "video",
"status": 1,
"version": "1.0.0",
"dependencies": ["board_001"],
"experiments": ["exp1"],
"target_rules": [
{
"type": "operator",
"operator": "video_new_user",
"params": {"days": 7}
}
]
}
go
type VideoExecutor struct {
adapter executor.Dapter
actManager types.ActivityManager
}
func (e VideoExecutor) Execute(ctx types.BusinessContext, req types.ActivityRequest, reward types.TaskRewardList) error {
// 实现业务逻辑
return nil
}
go
func main() {
// 初始化服务
serviceInitializer := activity.NewServiceInitializer()
// 执行活动
req := &types.ActivityRequest{
ActNames: []string{"video_001"},
UID: 12345,
}
resp, err := serviceInitializer.Execute(ctx, req)
}


## 扩展开发

### 添加新的业务模块

1. 实现业务适配器
2. 实现业务执行器
3. 实现业务算子
4. 注册到活动管理器

### 自定义算子


type CustomOperator struct{}
func (o CustomOperator) Execute(ctx types.BusinessContext, params map[string]interface{}) (bool, error) {
// 实现算子逻辑
return true, nil
}
)


## 配置说明

### 活动配置

- activity_id: 活动唯一标识
- type: 活动类型
- dependencies: 依赖的活动列表
- target_rules: 目标规则配置

### 规则配置

- type: 规则类型 (operator/group)
- operator: 算子名称
- params: 算子参数
- logic: 组合逻辑 (and/or)


这个 README 包含了：
项目简介和特性
架构设计说明
快速开始指南
扩展开发说明
配置文档
贡献指南

