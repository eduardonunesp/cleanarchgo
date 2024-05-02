package webserver

import (
	"log"

	"github.com/eduardonunesp/cleanarchgo/pkg/service"

	"github.com/labstack/echo/v4"
)

type APIServer struct {
	getAccount *service.GetAccount
	getRide    *service.GetRide
	reqRide    *service.RequestRide
	signup     *service.Signup
}

func NewAPIServer(
	getAccount *service.GetAccount,
	signup *service.Signup,
	getRide *service.GetRide,
	reqRide *service.RequestRide,
) *APIServer {
	return &APIServer{getAccount, getRide, reqRide, signup}
}

func (s APIServer) Listen(host string) {
	e := echo.New()
	s.Signup(e)
	s.GetAccount(e)
	s.RequestRide(e)
	s.GetRide(e)
	log.Printf("Server listen at %s", host)
	if err := e.Start(host); err != nil {
		log.Fatal(err)
	}
}
