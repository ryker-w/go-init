package ddd

import (
	"github.com/kataras/iris/v12"
	"github.com/ryker-w/go-init/cmd/go-init/ddd/userApi"
)

// web资源入口
func Router(app *iris.Application) {
	root := app.Party("/api") // 根路径
	router(root)
}

// 二级资源路径
func router(root iris.Party) {
	userApi.Route(root.Party("/users")) //rest 接口规范
	// userApi.Route(root.Party("/getUserList")) //RPC 接口规范
}
