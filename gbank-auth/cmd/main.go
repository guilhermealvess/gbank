package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/guilhermealvess/gbank/gbank-auth/internal/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
)

func main() {
	db, err := makeClientPostgresql()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	startHttpServer(db)

}

func startHttpServer(db *sql.DB) {
	server := echo.New()
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())

	server.GET("/ping", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "PONG")
	})

	applicationController := controller.NewApplicationController(db)
	server.POST("/application", applicationController.CreateApplication)

	port := 3000
	server.Logger.Fatal(server.Start(fmt.Sprintf(":%d", port)))
}

func makeClientPostgresql() (*sql.DB, error) {
	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "2023"
	dbname := "gbank_auth_db"

	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	return sql.Open("postgres", dataSourceName)

}

/* func sqlSelect(db *sql.DB) ([]*controller.ApplicationDTO, error) {
	var arr []*controller.ApplicationDTO
	sql := "SELECT id, name, description, status FROM application"

	sqlStatement, err := db.Query(sql)

	if err != nil {
		return arr, err
	}

	for sqlStatement.Next() {
		var app controller.ApplicationDTO
		err = sqlStatement.Scan(&app.ID, &app.Name, &app.Description, &app.Status)

		if err != nil {
			return arr, err
		}

		arr = append(arr, &app)
	}

	return arr, nil
}
*/
