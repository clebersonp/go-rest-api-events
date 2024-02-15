package models

import (
	"database/sql"
	"example.com/rest-api-events/db"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User) Save() error {
	query := `
		INSERT INTO users (email, password) VALUES (?, ?)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer func(s *sql.Stmt) {
		_ = s.Close()
	}(stmt)

	result, err := stmt.Exec(u.Email, u.Password)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	u.ID = id
	return err
}
