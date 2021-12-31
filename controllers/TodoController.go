package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"project1/model"
	todoservices "project1/services/todo"
)

func StoreTodo(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var todo model.Todo
	json.Unmarshal(reqBody, &todo)
	todo = todoservices.StoreTodo(todo)

	fmt.Println(todo)

	json.NewEncoder(w).Encode(todo)
}

func ListTodo(w http.ResponseWriter, r *http.Request) {
	data := todoservices.ListTodo()

	json.NewEncoder(w).Encode(data)
}
