package webserver

import (
	"errors"
	"net/http"

	"github.com/eduardonunesp/cleanarchgo/pkg/domain"
	"github.com/eduardonunesp/cleanarchgo/pkg/service"

	"github.com/labstack/echo/v4"
)

type (
	requestRideRequest struct {
		PassengerID string `json:"passengerId"`
		FromLat     string `json:"fromLat"`
		FromLong    string `json:"fromLong"`
		ToLat       string `json:"toLat"`
		ToLong      string `json:"toLong"`
	}

	requestRideResponse struct {
		RideID string `json:"rideID"`
	}
)

func (s apiServer) requestRideRequest(c echo.Context) error {
	reqBody := new(requestRideRequest)
	if err := c.Bind(reqBody); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			raiseWebserverError(errors.New("invalid input, malformed json")),
		)
	}
	result, err := s.reqRide.Execute(&service.RequestRideParams{
		PassengerID: reqBody.PassengerID,
		FromLat:     reqBody.FromLat,
		FromLong:    reqBody.FromLong,
		ToLat:       reqBody.ToLat,
		ToLong:      reqBody.ToLong,
	})
	serviceErr := new(service.ServiceError)
	if errors.As(err, &serviceErr) {
		return c.JSON(
			http.StatusBadRequest,
			raiseWebserverError(serviceErr),
		)
	}
	domainErr := new(domain.DomainError)
	if errors.As(err, &domainErr) {
		return c.JSON(
			http.StatusBadRequest,
			raiseWebserverError(domainErr),
		)
	}
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, requestRideResponse{
		RideID: result.RideID,
	})
}
