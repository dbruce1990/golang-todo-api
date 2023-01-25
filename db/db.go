package db

import (
	"database/sql"
	"fmt"
	"os"
)

type DB struct {
	*sql.DB
}

func Init() (*DB, error) {
	sqlDB, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s/%s", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME")))
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(5)
	sqlDB.SetMaxIdleConns(3)
	return &DB{sqlDB}, nil
}
