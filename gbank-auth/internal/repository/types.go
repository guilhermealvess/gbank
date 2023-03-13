package repository

/* type ClientDB interface {
	Insert(model T) error
} */

type Cache interface {
	Set(key string, value string) error

	Get(key string) (string, error)
}
