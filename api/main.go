package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// インスタンス作成
	e := echo.New()

	// ミドルウェア設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// エンドポイント
	e.GET("/", hello)

	// ポート番号1323で起動
	e.Logger.Fatal(e.Start(":1323"))
}

// 疎通確認用ハンドラー
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to Golang!")
}
