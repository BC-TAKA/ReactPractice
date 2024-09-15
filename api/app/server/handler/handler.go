package handler

import "github.com/ReactPractice/api/domain/service"

type Handler interface {
	AddressListHandler
}

type handler struct {
	AddressListHandler
}

func NewHandler(val service.AddressListService) Handler {
	return &handler{
		AddressListHandler: NewAddressListHandler(val),
	}
}
