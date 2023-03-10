package usecase

import (
	"github.com/google/uuid"
	"github.com/guilhermealvess/gbank/gbank-auth/internal/entity"
)

type CreateApplication struct {
	repository entity.ApplicationRepository
}

func (c *CreateApplication) ExecuteCreateApplication(input *InputCreateApplication) (*OutputCreateApplication, error) {
	application := entity.NewApplication(input.Name, input.Description)
	err := c.repository.Save(application)

	if err != nil {
		return &OutputCreateApplication{}, err
	}

	return &OutputCreateApplication{
		ID:          application.ID,
		Name:        application.Name,
		Description: application.Description,
		PublicKey:   application.PublicKey,
		Status:      application.Status,
	}, nil
}

type InputCreateApplication struct {
	Name        string
	Description string
}

type OutputCreateApplication struct {
	ID          uuid.UUID
	Name        string
	Description string
	PublicKey   string
	Status      string
}
