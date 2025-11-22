package dao

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"sort"

	"github.com/RdtyWorldd/client-server-to-do-list/server/client"
	"github.com/RdtyWorldd/client-server-to-do-list/server/task"
)

type FileClientDao struct {
	path_client string
	clientMap   map[int]client.Client
	task_dao    task.TaskDao
}

func NewFileClientDao(path_client string, task_dao task.TaskDao) *FileClientDao {
	file, err := os.OpenFile(path_client, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	dao := FileClientDao{path_client, make(map[int]client.Client), task_dao}
	if len(data) != 0 {
		var clients []client.Client
		err = json.Unmarshal(data, &clients)
		if err != nil {
			panic(err)
		}
		task_list := dao.task_dao.ReadAll()
		for i, value := range clients {
			for _, task := range task_list { //need optimization
				if task.OwnerID == value.GetID() {
					value.AddTask(task)
				}
			}
			dao.clientMap[clients[i].GetID()] = value
		}
	}
	return &dao
}

// question
// нужно ли проверять индекс или доверяться обработчикам комманд
func (dao *FileClientDao) Create(client client.Client) error {
	file, err := os.OpenFile(dao.path_client, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	client.SetId(len(dao.clientMap))
	file.Seek(-1, 2)
	task_json, err := json.Marshal(&client)
	if err != nil {
		return err
	}
	write_data := string(task_json) + "]"

	if len(dao.clientMap) != 0 {
		write_data = "," + write_data
	} else {
		file.Seek(0, 0)
		write_data = "[" + write_data
	}
	_, err = io.WriteString(file, write_data)
	if err != nil {
		return err
	}
	dao.clientMap[client.GetID()] = client
	return nil
}

func (dao *FileClientDao) Read(id int) (client.Client, error) {
	if id < 0 || id > len(dao.clientMap) {
		return client.Client{}, errors.New("index out of range") //пусть пока повисит пустая таска
	}
	if value, ok := dao.clientMap[id]; ok {
		return value, nil
	} else {
		return client.Client{}, errors.New("index out of range")
	}
}

func (dao *FileClientDao) ReadAll() []client.Client {
	res := make([]client.Client, 0, len(dao.clientMap))
	for _, value := range dao.clientMap {
		res = append(res, value)
	}

	sort.Slice(res, func(i int, j int) bool { return res[i].GetID() < res[j].GetID() })
	return res
}

func (dao *FileClientDao) Update(id int, upd client.Client) error {
	if id < 0 || id > len(dao.clientMap) {
		return errors.New("index out of range")
	}
	dao.clientMap[id] = upd
	file, err := os.OpenFile(dao.path_client, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := dao.marshal()
	if err != nil {
		return err
	}
	_, err = io.Writer.Write(file, data)
	if err != nil {
		return err
	}
	file.Truncate(int64(len(data)))
	return nil
}

func (dao *FileClientDao) Delete(id int) error {
	if id < 0 || id > len(dao.clientMap) {
		return errors.New("index out of range") //пусть пока повисит пустая таска
	}
	dao.delete_client(id)
	file, err := os.OpenFile(dao.path_client, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := dao.marshal()
	if err != nil {
		return err
	}
	_, err = io.Writer.Write(file, data)
	if err != nil {
		return err
	}
	file.Truncate(int64(len(data)))
	return nil
}

func (dao FileClientDao) marshal() ([]byte, error) {
	task_list := make([]client.Client, 0, len(dao.clientMap))
	for _, value := range dao.clientMap {
		task_list = append(task_list, value)
	}
	return json.Marshal(task_list)
}

func (dao *FileClientDao) delete_client(id int) {
	delete(dao.clientMap, id)
}
