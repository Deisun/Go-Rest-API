package main

import (
	"github.com/gorilla/mux"
	"log"
	"myBigNerdRanchProject/routes"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
