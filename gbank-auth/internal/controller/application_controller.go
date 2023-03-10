package controller

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/guilhermealvess/gbank/gbank-auth/internal/database"
	"github.com/guilhermealvess/gbank/gbank-auth/internal/repository"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type ApplicationController struct {
	//db *gorm.DB
	db_sql *database.SqlDB[repository.ApplicationModel]
}

func NewApplicationController(db *gorm.DB) *ApplicationController {
	db_sql := database.NewSqlDB[repository.ApplicationModel](db)
	return &ApplicationController{
		db_sql: db_sql,
	}
}

func CreateApplication(c echo.Context) error {
	application := new(Application)

	if err := c.Bind(application); err != nil {
		return err
	}

	return c.JSON(201, application)
}

func GetApplicationById(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		return err
	}

	return c.String(http.StatusOK, id.String())
}

type Application struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
}
