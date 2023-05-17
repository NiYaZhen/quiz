package controller

import (
	"comment/db"
	"comment/model"
	"context"
	"net/http"

	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson"
)

// UpdateComment - 更新評論
func UpdateComment(ctx iris.Context) {
	// 取得路由參數中的 UUID
	uuid := ctx.Params().Get("uuid")

	// 解析請求資料
	var comment model.Comment
	comment.UUID = uuid

	err := ctx.ReadJSON(&comment)
	if err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}

	// 將 comment 資料更新到 MongoDB
	filter := bson.M{"uuid": uuid}
	update := bson.M{"$set": comment}
	result, err := db.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		ctx.StatusCode(http.StatusInternalServerError)
		ctx.WriteString(err.Error())
		return
	}

	if result.MatchedCount == 0 {
		// 找不到符合條件的評論
		ctx.StatusCode(http.StatusBadRequest)
		ctx.WriteString("modify fail")
		return
	}

	// 回傳更新後的 comment 資料
	ctx.StatusCode(http.StatusOK)
	ctx.JSON(comment)
}
