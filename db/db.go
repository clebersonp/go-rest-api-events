package db

// the database/sql package needs a driver under the hood,
// that is we need to put _ in front of driver package to go does not remove it

// to install driver package: go get github.com/mattn/go-sqlite3

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func DBInit() {
	DB, err := sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Cannot open database connection!")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
}
