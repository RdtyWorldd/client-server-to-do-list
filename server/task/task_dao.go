package task

type TaskDao interface {
	Create(Task) error
	Read(id struct{ C_id, T_id int }) (Task, error)
	Update(id struct{ C_id, T_id int }, upd_task Task) error
	Delete(id struct{ C_id, T_id int }) error
	ReadAll() []Task
}
