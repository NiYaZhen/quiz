package router

import (
	"comment/controller"

	"github.com/kataras/iris/v12"
)

// NewRouter - 新建路由
func NewRouter(app *iris.Application) {
	api := app.Party("/comment")
	{
		api.Post("/", controller.CreateComment)
		api.Get("/{uuid}", controller.GetComment)
		api.Put("/{uuid}", controller.UpdateComment)
		api.Delete("/{uuid}", controller.DeleteComment)
	}

}
