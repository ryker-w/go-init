package model

type User struct {
	Pk
	Name string `orm:"column(name);null"`
	TableChangeInfo
}
