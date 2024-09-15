package infra

import "github.com/ReactPractice/api/domain/repo"

type addressListInfra struct{}

func NewAddressList() repo.AddressList {
	return &addressListInfra{}
}

// infra試験的関数
func (infra *addressListInfra) AddressListFunc() string {
	resp := "infra pass"

	return resp
}

// 住所録登録内容編集
func (infra *addressListInfra) EditAddress() {

}
