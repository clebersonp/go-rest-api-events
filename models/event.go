package models

import (
	"database/sql"
	"errors"
	"example.com/rest-api-events/db"
	"time"
)

// Event - struct tags for json deserialized string into structs and
// binding feature from gin package to required missing fields
type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"date_time" binding:"required"`
	UserID      int       `json:"user_id"`
}

func (e *Event) Save() error {
	query := `INSERT INTO events (name, description, location, date_time, user_id)
	VALUES (?, ?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer func(s *sql.Stmt) {
		_ = stmt.Close()
	}(stmt)

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	e.ID = id
	return err
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	var events []Event
	for rows.Next() {
		e := Event{}
		// the order to be passed need to be the same of column declaration on DDL of table
		err = rows.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, e)
	}

	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := `SELECT * FROM events e WHERE e.id = ?`

	row := db.DB.QueryRow(query, id)
	if row.Err() != nil {
		return nil, row.Err()
	}

	e := &Event{}
	err := row.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			e = nil
		} else {
			return nil, err
		}
	}

	// result or no result, and no error
	return e, nil
}

func (e *Event) Update() error {
	query := `UPDATE events
	SET name = ?, description = ?, location = ?, date_time = ?
	WHERE id = ?`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer func(s *sql.Stmt) {
		_ = s.Close()
	}(stmt)

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected <= 0 || rowsAffected > 1 {
		return errors.New(" Something went wrong. Try again later")
	}
	return nil
}
