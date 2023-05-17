package controller

import (
	"comment/db"
	"comment/model"
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
)

// CreateComment - 建立評論
func CreateComment(ctx iris.Context) {
	// 解析請求資料
	var comment model.Comment
	err := ctx.ReadJSON(&comment)
	if err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}

	// 生成一個新的 UUID
	uuid := uuid.New()

	// 設定 comment 資料的 UUID
	comment.UUID = uuid.String()

	// 將 comment 資料儲存到 MongoDB
	_, err = db.Collection.InsertOne(context.Background(), comment)
	if err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.WriteString(err.Error())
		ctx.WriteString("create comment fail")
		return
	}

	// 回傳新建立的 Comment 資料
	ctx.StatusCode(http.StatusOK)
	ctx.JSON(comment)
}
