package userApi

import (
	"github.com/kataras/iris/v12"
)

// 三级资源路径
func Route(p iris.Party) {
	p.Get("/", nil)             // 获取用户列表 == GET /api/users == GET /api/user/list
	p.Get("/{id}", getUserInfo) // 获取用户信息 == GET /api/users/1 == GET /api/user/getUser?id=1
	p.Post("/", createUser)     // 新增用户 == POST /api/users == POST /api/user/create
	// p.Post("/", utils.WithAuth(createUser)...) // 新增用户(with拦截器) == POST /api/users == POST /api/user/create
	p.Put("/{id}", nil) // 修改用户信息 == PUT /api/users/1 == PUT /api/user/update?id=1
}
