package db

import (
	"fmt"

	"github.com/dbruce1990/golang-todo-api/models"
)

func (db *DB) CreateOrUpdateUser(u models.User) {
	err := db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("CreateOrUpdate: connection still alive!")
}
