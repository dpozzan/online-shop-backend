package models

import (
	"errors"

	"github.com/dpozzan/db"
	"github.com/dpozzan/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User) Save() (int64, error) {
	checkQuery := "SELECT email FROM users WHERE email=?"

	row := db.DB.QueryRow(checkQuery, u.Email)

	var retriviedEmail string

	row.Scan(&retriviedEmail)

	if retriviedEmail != "" {
		return 0, errors.New("email already used")
	}

	query := `
	INSERT INTO users (email, password)
	VALUES (?, ?)
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)

	if err != nil {
		return 0, err
	}

	userID, err := result.LastInsertId()

	if err != nil {
		return 0, nil
	}

	return userID, nil
}


func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"

	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string

	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return err
	}

	isValidPassword := utils.CheckPasswordHash(retrievedPassword, u.Password)

	if !isValidPassword {
		return errors.New("Unauthorized")
	}

	return nil


}