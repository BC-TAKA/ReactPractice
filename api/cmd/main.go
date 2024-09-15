package main

import (
	"log"

	"github.com/ReactPractice/api/app/server"
	"github.com/ReactPractice/api/app/server/handler"
	"github.com/ReactPractice/api/domain/service"
	"github.com/ReactPractice/api/infra"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewHandler() handler.Handler {
	repo := infra.NewAddressList()
	service := service.NewAddressListService(repo)

	// 各層の構造体を渡す
	h := handler.NewHandler(service)

	return h
}

func main() {
	h := NewHandler()
	// インスタンス作成
	e := echo.New()

	// ミドルウェア設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	server := server.New(e).Handlers(h)

	// // エンドポイント
	// e.GET("/", hello)

	// ポート番号1323で起動
	// e.Logger.Fatal(e.Start(":1323"))

	log.Fatal(server.Start())
}

// // 疎通確認用ハンドラー
// func hello(c echo.Context) error {
// 	return c.String(http.StatusOK, "Welcome to Golang!")
// }
