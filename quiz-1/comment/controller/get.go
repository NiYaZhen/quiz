package controller

import (
	"comment/db"
	"comment/model"
	"context"
	"net/http"

	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson"
)

// GetComment - 取得單一評論
func GetComment(ctx iris.Context) {
	var comment model.Comment
	// 取得路由參數 "uuid"的值
	uuid := ctx.Params().Get("uuid")

	// 查詢特定評論資料
	filter := bson.M{"uuid": uuid}
	result := db.Collection.FindOne(context.Background(), filter)
	if result.Err() != nil {
		ctx.StatusCode(http.StatusInternalServerError)
		ctx.WriteString(result.Err().Error())
		return
	}

	// 解碼評論資料
	err := result.Decode(&comment)
	if err != nil {
		ctx.StatusCode(http.StatusInternalServerError)
		ctx.WriteString(err.Error())
		ctx.WriteString("comment uuid not found")
		return
	}

	// 回傳單個評論資料
	ctx.JSON(comment)
}
