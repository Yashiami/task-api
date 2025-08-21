package database

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func InitDB() {
	connName := "postgres://postgres:1234@localhost:5432/task?sslmode=disable"
	var err error
	DB, err = sql.Open("pgx", connName)
	if err != nil {
		panic(err)
	}
}
