package routes

import (
	"fmt"
	"net/http"

	"github.com/dbruce1990/golang-todo-api/db"
	"github.com/gorilla/mux"
)

type App struct {
	r  *mux.Router
	db *db.DB
}

func (app *App) Init(r *mux.Router, db *db.DB) {
	app.r = r
	app.db = db
	app.r.HandleFunc("/ping", app.pingPongHandler)
	app.InitUserRoutes()
	app.r.NotFoundHandler = http.HandlerFunc(app.handle404)
}

func (app *App) pingPongHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pong"))
}

func (app *App) handle404(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Sorry, the requested resource was not found", http.StatusNotFound)
}

func (app *App) ErrorHandler(w http.ResponseWriter, err error, status int) {
	fmt.Printf("Status: %d\t Error: %s\n", status, err.Error())
	http.Error(w, err.Error(), status)
}
