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
	createUsersTable()
	createEventsTable()
	createRegistrationTable()
}

func createUsersTable() {
	ddl := `
	CREATE TABLE IF NOT EXISTS users (
    	id INTEGER PRIMARY KEY AUTOINCREMENT,
    	email TEXT NOT NULL UNIQUE,
    	password TEXT NOT NULL
	)`

	_, err := DB.Exec(ddl)
	if err != nil {
		panic("Could not create table 'users'!")
	}
}

func createEventsTable() {
	ddl := `
	CREATE TABLE IF NOT EXISTS events (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    name TEXT NOT NULL,
	    description TEXT NOT NULL,
	    location TEXT NOT NULL,
	    date_time DATETIME NOT NULL,
	    user_id INTEGER,
	    FOREIGN KEY (user_id) REFERENCES users(id)	    
	)
`
	_, err := DB.Exec(ddl)
	if err != nil {
		panic("Could not create table 'events'!")
	}
}

func createRegistrationTable() {
	query := `
	CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY (event_id) REFERENCES events(id),
		FOREIGN KEY (user_id) REFERENCES users(id)
		)
`
	_, err := DB.Exec(query)
	if err != nil {
		panic("Could not create table 'registrations'!")
	}
}
