package webserver

import (
	"errors"
	"net/http"

	"github.com/eduardonunesp/cleanarchgo/pkg/service"

	"github.com/labstack/echo/v4"
)

type (
	rideResponse struct {
		ID             string `json:"id" form:"id"`
		PassengerID    string `json:"passenger_id" form:"passenger_id"`
		FromLat        string `json:"from_lat" form:"from_lat"`
		FromLong       string `json:"from_long" form:"from_long"`
		ToLat          string `json:"to_lat" form:"to_lat"`
		ToLong         string `json:"to_long" form:"to_long"`
		Status         string `json:"status" form:"status"`
		PassengerName  string `json:"passenger_name" form:"passenger_name"`
		PassengerEmail string `json:"passenger_email" form:"passenger_email"`
	}
)

func (s apiServer) GetRide(e *echo.Echo) {
	e.GET("/rides/:rideID", func(c echo.Context) error {
		rideID := c.Param("rideID")
		if rideID == "" {
			return c.JSON(
				http.StatusBadRequest,
				raiseWebserverError(errors.New("accountID cannot be empty")),
			)
		}
		result, err := s.getRide.Execute(&service.GetRideParams{
			RideID: rideID,
		})
		serviceErr := new(service.ServiceError)
		if errors.As(err, &serviceErr) {
			return c.NoContent(http.StatusNotFound)
		}
		if err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}
		if result == nil {
			return c.NoContent(http.StatusNotFound)
		}
		return c.JSON(http.StatusOK, rideResponse{
			ID:             result.ID,
			PassengerID:    result.PassengerID,
			FromLat:        result.FromLat,
			FromLong:       result.FromLong,
			ToLat:          result.ToLat,
			ToLong:         result.ToLong,
			Status:         result.Status,
			PassengerName:  result.PassengerName,
			PassengerEmail: result.PassengerEmail,
		})
	})
}
