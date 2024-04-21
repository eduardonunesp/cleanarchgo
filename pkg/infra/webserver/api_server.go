package webserver

import (
	"log"

	"github.com/eduardonunesp/cleanarchgo/pkg/service"

	"github.com/labstack/echo/v4"
)

type apiServer struct {
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
) *apiServer {
	return &apiServer{getAccount, getRide, reqRide, signup}
}

func (s apiServer) Listen(host string) {
	e := echo.New()
	e.POST("/signup", s.signupRequest)
	e.POST("/request_ride", s.requestRideRequest)
	e.GET("/rides/:rideID", s.getRideRequest)
	e.GET("/accounts/:accountID", s.getAccountRequest)
	log.Printf("Server listen at %s", host)
	e.Start(host)
}
