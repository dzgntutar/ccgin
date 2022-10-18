package service

type IService[T any] interface {
	GetAll() ([]T, error)
	Create(t T) error
	GetById(id string) (T, error)
	Delete(id string) error
}
