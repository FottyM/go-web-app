package main

import (
	"net/http"

	"github.com/kenigbolo/go-web-app/controller"
	"github.com/kenigbolo/go-web-app/templates"
)

var filePath = "../src/github.com/kenigbolo/go-web-app/"

func main() {
	templates := templates.PopulateTemplates(filePath)
	controller.Startup(filePath, templates)
	http.ListenAndServe(":3000", nil)
}
