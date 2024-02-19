package models

import (
	"errors"

	"example.com/booking-app/db"
	"example.com/booking-app/utils"
)

type User struct {
	ID       int64
	Name     string
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := `
	INSERT INTO users(name, email, password) 
	VALUES (?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Name, u.Email, hashedPassword)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()

	u.ID = id
	return err
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"

	row := db.DB.QueryRow(query, u.Email)

	var rowPass string

	err := row.Scan(&u.ID, &rowPass)
	if err != nil {
		return errors.New("credentials invalid")
	}

	passIsValid := utils.CheckPasswordHash(u.Password, rowPass)

	if !passIsValid {
		return errors.New("credentials invalid")
	}

	return nil
}
