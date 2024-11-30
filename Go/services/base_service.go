package services

type BaseService[T any] interface {
	Create(data T) (T, error)
	Read(param interface{}) interface{}
	Update(id string, data interface{}) (T, error)
	Delete(id string) (T, error)
}
