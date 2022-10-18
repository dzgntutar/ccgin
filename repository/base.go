package repository

type IRepository[T any] interface {
	GetAll() ([]T, error)
	Create(t T) error
	GetById(id string) (T, error)
	Delete(id string) error
}
