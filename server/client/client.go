package client

import "github.com/RdtyWorldd/client-server-to-do-list/server/task"

type Client struct {
	login     string
	password  string
	firstName string
	lastName  string
	taskList  []task.Task
}

func NewClient(login string, password string, firstName string, lastName string) *Client {
	return &Client{login, password, firstName, lastName, nil}
}

type ClientDTO struct {
	Login     string      `json:"login"`
	FirstName string      `json:"first"`
	LastName  string      `json:"last"`
	TaskList  []task.Task `json:"task_list"`
}
