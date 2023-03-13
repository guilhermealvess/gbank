package repository

import (
	"database/sql"
	"fmt"

	"github.com/guilhermealvess/gbank/gbank-auth/internal/entity"
)

type ApplicationRepository struct {
	entity.ApplicationRepository
	db *sql.DB
	//cache Cache
}

func NewApplicationRepository(clientDB *sql.DB /*, cache Cache*/) *ApplicationRepository {
	return &ApplicationRepository{
		db: clientDB,
	}
}

func (a *ApplicationRepository) Save(application *entity.Application) error {
	sqlStatement := fmt.Sprintf("INSERT INTO application VALUES ($1,$2,$3,$4,$5,$6)")

	insert, err := a.db.Prepare(sqlStatement)
	if err != nil {
		return err
	}

	result, err := insert.Exec(application.ID, application.Name, application.Description, application.Status, application.CreatedAt, application.UpdatedAt)
	if err != nil {
		return err
	}

	affect, err := result.RowsAffected()
	if err != nil {
		return err
	}

	println(affect)
	return nil
}
