package service

import (
	"github.com/eduardonunesp/cleanarchgo/pkg/domain"
	"github.com/eduardonunesp/cleanarchgo/pkg/infra/repository"
)

type UpdatePositionParams struct {
	RideID string
	Lat    string
	Long   string
}

type UpdatePosition struct {
	posRepo repository.PositionRepository
}

func NewUpdatePosition(posRepo repository.PositionRepository) *UpdatePosition {
	return &UpdatePosition{posRepo}
}

func (s UpdatePosition) Execute(input UpdatePositionParams) error {
	position, err := domain.BuildPosition(
		domain.PositionWithRideID(input.RideID),
		domain.PositionWithCoord(input.Lat, input.Long),
	)
	if err != nil {
		return err
	}
	return s.posRepo.SavePosition(position)
}
