package main

import (
	"comment/db"
	"comment/router"

	"github.com/kataras/iris/v12"
)

func init() {
	// 讀取env檔案
	db.LoadTheEnv()
	// 創建資料庫實例
	db.CreateDBInstance()
}

func main() {

	// 建立 Iris 應用
	app := iris.New()

	router.NewRouter(app)

	// 啟動 Iris 應用
	app.Run(iris.Addr(":8080"))
}
