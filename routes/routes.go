package routes

import (
	"github.com/gorilla/mux"
	"myBigNerdRanchProject/controllers"
)

var RegisterRoutes = func(router *mux.Router) {
	router.HandleFunc("/post/", controllers.GetPosts).Methods("GET")
	router.HandleFunc("/post/", controllers.CreatePost).Methods("POST")
	router.HandleFunc("/post/{postId}", controllers.GetPost).Methods("GET")
	router.HandleFunc("/post/{postId}", controllers.UpdatePost).Methods("PUT")
	router.HandleFunc("/post/{postId}", controllers.DeletePost).Methods("DELETE")
}
