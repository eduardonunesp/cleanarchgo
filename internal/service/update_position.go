package service

import (
	"github.com/eduardonunesp/cleanarchgo/internal/domain"
	"github.com/eduardonunesp/cleanarchgo/internal/infra/repository"
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
	position, err := domain.PositionCreate(domain.PositionCreateParams{
		RideID:         input.RideID,
		CoordinateLat:  input.Lat,
		CoordinateLong: input.Long,
	})
	if err != nil {
		return err
	}
	return s.posRepo.SavePosition(position)
}
