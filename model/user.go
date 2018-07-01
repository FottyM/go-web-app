package model

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	helper "github.com/kenigbolo/go-web-app/helpers"
)

// User model
type User struct {
	ID        int
	Email     string
	Password  string
	FirstName string
	LastName  string
	LastLogin *time.Time
}

// Login function
func Login(email, password string) (*User, error) {
	result := &User{}
	pwd := helper.EncryptPassword(email, password)
	row := db.QueryRow(`
		SELECT id, email, firstname, lastname
		FROM public.user
		WHERE email = $1
		  AND password = $2`, email, pwd)
	err := row.Scan(&result.ID, &result.Email, &result.FirstName, &result.LastName)
	switch {
	case err == sql.ErrNoRows:
		return nil, fmt.Errorf("User not found")
	case err != nil:
		return nil, err
	}
	t := time.Now()
	_, err = db.Exec(`
		UPDATE public.user
		SET lastlogin = $1
		WHERE id = $2`, t, result.ID)
	if err != nil {
		log.Printf("Failed to update login time for user %v to %v: %v", result.Email, t, err)
	}
	return result, nil
}
