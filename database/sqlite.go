package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var (
	DB *sql.DB
)

func Setup() {
	var err error
	DB, err = sql.Open("sqlite3", "./database.db")
	if err != nil {
		fmt.Println("Error opening database")
	}

	err = DB.Ping()
	if err != nil {
		fmt.Println("Error pinging database")
	}
}
