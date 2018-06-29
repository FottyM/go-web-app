package main

import (
	"html/template"
	"log"
	"net/http"
)

var filePath = "../src/github.com/kenigbolo/go-web-app/"

func main() {
	templates := populateTemplates()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestedFile := r.URL.Path[1:]
		t := templates.Lookup(requestedFile + ".html")
		if t != nil {
			err := t.Execute(w, nil)
			if err != nil {
				log.Println(err)
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})
	http.Handle("/img/", http.FileServer(http.Dir(filePath+"public")))
	http.Handle("/css/", http.FileServer(http.Dir(filePath+"public")))
	http.ListenAndServe(":8000", nil)
}

func populateTemplates() *template.Template {
	result := template.New("templates")
	basePath := filePath + "templates"
	template.Must(result.ParseGlob(basePath + "/*.html"))
	return result
}
