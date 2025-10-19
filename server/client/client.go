package client

import (
	"encoding/json"

	"github.com/RdtyWorldd/client-server-to-do-list/server/task"
)

type Client struct {
	id        int
	login     string
	password  string
	firstName string
	lastName  string
	taskList  []task.Task
}

// type ClientDTO struct {
// 	Login     string      `json:"login"`
// 	FirstName string      `json:"first"`
// 	LastName  string      `json:"last"`
// 	TaskList  []task.Task `json:"task_list"`
// }

func NewClient(login string, password string, firstName string, lastName string) *Client {
	return &Client{-1, login, password, firstName, lastName, nil}
}

// Геттеры
func (c *Client) GetID() int {
	return c.id
}

func (c *Client) GetLogin() string {
	return c.login
}

func (c *Client) GetPassword() string {
	return c.password
}

func (c *Client) GetFirstName() string {
	return c.firstName
}

func (c *Client) GetLastName() string {
	return c.lastName
}

func (c *Client) GetTaskList() []task.Task {
	return c.taskList
}

// Сеттеры
func (c *Client) SetId(id int) {
	c.id = id
}

func (c *Client) SetLogin(login string) {
	c.login = login
}

func (c *Client) SetPassword(password string) {
	c.password = password
}

func (c *Client) SetFirstName(firstName string) {
	c.firstName = firstName
}

func (c *Client) SetLastName(lastName string) {
	c.lastName = lastName
}

// Методы для работы со списком задач
func (c *Client) AddTask(task task.Task) {
	c.taskList = append(c.taskList, task)
}

func (c *Client) RemoveTask(index int) {
	if index >= 0 && index < len(c.taskList) {
		c.taskList = append(c.taskList[:index], c.taskList[index+1:]...)
	}
}

func (c *Client) ClearTaskList() {
	c.taskList = nil
}

func (c *Client) GetTaskCount() int {
	return len(c.taskList)
}

// MarshalJSON преобразует Client в JSON
func (c *Client) MarshalJSON() ([]byte, error) {
	type ClientJSON struct {
		ID        int    `json:"id"`
		Login     string `json:"login"`
		Password  string `json:"password"` // omitempty скрывает поле если оно пустое
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	clientJSON := ClientJSON{
		ID:        c.id,
		Login:     c.login,
		Password:  c.password, // Внимание: пароль будет в JSON!
		FirstName: c.firstName,
		LastName:  c.lastName,
	}

	return json.Marshal(clientJSON)
}

// UnmarshalJSON преобразует JSON в Client
func (c *Client) UnmarshalJSON(data []byte) error {
	type ClientJSON struct {
		ID        int    `json:"id"`
		Login     string `json:"login"`
		Password  string `json:"password"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	var clientJSON ClientJSON
	if err := json.Unmarshal(data, &clientJSON); err != nil {
		return err
	}

	c.id = clientJSON.ID
	c.login = clientJSON.Login
	c.password = clientJSON.Password
	c.firstName = clientJSON.FirstName
	c.lastName = clientJSON.LastName

	return nil
}
