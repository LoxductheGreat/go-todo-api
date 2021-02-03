package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Todo struct {
	ID   string `json:"id"`
	Body string `json:"body"`
}

var todos []Todo = []Todo{}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", hello)

	router.HandleFunc("/todos/", additem).Methods("POST")

	router.HandleFunc("/todos/", getAll).Methods("GET")

	router.HandleFunc("/todos/{id}", deleteItem).Methods("DELETE")

	http.ListenAndServe(":8000", router)
}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World")
}

func getAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func additem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newTodo Todo

	json.NewDecoder(r.Body).Decode(&newTodo)

	newTodo.ID = strconv.Itoa(len(todos) + 1)

	todos = append(todos, newTodo)

	json.NewEncoder(w).Encode(newTodo)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for idx, item := range todos {
		if item.ID == params["id"] {
			todos = append(todos[:idx], todos[idx+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(todos)
}
