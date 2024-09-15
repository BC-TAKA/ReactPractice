package handler

import (
	"net/http"

	"github.com/ReactPractice/api/domain/service"
	"github.com/labstack/echo/v4"
)

type AddressListHandler interface {
	// 関数を宣言
	AddressListFunc(c echo.Context) error
	EditAddress(e echo.Context) error
}

type addressListHandler struct {
	addressList service.AddressListService
}

func NewAddressListHandler(val service.AddressListService) AddressListHandler {
	return &addressListHandler{
		addressList: val,
	}
}

// 試験的関数
func (handle *addressListHandler) AddressListFunc(c echo.Context) error {
	// リクエスト内容の処理が必要

	const resp = "handler dummy func"
	// service層からの結果をrespに格納して返す
	return c.JSON(http.StatusOK, resp)
}

// 住所録登録内容編集
func (handle *addressListHandler) EditAddress(c echo.Context) error {
	const resp = "handler edit address"
	return c.JSON(http.StatusOK, resp)
}
