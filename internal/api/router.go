package api

import (
	"github.com/kataras/iris/v12"
)

func Route(app *iris.Application) {
	root := app.Party("/api")
	router(root)
	return
}
func router(root iris.Party) {
	User(root.Party("/user"))
}
func User(p iris.Party) {
	p.Get("/",nil)
	p.Get("/{id}",nil)

	p.Post("/",nil)
	p.Put("/{id}",nil)
}
