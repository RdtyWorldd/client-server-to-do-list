package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/RdtyWorldd/client-server-to-do-list/server/client"
	"github.com/RdtyWorldd/client-server-to-do-list/server/dao"
	"github.com/RdtyWorldd/client-server-to-do-list/server/task"
)

type TaskListHandler struct {
	client_dao dao.CrudDao[client.Client]
	task_dao   dao.CrudDao[task.Task]
}

func NewTaskListHandler(client_dao dao.CrudDao[client.Client], task_dao dao.CrudDao[task.Task]) *TaskListHandler {
	return &TaskListHandler{client_dao, task_dao}
}

func (handler *TaskListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if _, ok := r.Header["Cookie"]; !ok {
		//redirect
		log.Panic("Request didnt had cookies")
	}
	cookie, err := r.Cookie("session_id")
	if err != nil {
		log.Panic(err)
	}

	c_id, err := strconv.Atoi(cookie.Value)
	if err != nil {
		log.Panic(err)
	}
	_, err = handler.client_dao.Read(c_id)
	if err != nil {

	}
	tasks := handler.task_dao.ReadAll()
	var client_tasks []task.Task
	for _, value := range tasks {
		if value.OwnerID == c_id {
			client_tasks = append(client_tasks, value)
		}
	}
	marshaled, _ := json.Marshal(client_tasks)
	w.Write(marshaled)
}
