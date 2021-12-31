package services

import (
	"project1/model"
)

func StorePost(post model.Post) model.Post {
	model.NewPostID(&post, int32(len(posts)))
	model.NewPostTimestamp(&post)

	posts = append(posts, post)
	return post
}
