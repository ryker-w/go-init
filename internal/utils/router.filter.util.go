package utils

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter/midware/auth"
	"github.com/ryker-w/go-init/internal/etc"
)

// WithAuth token验证器,
// auth.JwtBasic header预处理
// auth.Forbidden401Handler 无权限时返回401, 返回格式按照参数 auth.ForbiddenOption
func WithAuth(handler func(ctx iris.Context)) []iris.Handler {
	var handlers []iris.Handler
	if etc.Config.Token.Enable {
		handlers = append(handlers, auth.JwtBasic(), auth.Forbidden401Handler(auth.WithJsonResp()))
	}
	handlers = append(handlers, Filter, handler)
	return handlers
}

// Filter 自定义请求拦截器
func Filter(ctx iris.Context) {
	// TODO Something

	// 下一个 handler
	ctx.Next()
}
