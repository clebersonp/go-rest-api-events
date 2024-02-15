package models

import (
	"database/sql"
	"errors"
	"example.com/rest-api-events/db"
	"example.com/rest-api-events/utils"
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

	hashPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashPassword)
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

func (u *User) ValidateCredentials() (bool, error) {
	query := `
		SELECT u.password FROM users u WHERE u.email = ?
	`

	row := db.DB.QueryRow(query, u.Email)
	if row.Err() != nil {
		return false, row.Err()
	}

	var hashedPassword string
	err := row.Scan(&hashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return utils.CheckPasswordHash(hashedPassword, u.Password), nil
}
