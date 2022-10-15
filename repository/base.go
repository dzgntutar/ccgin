package repository

type IRepository[T any] interface {
	GetAll() ([]T, error)
	Create(t T)
}
