package dao

import (
	"github.com/RdtyWorldd/client-server-to-do-list/server/client"
	"github.com/RdtyWorldd/client-server-to-do-list/server/task"
)

type ClientDao interface {
	//CRUD operations
	Create(client.Client) error
	Read(id int) (client.Client, error)
	Update(id int, upd_client client.Client) error
	Delete(id int) error
	ReadAll() []client.Client

	//Client specifice
	// нужно ли тут использовать указатель?
	AddTask(c_id int, task task.Task) (client.Client, error)
	DeleteTask(c_id int, t_id int) (client.Client, error)
	UpdateTask(c_id int, task task.Task) (client.Client, error)
}
