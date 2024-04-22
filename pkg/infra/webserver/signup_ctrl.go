package webserver

import (
	"errors"
	"log"
	"net/http"

	"github.com/eduardonunesp/cleanarchgo/pkg/domain"
	"github.com/eduardonunesp/cleanarchgo/pkg/service"

	"github.com/labstack/echo/v4"
)

type (
	signupRequest struct {
		Name        string `json:"name"`
		Email       string `json:"email"`
		CPF         string `json:"cpf"`
		CarPlate    string `json:"carPlate"`
		IsPassenger bool   `json:"isPassenger"`
		IsDriver    bool   `json:"isDriver"`
	}

	signupResponse struct {
		AccountUID string `json:"acountId"`
	}
)

func (s apiServer) Signup(e *echo.Echo) {
	e.POST("/signup", func(c echo.Context) error {
		reqBody := new(signupRequest)
		if err := c.Bind(reqBody); err != nil {
			return c.JSON(
				http.StatusBadRequest,
				raiseWebserverError(errors.New("invalid input, malformed json")),
			)
		}
		result, err := s.signup.Execute(&service.SignupParams{
			Name:        reqBody.Name,
			Email:       reqBody.Email,
			CPF:         reqBody.CPF,
			CarPlate:    reqBody.CarPlate,
			IsPassenger: reqBody.IsPassenger,
			IsDriver:    reqBody.IsDriver,
		})
		signupErr := new(service.ServiceError)
		if errors.As(err, &signupErr) {
			return c.JSON(
				http.StatusBadRequest,
				raiseWebserverError(signupErr),
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
			log.Println("Error: ", err.Error())
			return c.NoContent(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, signupResponse{
			AccountUID: result.AccountUID,
		})
	})
}
