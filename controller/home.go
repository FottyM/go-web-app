package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/kenigbolo/go-web-app/model"
	"github.com/kenigbolo/go-web-app/viewmodel"
)

type home struct {
	homeTemplate         *template.Template
	standLocatorTemplate *template.Template
	loginTemplate        *template.Template
	signupTemplate       *template.Template
}

func (h home) registerRoutes() {
	http.HandleFunc("/", h.handleHome)
	http.HandleFunc("/home", h.handleHome)
	http.HandleFunc("/stand_locator", h.handleStandLocator)
	http.HandleFunc("/login", h.handleLogin)
	http.HandleFunc("/signup", h.handleSignup)
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

func (h home) handleLogin(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewLogin()
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Println(fmt.Errorf("Error logging in: %v", err))
		}
		email := r.Form.Get("email")
		password := r.Form.Get("password")
		var errMsg string
		if user, errMsg := model.Login(email, password); errMsg == nil {
			log.Printf("User has logged in: %v\n", user)
			http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)
			return
		}
		log.Printf("Failed to log user in with email: %v, error was: %v\n", email, errMsg)
		vm.Email = email
		vm.Password = password
	}
	w.Header().Add("Content-Type", "text/html")
	h.loginTemplate.Execute(w, vm)
}

func (h home) handleSignup(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewSignup()
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Println(fmt.Errorf("Error signing up: %v", err))
		}
		email := r.Form.Get("email")
		firstName := r.Form.Get("first-name")
		lastName := r.Form.Get("last-name")
		password := r.Form.Get("password")
		passwordConfirmation := r.Form.Get("password-confirmation")
		if password != passwordConfirmation {
			log.Println("Password Missmatch")
			http.Redirect(w, r, "/signup", http.StatusTemporaryRedirect)
			return
		}
		var errMsg string
		if user, errMsg := model.Signup(email, firstName, lastName, password); errMsg == nil {
			log.Printf("User successfully signed up: %v\n", user)
			http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)
			return
		}
		log.Printf("Failed to signup user with email: %v, error was: %v\n", email, errMsg)
		vm.Email = email
		vm.FirstName = firstName
		vm.LastName = lastName
		vm.Password = password
		vm.PasswordConfirmation = passwordConfirmation
	}
	w.Header().Add("Content-Type", "text/html")
	h.signupTemplate.Execute(w, vm)
}
