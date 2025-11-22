package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/RdtyWorldd/client-server-to-do-list/server/client/dao"
)

type TaskListHandler struct {
	client_dao dao.ClientDao
}

func NewTaskListHandler(client_dao dao.ClientDao) *TaskListHandler {
	return &TaskListHandler{client_dao}
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

	client, err := handler.client_dao.Read(c_id)
	if err != nil {

	}
	marshaled, _ := json.Marshal(client.GetTaskList())
	w.Write(marshaled)
}
