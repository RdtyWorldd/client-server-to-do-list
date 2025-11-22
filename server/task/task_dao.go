package task

type TaskDao interface {
	Create(Task) error
	Read(id int) (Task, error)
	Update(id int, upd_task Task) error
	Delete(id int) error
	ReadAll() []Task
}
