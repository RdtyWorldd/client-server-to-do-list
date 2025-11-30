package task

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"sort"
)

type FileTaskDao struct {
	path    string
	taskMap map[struct{ C_id, T_id int }]Task
}

func NewFileTaskDao(path string) *FileTaskDao {
	file, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	dao := FileTaskDao{path, make(map[struct{ C_id, T_id int }]Task)}
	if len(data) != 0 {
		var tasks []Task
		err = json.Unmarshal(data, &tasks)
		if err != nil {
			panic(err)
		}
		for _, value := range tasks {
			key := struct{ C_id, T_id int }{value.OwnerID, value.ID}
			dao.taskMap[key] = value
		}
	}
	return &dao
}

// question
// нужно ли проверять индекс или доверяться обработчикам комманд
func (dao *FileTaskDao) Create(task Task) error {
	file, err := os.OpenFile(dao.path, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	file.Seek(-1, 2)
	task_json, err := json.Marshal(task)
	if err != nil {
		return err
	}
	write_data := string(task_json) + "]"

	if len(dao.taskMap) != 0 {
		write_data = "," + write_data
	} else {
		file.Seek(0, 0)
		write_data = "[" + write_data
	}
	_, err = io.WriteString(file, write_data)
	if err != nil {
		return err
	}
	return nil
}

func (dao *FileTaskDao) Read(key struct{ C_id, T_id int }) (Task, error) {
	if value, ok := dao.taskMap[key]; ok {
		return value, nil
	} else {
		return Task{}, errors.New("index out of range")
	}
}

func (dao *FileTaskDao) ReadAll() []Task {
	res := make([]Task, 0, len(dao.taskMap))
	for _, value := range dao.taskMap {
		res = append(res, value)
	}

	//надо не надо непонятно
	sort.Slice(res, func(i int, j int) bool {
		return res[i].OwnerID < res[j].OwnerID ||
			(res[i].OwnerID == res[j].OwnerID && res[i].ID < res[j].ID)
	})
	return res
}

func (dao *FileTaskDao) Update(key struct{ C_id, T_id int }, upd_task Task) error {
	if _, ok := dao.taskMap[key]; !ok {
		return errors.New("index out of range")
	}
	dao.taskMap[key] = upd_task
	file, err := os.OpenFile(dao.path, os.O_WRONLY, 0666)
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

func (dao *FileTaskDao) Delete(key struct{ C_id, T_id int }) error {
	if _, ok := dao.taskMap[key]; !ok {
		return errors.New("index out of range")
	}
	delete(dao.taskMap, key)
	file, err := os.OpenFile(dao.path, os.O_WRONLY, 0666)
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

func (dao FileTaskDao) marshal() ([]byte, error) {
	task_list := make([]Task, 0, len(dao.taskMap))
	for _, value := range dao.taskMap {
		task_list = append(task_list, value)
	}
	return json.Marshal(task_list)
}
