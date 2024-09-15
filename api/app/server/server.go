package server

import (
	"github.com/ReactPractice/api/app/server/handler"
	"github.com/labstack/echo/v4"
)

type Server struct {
	echo *echo.Echo
}

func New(e *echo.Echo) *Server {
	return &Server{
		echo: e,
	}
}

func (s *Server) Handlers(h handler.Handler) *Server {
	// curl localhost:1234/api/test
	s.echo.GET("api/test", h.AddressListFunc)

	return s
}

func (s *Server) Start() error {
	port := "1234"
	return s.echo.Start(":" + port)
}
