package userApi

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
)

func createUser(ctx iris.Context) {
	// var req interface{}
	var resp app.Response
	// TODO Something

	// response
	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}
