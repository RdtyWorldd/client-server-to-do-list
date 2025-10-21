package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/RdtyWorldd/client-server-to-do-list/server/client"
	"github.com/RdtyWorldd/client-server-to-do-list/server/dao"
)

type LoginHandler struct {
	dao dao.CrudDao[client.Client]
}

func NewLoginHandler(dao dao.CrudDao[client.Client]) *LoginHandler {
	return &LoginHandler{dao}
}
func (handler *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Panic(err)
	}

	var req_data struct {
		Login    string
		Password string
	}

	err = json.Unmarshal(data, &req_data)
	if err != nil {
		panic(err)
	}

	clients := handler.dao.ReadAll()
	is_exist := false
	var client_pos int
	for i, value := range clients {
		if value.GetLogin() == req_data.Login {
			is_exist = true
			client_pos = i
			break
		}
	}
	if (!is_exist) || (clients[client_pos].GetPassword() != req_data.Password) {
		io.WriteString(w, "Wrong login or password, please try again or Sign up")
		return
	}
	log.Println("In login handler, user exist, password correct")
	http.Redirect(w, r, r.Host+"/task-list", http.StatusFound)
}
