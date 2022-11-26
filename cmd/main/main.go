package main

import (
	"log"
	"net/http"

	"github.com/doduykhang/musik/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.PathPrefix("/assets").Handler(http.FileServer(http.Dir("./assets/")))
	routes.RegisterRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
