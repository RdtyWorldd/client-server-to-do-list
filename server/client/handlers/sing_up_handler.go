package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/RdtyWorldd/client-server-to-do-list/server/client"
	"github.com/RdtyWorldd/client-server-to-do-list/server/dao"
)

type SignUpHandler struct {
	dao dao.CrudDao[client.Client]
}

func NewSingUpHandler(dao dao.CrudDao[client.Client]) *SignUpHandler {
	return &SignUpHandler{dao}
}

func (handler *SignUpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Panic(err)
	}
	var parse struct {
		Login     string
		Password  string
		FirstName string
		LastName  string
	}
	err = json.Unmarshal(data, &parse)
	if err != nil {
		log.Panic(err)
	}
	handler.dao.Create(*client.NewClient(parse.Login, parse.Password, parse.FirstName, parse.LastName))
	io.WriteString(w, "New account created")
}
