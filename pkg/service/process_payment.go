package service

import (
	"github.com/eduardonunesp/cleanarchgo/pkg/infra/gateway"
	"github.com/eduardonunesp/cleanarchgo/pkg/infra/repository"
)

type ProcessPaymentParams struct {
	RideID          string
	CreditCardToken string
	Amount          string
}

type ProcessPayment struct {
	rideRepo  repository.RideRepository
	ccGateway gateway.CreditCardGW
}

func NewProcessPayment(rideRepo repository.RideRepository, ccGateway gateway.CreditCardGW) *ProcessPayment {
	return &ProcessPayment{rideRepo, ccGateway}
}

func (s ProcessPayment) Execute(input ProcessPaymentParams) error {
	ride, err := s.rideRepo.GetRideByID(input.RideID)
	if err != nil {
		return err
	}
	_ = ride
	return s.ccGateway.ProcessPayment(input.CreditCardToken, input.Amount)
}
