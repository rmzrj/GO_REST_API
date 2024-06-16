package models

import (
	"errors"

	"rest_api_example.com/db"
	"rest_api_example.com/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?,?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	u.ID = userId
	return err
}

func (u *User) ValidateCred() error {
	query := "SELECT id, password FROM users WHERE email = ?"

	row := db.DB.QueryRow(query, u.Email)

	var e User

	err := row.Scan(&u.ID, &e.Password)

	if err != nil {
		return err
	}

	passwordValid := utils.CheckPassordHash(u.Password, e.Password)

	if !passwordValid {
		return errors.New("INVALID")
	}

	return nil

}
