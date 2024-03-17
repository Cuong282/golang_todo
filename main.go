package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Todo struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

var todos []Todo

func GetTodo(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(todos)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(r.Body)
	var todo Todo
	json.NewDecoder(r.Body).Decode(&todo)
	fmt.Println("todo: ", todo)
	todos = append(todos, todo)
	fmt.Println("todos: ", todos)
	json.NewEncoder(w).Encode(todos)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	todoID := r.URL.Query().Get("id")
	fmt.Println("1: ", todoID)
	var updatedTodo Todo
	fmt.Println("2: ", updatedTodo)
	_ = json.NewDecoder(r.Body).Decode(&updatedTodo)
	for index, todo := range todos {
		if strconv.Itoa(todo.ID) == todoID {
			todos[index] = updatedTodo
			break
		}
	}
	json.NewEncoder(w).Encode(todos)
	fmt.Println("xong: ", updatedTodo)

}

func main() {

	http.HandleFunc("/:id", GetTodo)
	http.HandleFunc("/todo/create", CreateTodo)
	http.HandleFunc("/todo/Change", UpdateTodo)

	log.Fatal(http.ListenAndServe(":3000", nil))

}
