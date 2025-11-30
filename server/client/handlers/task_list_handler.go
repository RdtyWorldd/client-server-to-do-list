package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/RdtyWorldd/client-server-to-do-list/server/actions"
	"github.com/RdtyWorldd/client-server-to-do-list/server/client"
	"github.com/RdtyWorldd/client-server-to-do-list/server/client/dao"
	"github.com/RdtyWorldd/client-server-to-do-list/server/task"
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
	switch r.Method {
	case "GET":
		handler.get(w, r, client)
	case "POST":
		handler.postActionsRouting(w, r, client)
	}
}

func (handler *TaskListHandler) get(w http.ResponseWriter, r *http.Request, client client.Client) {
	marshaled, err := json.Marshal(client.GetTaskList())
	if err != nil {
		http.Error(w, "Server couldn't load page", http.StatusInternalServerError)
		return
	}
	w.Write(marshaled)
}

// не размазано по всем методам для jsonApi
func (handler *TaskListHandler) postActionsRouting(w http.ResponseWriter, r *http.Request, c client.Client) {
	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Panic(err)
	}

	var req_data struct {
		Action string    `json:"action"`
		Object task.Task `json:"object"`
	}

	err = json.Unmarshal(data, &req_data)
	if err != nil {
		panic(err)
	}
	switch req_data.Action {
	case actions.ADD:
		handler.add(w, r, c, req_data.Object)
	case actions.DEL:
		handler.delete(w, r, c, req_data.Object)
	case actions.UPD:
		handler.update(w, r, c, req_data.Object)
	default:
		log.Println("Такого действия не существует")
		//выдать что такого действия не существует
	}
}

func (handler *TaskListHandler) add(w http.ResponseWriter, r *http.Request, c client.Client, t task.Task) {
	upd_c, err := handler.client_dao.AddTask(c.GetID(), t)
	if err != nil {
		//узнать что не получилось
		//написать в ответ что таска не записанаб причина
		log.Panicln(err)
		http.Error(w, "Can't create new task", http.StatusBadRequest)
		return
	}
	// вопросы на проверку ошибки и ответа клиенту
	// сообщение что все ок и новый список тасок или обновленную таску
	handler.get(w, r, upd_c)
}

func (handler *TaskListHandler) delete(w http.ResponseWriter, r *http.Request, c client.Client, t task.Task) {
	upd_c, err := handler.client_dao.DeleteTask(c.GetID(), t.ID)
	if err != nil {
		log.Panicln(err)
		http.Error(w, "Can't delete task", http.StatusBadRequest)
		return
	}
	handler.get(w, r, upd_c)
}

func (handler *TaskListHandler) update(w http.ResponseWriter, r *http.Request, c client.Client, t task.Task) {
	upd_c, err := handler.client_dao.UpdateTask(c.GetID(), t)
	if err != nil {
		log.Panicln(err)
		http.Error(w, "Can't update new task", http.StatusBadRequest)
		return
	}
	handler.get(w, r, upd_c)
}
