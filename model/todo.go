package model

import "time"

type Todo struct {
	ID        int32  `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

func NewTodoID(todo *Todo, count int32) {
	todo.ID = count + 1
	todo.Type = "Todo"
}

func NewTodoTimestamp(todo *Todo) {
	todo.CreatedAt = time.Now().Unix()
	todo.UpdatedAt = time.Now().Unix()
}
