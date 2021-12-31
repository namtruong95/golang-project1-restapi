package services

import (
	"project1/model"
)

var posts []model.Post = make([]model.Post, 0)

func ListPost() []model.Post {
	return posts
}
