package routes

import (
	"github.com/gorilla/mux"
	"myBigNerdRanchProject/controllers"
)

var RegisterRoutes = func(router *mux.Router) {
	router.HandleFunc("/post/", controllers.GetPosts).Methods("GET")
	router.HandleFunc("/post/", controllers.CreatePost).Methods("POST")
	router.HandleFunc("/post/{id}", controllers.GetPost).Methods("GET")
	router.HandleFunc("/post/{id}", controllers.UpdatePost).Methods("PUT")
	router.HandleFunc("/post/{id}", controllers.DeletePost).Methods("DELETE")
}
