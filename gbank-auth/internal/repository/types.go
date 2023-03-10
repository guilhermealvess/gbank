package repository

type SqlDB interface {
	Insert(model any) error
}

type Cache interface {
	Set(key string, value string) error

	Get(key string) (string, error)
}
