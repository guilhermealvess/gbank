package controller

import (
	"database/sql"
	"net/http"

	"github.com/google/uuid"
	"github.com/guilhermealvess/gbank/gbank-auth/internal/repository"
	"github.com/guilhermealvess/gbank/gbank-auth/internal/usecase"
	"github.com/labstack/echo"
)

type ApplicationController struct {
	clientDB *sql.DB
}

func NewApplicationController(db *sql.DB) *ApplicationController {
	return &ApplicationController{
		clientDB: db,
	}
}

func (a *ApplicationController) CreateApplication(c echo.Context) error {
	applicationDto := new(ApplicationDTO)

	if err := c.Bind(applicationDto); err != nil {
		return err
	}

	applicationRepository := repository.NewApplicationRepository(a.clientDB)
	u := usecase.NewCreateApplication(applicationRepository)
	input := &usecase.InputCreateApplication{
		Name:        applicationDto.Name,
		Description: applicationDto.Description,
	}

	output, err := u.ExecuteCreateApplication(input)
	if err != nil {
		return c.String(http.StatusBadGateway, err.Error())
	}

	return c.JSON(201, &ApplicationDTO{
		ID:          output.ID,
		Name:        output.Name,
		Description: output.Description,
		Status:      output.Status,
	})
}

func GetApplicationById(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		return err
	}

	return c.String(http.StatusOK, id.String())
}
