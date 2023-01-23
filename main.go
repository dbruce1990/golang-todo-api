package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	InitRoutes()
	log.Fatal(http.ListenAndServe(":8000", r))
}
