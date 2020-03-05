package controllers

import (
	"fmt"
	"myBigNerdRanchProject/models"
	"net/http"
)

var NewPost models.Post

func GetPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get Posts")
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create Post")
}

func GetPost(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Get a Post")
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update Post")
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete Post")
}
