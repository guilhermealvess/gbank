package entity

import (
	"time"

	"github.com/google/uuid"
)

type Application struct {
	ID          uuid.UUID
	Name        string
	Description string
	PrivateKey  string
	PublicKey   string
	Status      string
	CreatedAt   time.Time
}

func NewApplication(name string, description string) *Application {
	id, _ := uuid.NewUUID()
	privateKey := ""
	publicKey := ""

	return &Application{
		ID:          id,
		Name:        name,
		Description: description,
		PrivateKey:  privateKey,
		PublicKey:   publicKey,
		CreatedAt:   time.Now(),
	}
}
