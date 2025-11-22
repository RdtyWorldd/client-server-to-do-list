package dao

import "github.com/RdtyWorldd/client-server-to-do-list/server/client"

type ClientDao interface {
	Create(client.Client) error
	Read(id int) (client.Client, error)
	Update(id int, upd_client client.Client) error
	Delete(id int) error
	ReadAll() []client.Client
}
