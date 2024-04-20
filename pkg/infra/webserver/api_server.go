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

	accountResponse struct {
		ID          string `json:"id" form:"id"`
		Name        string `json:"name" form:"name"`
		Email       string `json:"email" form:"email"`
		CPF         string `json:"cpf" form:"cpf"`
		CarPlate    string `json:"carPlate" form:"carPlate"`
		IsPassenger bool   `json:"isPassenger" form:"isPassenger"`
		IsDriver    bool   `json:"isDriver" form:"isDriver"`
	}
)

type apiServer struct {
	getAccount *service.GetAccount
	signup     *service.Signup
}

func NewAPIServer(getAccount *service.GetAccount, signup *service.Signup) *apiServer {
	return &apiServer{getAccount, signup}
}

func (s apiServer) signupRequest(c echo.Context) error {
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
}

func (s apiServer) getAccountRequest(c echo.Context) error {
	accountID := c.Param("accountID")
	if accountID == "" {
		return c.JSON(
			http.StatusBadRequest,
			raiseWebserverError(errors.New("accountID cannot be empty")),
		)
	}
	result, err := s.getAccount.Execute(&service.GetAccountParams{
		ID: accountID,
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
	return c.JSON(http.StatusOK, accountResponse{
		ID:          result.ID,
		Name:        result.Name,
		CPF:         result.CPF,
		Email:       result.Email,
		CarPlate:    result.CarPlate,
		IsPassenger: result.IsPassenger,
		IsDriver:    result.IsDriver,
	})
}

func (s apiServer) Listen(host string) {
	e := echo.New()
	e.POST("/signup", s.signupRequest)
	e.GET("/accounts/:accountID", s.getAccountRequest)
	log.Printf("Server listen at %s", host)
	e.Start(host)
}
