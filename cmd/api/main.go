package main

import (
	"fmt"
	"os"

	"github.com/eduardonunesp/cleanarchgo/pkg/infra/gateway"
	"github.com/eduardonunesp/cleanarchgo/pkg/infra/repository"
	"github.com/eduardonunesp/cleanarchgo/pkg/infra/webserver"
	"github.com/eduardonunesp/cleanarchgo/pkg/service"
)

func main() {
	var connStr string
	if str, ok := os.LookupEnv("PG_CONN_STR"); ok {
		connStr = str
	}
	connStr = fmt.Sprintf("postgres://%s", connStr)
	fmt.Println(connStr)
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
