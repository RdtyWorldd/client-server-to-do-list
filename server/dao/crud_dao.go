package dao

type CrudDao[T any] interface {
	Create(T) error
	Read(id int) (T, error)
	Update(id int, upd_task T) error
	Delete(id int) error
	ReadAll() []T
}
