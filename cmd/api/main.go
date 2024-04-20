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
	mailerGW := gateway.NewMailerGatewayMemory()
	accountServ := service.NewGetAccount(accountRepo)
	signupServ := service.NewSignup(accountRepo, mailerGW)
	api := webserver.NewAPIServer(accountServ, signupServ)
	api.Listen(":3000")
}
