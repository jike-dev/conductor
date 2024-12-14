package main

import (
	"context"
	"database/sql"
	"log"

	"github.com/your-project/pkg/executor/activity"
	"github.com/your-project/pkg/types"
)

func main() {
	// 初始化数据库连接
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/dbname")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 2. 初始化服务
	serviceInitializer := activity.NewServiceInitializer()
	if err := serviceInitializer.InitializeServices(context.Background()); err != nil {
		log.Fatal(err)
	}

	// 3. 处理请求
	ctx := context.Background()
	req := &types.ActivityRequest{
		ActNames: []string{"reward_points", "coupon"},
		UID:      12345,
	}
	_, err = serviceInitializer.Execute(ctx, req)
	if err != nil {
		log.Fatal(err)
	}

	return

}
