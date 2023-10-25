package db

import "github.com/ryker-w/go-init/internal/db/model"

// 初始化数据库表
func RegisterTables() (tables []interface{}) {
	tables = append(tables,
		new(model.UserInfo),
	)
	return
}
