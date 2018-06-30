package main

import (
	"net/http"

	"github.com/kenigbolo/go-web-app/controller"
	"github.com/kenigbolo/go-web-app/database"
	"github.com/kenigbolo/go-web-app/middleware"
	"github.com/kenigbolo/go-web-app/templates"

	_ "net/http/pprof"

	_ "github.com/lib/pq"
)

var filePath = "../src/github.com/kenigbolo/go-web-app/"

func main() {
	templates := templates.PopulateTemplates(filePath)
	db := dbinitializer.ConnectToDatabase()
	defer db.Close()
	controller.Startup(filePath, templates)
	go http.ListenAndServe(":8080", nil)
	http.ListenAndServeTLS(":8000", filePath+"cert.pem", filePath+"key.pem", &middleware.TimeoutMiddleware{Next: new(middleware.GzipMiddleware)})
}
