package model

import "database/sql"

var db *sql.DB

// SetDatabase func
func SetDatabase(database *sql.DB) {
	db = database
}
