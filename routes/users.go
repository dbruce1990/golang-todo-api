package routes

import (
	"encoding/json"
	"net/http"

	"github.com/dbruce1990/golang-todo-api/models"
)

func (app *App) InitUserRoutes() {
	sr := app.r.PathPrefix("/users/").Subrouter()
	sr.HandleFunc("/register", app.createUserHandler).Methods("POST")
	sr.HandleFunc("/destroy", app.destroyUserHandler).Methods("DELETE")
}

// Attempts to register a new user account in the database
func (app *App) createUserHandler(w http.ResponseWriter, r *http.Request) {
	var u models.User

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		app.ErrorHandler(w, err, http.StatusInternalServerError)
		return
	}

	if err := u.ConfirmPasswordMatch(); err != nil {
		app.ErrorHandler(w, err, http.StatusBadRequest)
		return
	}
	app.db.CreateOrUpdateUser(u)
	w.Write([]byte("got here"))
}

func (app *App) destroyUserHandler(w http.ResponseWriter, r *http.Request) {

}
