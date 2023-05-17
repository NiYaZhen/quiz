package controller

import (
	"comment/db"
	"context"
	"net/http"

	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson"
)

// DeleteComment - 刪除評論
func DeleteComment(ctx iris.Context) {
	// 取得路由參數中的 UUID
	uuid := ctx.Params().Get("uuid")

	// 將 comment 資料從 MongoDB 刪除
	filter := bson.M{"uuid": uuid}
	result, err := db.Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		ctx.StatusCode(http.StatusInternalServerError)
		ctx.WriteString(err.Error())
		return
	}

	if result.DeletedCount == 0 {
		// 找不到符合條件的評論
		ctx.StatusCode(http.StatusBadRequest)
		ctx.WriteString("delete fail")
		return
	}

	// 回傳成功訊息
	ctx.StatusCode(http.StatusOK)
	ctx.WriteString("success")
}
