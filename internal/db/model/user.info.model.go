package model

import "github.com/lishimeng/app-starter"

type UserInfo struct {
	app.Pk
	Name string `orm:"column(name);null"`
	app.TableChangeInfo
}
