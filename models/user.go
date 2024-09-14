package models

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Fidel-wolee/event-booking/db"
	"github.com/Fidel-wolee/event-booking/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUser struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (user *User) Save() error {
	query := `INSERT INTO users(name, email, password)
	VALUES(?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	hashedPassword, err := utils.HashPassword(user.Password)
	fmt.Println(hashedPassword)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(user.Name, user.Email, hashedPassword)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = userId
	return nil
}

func GetUser(id int64) (*User, error) {
	query := "SELECT id, name, email, password FROM users WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		// Handle no rows case if necessary
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no user found with id %d", id)
		}
		return nil, err
	}

	return &user, nil
}
func (u *LoginUser) ValidateCredentials() (*LoginUser, error) {
	query := "SELECT id, password FROM users WHERE email=?"
	row := db.DB.QueryRow(query, u.Email)

	// Variables to scan the result
	var storedID int64
	var storedPassword string

	// Scan the result into the variables
	err := row.Scan(&storedID, &storedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("no user found with the provided email")
		}
		return nil, err
	}

	// Check if the provided password is correct
	passwordIsValid := utils.CheckPasswordHash(u.Password, storedPassword)
	if !passwordIsValid {
		return nil, errors.New("incorrect password")
	}

	// Set the user's ID after successful validation
	u.ID = storedID

	// Return the user with the populated ID
	return u, nil
}
