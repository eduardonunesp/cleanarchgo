package main

import (
	"fmt"
	"os"

	"github.com/eduardonunesp/cleanarchgo/internal/infra/gateway"
	"github.com/eduardonunesp/cleanarchgo/internal/infra/repository"
	"github.com/eduardonunesp/cleanarchgo/internal/infra/webserver"
	"github.com/eduardonunesp/cleanarchgo/internal/service"
)

func main() {
	connStr := fmt.Sprintf("postgres://%s", os.Getenv("PG_CONN_STR"))
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
