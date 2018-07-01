package model

import (
	"testing"

	helper "github.com/kenigbolo/go-web-app/helpers"
)

func TestLoginSendsCorrectPasswordHash(t *testing.T) {
	testDB := new(mockDB)
	testDB.returnedRow = &mockRow{}
	db = testDB

	password := "the password"
	email := "the email"
	Login(email, password)

	pwd := helper.EncryptPassword(email, password)

	if testDB.lastArgs[1] != pwd {
		t.Errorf("Login function failed to send correct password hash to database")
	}
}
