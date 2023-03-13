package entity

type ApplicationRepository interface {
	Save(application *Application) error

	// FindAll() ([]*Application, error)

	// FindById(id uuid.UUID) (*Application, error)
}
