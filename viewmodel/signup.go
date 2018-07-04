package viewmodel

// Signup struct
type Signup struct {
	Title                string
	Active               string
	Email                string
	Password             string
	PasswordConfirmation string
	FirstName            string
	LastName             string
	Alert                string
	AlertMessage         string
	AlertDanger          string
	AlertSuccess         string
}

// NewSignup function
func NewSignup() Signup {
	return Signup{Title: "Lemonade Stand Supply", Active: "home", Alert: "invisible"}
}
