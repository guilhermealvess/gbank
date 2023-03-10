package database

import "github.com/guilhermealvess/gbank/gbank-auth/internal/repository"

type DatabaseEntities interface {
	repository.ApplicationModel
}
