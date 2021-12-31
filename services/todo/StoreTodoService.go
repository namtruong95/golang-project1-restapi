package services

import (
	"project1/model"
)

func StoreTodo(todo model.Todo) model.Todo {
	model.NewTodoID(&todo, int32(len(todos)))
	model.NewTodoTimestamp(&todo)

	todos = append(todos, todo)
	return todo
}
