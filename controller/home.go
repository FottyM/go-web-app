package controller

import (
	"html/template"
	"net/http"

	"github.com/kenigbolo/go-web-app/viewmodel"
)

type home struct {
	homeTemplate         *template.Template
	standLocatorTemplate *template.Template
}

func (h home) registerRoutes() {
	http.HandleFunc("/", h.handleHome)
	http.HandleFunc("/home", h.handleHome)
	http.HandleFunc("/stand_locator", h.handleStandLocator)
}

func (h home) handleHome(w http.ResponseWriter, r *http.Request) {
	if pusher, ok := w.(http.Pusher); ok {
		pusher.Push("../src/github.com/kenigbolo/go-web-app/public/css/app.css", &http.PushOptions{
			Header: http.Header{"Content-Type": []string{"text/css"}},
		})
	}
	vm := viewmodel.NewHome()
	w.Header().Add("Content-Type", "text/html")
	h.homeTemplate.Execute(w, vm)
}

func (h home) handleStandLocator(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewStandLocator()
	w.Header().Add("Content-Type", "text/html")
	h.standLocatorTemplate.Execute(w, vm)
}
