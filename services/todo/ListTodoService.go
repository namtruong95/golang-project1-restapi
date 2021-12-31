package services

import (
	"project1/model"
)

var todos []model.Todo = make([]model.Todo, 0)

func ListTodo() []model.Todo {
	return todos
}
