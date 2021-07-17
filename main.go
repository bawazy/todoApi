package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type todo struct {
	Id        string `json:"id"`
	Task      string `json:"task"`
	Completed string `json:"completed"`
}

var todos []todo

func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func getTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, todo := range todos {
		if todo.Id == params["id"] {
			json.NewEncoder(w).Encode(todo)
			return
		}
	}
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var todo todo
	_ = json.NewDecoder(r.Body).Decode(&todo)
	todo.Id = strconv.Itoa(rand.Intn(100000000))
	todos = append(todos, todo)
	json.NewEncoder(w).Encode(todo)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, todo := range todos {
		if todo.Id == params["id"] {
			todos = append(todos[:index], todos[1+index:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(todos)
}
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range todos {
		if item.Id == params["id"] {
			todos = append(todos[:index], todos[1+index:]...)
			var todo todo
			_ = json.NewDecoder(r.Body).Decode(&todo)
			todo.Id = params["id"]
			todos = append(todos, todo)
			json.NewEncoder(w).Encode(todo)

		}
	}
}

func main() {
	r := mux.NewRouter()

	todos = append(todos, todo{Id: "1", Task: "Do Laundry", Completed: "false"})
	todos = append(todos, todo{Id: "2", Task: "Clean The House", Completed: "false"})

	r.HandleFunc("/todo", getTodos).Methods("GET")
	r.HandleFunc("/todo/{id}", getTodo).Methods("GET")
	r.HandleFunc("/todo", createTodo).Methods("POST")
	r.HandleFunc("/todo/{id}", deleteTodo).Methods("DELETE")
	r.HandleFunc("/todo/{id}", UpdateTodo).Methods("PUT")
	fmt.Printf("Starting Server at port 8080 \n")
	log.Fatal(http.ListenAndServe(":8080", r))
}
