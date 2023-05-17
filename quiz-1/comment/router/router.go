package router

import (
	"comment/controller"

	"github.com/kataras/iris/v12"
)

// NewRouter - 新建路由
func NewRouter(app *iris.Application) {
	// 評論
	api := app.Party("/comment")
	{
		// 創建評論
		api.Post("/", controller.CreateComment)
		// 取得評論
		api.Get("/{uuid}", controller.GetComment)
		// 更新評論
		api.Put("/{uuid}", controller.UpdateComment)
		// 刪除評論
		api.Delete("/{uuid}", controller.DeleteComment)
	}

}
