package repo

type AddressList interface {
	// service層とinfra層を繋ぐ役割
	// 使用する関数はここに宣言
	AddressListFunc() string
	EditAddress()
}
