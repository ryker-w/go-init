package model

import (
	"fmt"
	"time"
)

type Pk struct {
	// ID
	Id int `orm:"pk;auto;column(id)"`
}

func (pk *Pk) PkString() string {
	return fmt.Sprintf("%d", pk.Id)
}

type TableChangeInfo struct {
	// 状态
	Status int `orm:"column(status)"`
	// 创建时间
	CreateTime time.Time `orm:"auto_now_add;type(datetime);column(ctime);null"`
	// 修改时间
	UpdateTime time.Time `orm:"auto_now;type(datetime);column(mtime);null"`
}

type TableInfo struct {
	// 创建时间
	CreateTime time.Time `orm:"auto_now_add;type(datetime);column(ctime)"`
}

const (
	Inactive = 10
	Active   = 20
)
