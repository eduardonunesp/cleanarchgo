package main

import (
	"github.com/eduardonunesp/cleanarchgo/pkg/infra/gateway"
	"github.com/eduardonunesp/cleanarchgo/pkg/infra/repository"
	"github.com/eduardonunesp/cleanarchgo/pkg/infra/webserver"
	"github.com/eduardonunesp/cleanarchgo/pkg/service"
)

func main() {
	connStr := "postgres://postgres:123456@localhost:5432/postgres?sslmode=disable"
	accountRepo := repository.NewAccountRepositoryPG(connStr)
	rideRepo := repository.NewRideRepositoryPG(connStr)
	mailerGW := gateway.NewMailerGatewayMemory()
	accountServ := service.NewGetAccount(accountRepo)
	signupServ := service.NewSignup(accountRepo, mailerGW)
	reqRideServ := service.NewRequestRide(rideRepo, accountRepo)
	getRideServ := service.NewGetRide(rideRepo, accountRepo)
	api := webserver.NewAPIServer(accountServ, signupServ, getRideServ, reqRideServ)
	api.Listen(":3000")
}
