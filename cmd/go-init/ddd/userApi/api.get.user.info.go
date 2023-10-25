package userApi

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/ryker-w/go-init/internal/db/model"
	"github.com/ryker-w/go-init/internal/utils"
)

type userInfo struct {
	Id         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Status     int    `json:"status,omitempty"`
	CreateTime string `json:"createTime,omitempty"`
}

type respUserInfo struct {
	app.Response
	Data userInfo `json:"data"`
}

func getUserInfo(ctx iris.Context) {
	var resp respUserInfo
	// 用户id
	id := ctx.Params().GetIntDefault("id", 0)
	if id == 0 {
		log.Info("id is 0")
		resp.Code = tool.RespCodeNotFound
		tool.ResponseJSON(ctx, resp)
		return
	}
	// 查询
	var user model.UserInfo
	user.Id = id
	err := app.GetOrm().Context.Read(&user)
	if err != nil {
		log.Info(err)
		resp.Code = tool.RespCodeNotFound
		tool.ResponseJSON(ctx, resp)
		return
	}
	utils.SimpleCopyProperties2(user, &resp.Data)
	// 返回结果
	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}
