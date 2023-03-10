package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/guilhermealvess/gbank/gbank-auth/internal/controller"
	"github.com/guilhermealvess/gbank/gbank-auth/internal/database"
	"github.com/guilhermealvess/gbank/gbank-auth/internal/repository"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	//go startHttpServer()

	testDatabase()
}

func startHttpServer() {
	port := 3000

	server := echo.New()
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())

	server.GET("/ping", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "PONG")
	})

	server.POST("/application", controller.CreateApplication)
	server.GET("/application", controller.CreateApplication)
	server.GET("/application/:id", controller.GetApplicationById)
	server.PATCH("/application", controller.CreateApplication)

	server.Logger.Fatal(server.Start(fmt.Sprintf(":%d", port)))
}

func testDatabase() {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=postgres password=2023 dbname=gbank_auth port=5432 sslmode=disable TimeZone=America/Sao_Paulo host=localhost",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db_sql := database.NewSqlDB[repository.ApplicationModel](db)
	println(db_sql)

	id, _ := uuid.NewRandom()
	now := time.Now()

	model := &repository.ApplicationModel{
		ID:          id,
		Name:        "Guilherme",
		Description: "Alves",
		Status:      "new",
		PrivateKey:  "",
		PublicKey:   "",
		CreatedAt:   now,
	}

	// Mock
	type Product struct {
		gorm.Model
		ID    uint `gorm:"primaryKey;default:auto_random()"`
		Code  string
		Price uint
	}

	db.Create(&Product{})

	db_sql.Insert(*model)
}
