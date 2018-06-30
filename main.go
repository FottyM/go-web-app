package main

import (
	"log"
	"net/http"

	"github.com/kenigbolo/go-web-app/templates"
	"github.com/kenigbolo/go-web-app/viewmodel"
)

var filePath = "../src/github.com/kenigbolo/go-web-app/"

func main() {
	templates := templates.PopulateTemplates(filePath)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestedFile := r.URL.Path[1:]
		template := templates[requestedFile+".html"]
		var context interface{}
		switch requestedFile {
		case "shop":
			context = viewmodel.NewShop()
		case "stand_locator":
			context = viewmodel.NewStandLocator()
		default:
			context = viewmodel.NewHome()
		}
		if template != nil {
			err := template.Execute(w, context)
			if err != nil {
				log.Println(err)
			}
		} else {
			w.WriteHeader(404)
		}
	})
	http.Handle("/img/", http.FileServer(http.Dir(filePath+"public")))
	http.Handle("/css/", http.FileServer(http.Dir(filePath+"public")))
	http.ListenAndServe(":8000", nil)
}
