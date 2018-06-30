package main

import (
	"net/http"

	"github.com/kenigbolo/go-web-app/controller"
	"github.com/kenigbolo/go-web-app/database"
	"github.com/kenigbolo/go-web-app/middleware"
	"github.com/kenigbolo/go-web-app/templates"

	_ "github.com/lib/pq"
)

var filePath = "../src/github.com/kenigbolo/go-web-app/"

func main() {
	templates := templates.PopulateTemplates(filePath)
	// db := dbinitializer.connectToDatabase()
	// db := connectToDatabase()
	db := dbinitializer.ConnectToDatabase()
	defer db.Close()
	controller.Startup(filePath, templates)
	http.ListenAndServe(":8000", &middleware.TimeoutMiddleware{Next: new(middleware.GzipMiddleware)})
}

// func connectToDatabase() *sql.DB {
// 	db, err := sql.Open("postgres", "postgres://kenigbolo:kenigbolo@localhost/gowebapp?sslmode=disable")
// 	if err != nil {
// 		log.Fatalln(fmt.Errorf("Unable to connect to database: %v", err))
// 	}
// 	model.SetDatabase(db)
// 	return db
// }
