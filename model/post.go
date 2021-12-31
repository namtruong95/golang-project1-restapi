package model

import "time"

type Post struct {
	ID        int32  `json:"id"`
	Type      string `json:"type"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

func NewPostID(p *Post, count int32) {
	p.ID = count + 1
	p.Type = "Post"
}

func NewPostTimestamp(p *Post) {
	p.CreatedAt = time.Now().Unix()
	p.UpdatedAt = time.Now().Unix()
}
