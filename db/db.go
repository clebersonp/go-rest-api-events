package db

// the database/sql package needs a driver under the hood,
// that is we need to put _ in front of driver package to go does not remove it

// to install driver package: go get github.com/mattn/go-sqlite3

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	// workaround for this problem using DB, err := ->
	// creates a new local variable named DB within the scope of the InitDB function,
	// shadowing the global variable DB declared at the package level.
	var err error
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		return
	}

	if err != nil {
		panic("Cannot open database connection!")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	// create all tables
	createTables()
}

func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    name TEXT NOT NULL,
	    description TEXT NOT NULL,
	    location TEXT NOT NULL,
	    date_time DATETIME NOT NULL,
	    user_id INTEGER	    
	)
`
	_, err := DB.Exec(createEventsTable)
	if err != nil {
		panic("Could not create table 'events'!")
	}
}
