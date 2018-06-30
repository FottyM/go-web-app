package dbinitializer

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/kenigbolo/go-web-app/model"
	_ "github.com/lib/pq" // Database initializer
)

// ConnectToDatabase function
func ConnectToDatabase() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost/go_web_app?sslmode=disable")
	if err != nil {
		log.Fatalln(fmt.Errorf("Unable to connect to database: %v", err))
	}
	model.SetDatabase(db)
	return db
}
