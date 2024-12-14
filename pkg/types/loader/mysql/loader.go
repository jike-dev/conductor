package mysql

import (
	"context"
	"database/sql"
)

// MysqlLoader MySQL配置加载器
type MysqlLoader struct {
	db *sql.DB
}

func NewMysqlLoader(db *sql.DB) *MysqlLoader {
	return &MysqlLoader{db: db}
}

// LoadConfig 从MySQL加载配置
func (l *MysqlLoader) LoadConfig(ctx context.Context, activityID string) ([]byte, error) {
	var config []byte
	err := l.db.QueryRowContext(ctx,
		"SELECT config FROM activity_configs WHERE activity_id = ? AND status = 1",
		activityID,
	).Scan(&config)
	return config, err
}
