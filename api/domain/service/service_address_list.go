package service

import (
	"github.com/ReactPractice/api/domain/repo"
)

type AddressListService interface {
	// service層での関数をここに追加していく
	AddressListFunc() string
	EditAddress()
}

type addressListService struct {
	// repo.AddressListの宣言内容に依存する
	addressListRepo repo.AddressList
}

func NewAddressListService(val repo.AddressList) AddressListService {
	return &addressListService{
		addressListRepo: val,
	}
}

// service試験的関数
func (service *addressListService) AddressListFunc() string {
	resp := "service pass"

	return resp
}

// 住所録登録内容編集
func (service *addressListService) EditAddress() {

}
