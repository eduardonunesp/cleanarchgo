package webserver

import (
	"errors"
	"net/http"

	"github.com/eduardonunesp/cleanarchgo/pkg/service"

	"github.com/labstack/echo/v4"
)

type (
	accountResponse struct {
		ID          string `json:"id" form:"id"`
		Name        string `json:"name" form:"name"`
		Email       string `json:"email" form:"email"`
		CPF         string `json:"cpf" form:"cpf"`
		CarPlate    string `json:"carPlate" form:"carPlate"`
		AccountType string `json:"accountType" form:"accountType"`
	}
)

func (s apiServer) GetAccount(e *echo.Echo) {
	e.POST("/request_ride", func(c echo.Context) error {
		accountID := c.Param("accountID")
		if accountID == "" {
			return c.JSON(
				http.StatusBadRequest,
				raiseWebserverError(errors.New("accountID cannot be empty")),
			)
		}
		result, err := s.getAccount.Execute(service.GetAccountInput{
			ID: accountID,
		})
		serviceErr := new(service.ServiceError)
		if errors.As(err, &serviceErr) {
			return c.NoContent(http.StatusInternalServerError)
		}
		if err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}
		if result == nil {
			return c.NoContent(http.StatusNotFound)
		}
		return c.JSON(http.StatusOK, accountResponse{
			ID:          result.ID,
			Name:        result.Name,
			CPF:         result.CPF,
			Email:       result.Email,
			CarPlate:    result.CarPlate,
			AccountType: result.AccountType,
		})
	})
}
