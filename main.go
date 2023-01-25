package main

import (
	"log"
	"net/http"

	"github.com/dbruce1990/golang-todo-api/db"
	"github.com/dbruce1990/golang-todo-api/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := db.Init()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	r := mux.NewRouter()
	app := routes.App{}
	app.Init(r, db)

	log.Fatal(http.ListenAndServe(":8000", r))
}
