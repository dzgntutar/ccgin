package service

type IService[T any] interface {
	GetAll() ([]T, error)
	Create()
}
