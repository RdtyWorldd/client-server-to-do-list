package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/RdtyWorldd/client-server-to-do-list/server/client"
	cdao "github.com/RdtyWorldd/client-server-to-do-list/server/client/dao"
	"github.com/RdtyWorldd/client-server-to-do-list/server/client/handlers"
	"github.com/RdtyWorldd/client-server-to-do-list/server/task"
)

func generateRandomClient() client.Client {
	// Генерация случайных имен и логинов
	firstNames := []string{"Иван", "Петр", "Мария", "Анна", "Алексей", "Елена", "Сергей", "Ольга"}
	lastNames := []string{"Иванов", "Петров", "Сидоров", "Кузнецов", "Смирнов", "Попов", "Васильев"}

	firstName := firstNames[rand.Intn(len(firstNames))]
	lastName := lastNames[rand.Intn(len(lastNames))]
	login := fmt.Sprintf("%s%d", firstName, rand.Intn(1000))
	password := fmt.Sprintf("pass%d", rand.Intn(10000))
	return *client.NewClient(login, password, firstName, lastName)
}

type HandlerFunc func(http.ResponseWriter, *http.Request)

func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(w, r)
}

func TaskListLogHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("User login")
}

func main() {
	// rand.Seed(time.Now().UnixNano())

	// // Создаем массив из 10 случайных клиентов
	// clients := make([]client.Client, 10)
	// for i := 0; i < 10; i++ {
	// 	clients[i] = generateRandomClient()
	// }

	//ih := http.HandlerFunc(TaskListLogHandler)
	task_dao := task.NewFileTaskDao("./json/tasks.json")
	client_dao := cdao.NewFileClientDao("./json/clients.json", task_dao)
	// dao.Delete(2)
	mux := http.NewServeMux()
	mux.Handle("/singup", handlers.NewSingUpHandler(client_dao))
	mux.Handle("/login", handlers.NewLoginHandler(client_dao))
	// mux.Handle("/task-list", ih)
	mux.Handle("/task-list", handlers.NewTaskListHandler(client_dao))
	log.Print("Сервер запущен...")

	// Запуск сервера на порту 8080
	http.ListenAndServe(":8080", mux)
}
