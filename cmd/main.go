package main

import (
	"context"
	"log"

	"github.com/jike-dev/conductor/pkg/executor/activity"
	"github.com/jike-dev/conductor/pkg/types"
)

func main() {
	// 1. 创建服务初始化器
	serviceInitializer := activity.NewServiceInitializer()

	// 2. 初始化所有服务（包括视频模块）
	if err := serviceInitializer.InitializeServices(context.Background()); err != nil {
		log.Fatal(err)
	}

	// 3. 处理请求
	ctx := context.Background()
	req := &types.ActivityRequest{
		ActNames: []string{"video_001"}, // 使用视频模块的活动ID
		UID:      12345,
		BusinessParams: map[string]interface{}{
			"scene": "home",
		},
	}

	// 4. 执行活动
	resp, err := serviceInitializer.Execute(ctx, req)
	if err != nil {
		log.Fatal(err)
	}

	// 5. 处理响应
	if resp.Code == 0 {
		log.Printf("执行成功: %+v\n", resp.Data)
	} else {
		log.Printf("执行失败: %s\n", resp.Message)
	}
}
