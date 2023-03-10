package repository

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/guilhermealvess/gbank/gbank-auth/internal/entity"
)

type ApplicationRepository struct {
	entity.ApplicationRepository
	sqlDB SqlDB
	cache Cache
}

func NewApplicationRepository(sqlDB SqlDB, cache Cache) *ApplicationRepository {
	return &ApplicationRepository{
		sqlDB: sqlDB,
		cache: cache,
	}
}

func (a *ApplicationRepository) Save(application *entity.Application) error {
	applicationModel := entityToModel(application)
	err := a.sqlDB.Insert(&applicationModel)

	if err != nil {
		return err
	}

	key := fmt.Sprintf("application:%s", application.ID.String())
	b, err := json.Marshal(applicationModel)
	if err != nil {
		return err
	}
	value := string(b)

	err = a.cache.Set(key, value)

	return nil
}

type ApplicationModel struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	PrivateKey  string    `json:"privateKey"`
	PublicKey   string    `json:"publicKey"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
}

func entityToModel(a *entity.Application) *ApplicationModel {
	return &ApplicationModel{
		ID:          a.ID,
		Name:        a.Name,
		Description: a.Description,
		PrivateKey:  a.PrivateKey,
		PublicKey:   a.PublicKey,
		Status:      a.Status,
		CreatedAt:   a.CreatedAt,
	}
}

func modelToEntity(a *ApplicationModel) *entity.Application {
	return &entity.Application{
		ID:          a.ID,
		Name:        a.Name,
		Description: a.Description,
		PrivateKey:  a.PrivateKey,
		PublicKey:   a.PublicKey,
		Status:      a.Status,
		CreatedAt:   a.CreatedAt,
	}
}
