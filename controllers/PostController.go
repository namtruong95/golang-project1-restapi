package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"project1/model"
	postservice "project1/services/post"
)

func StorePost(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var post model.Post
	json.Unmarshal(reqBody, &post)
	post = postservice.StorePost(post)

	fmt.Println(post)

	json.NewEncoder(w).Encode(post)
}

func ListPost(w http.ResponseWriter, r *http.Request) {
	data := postservice.ListPost()

	json.NewEncoder(w).Encode(data)
}
