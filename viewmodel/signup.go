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
}

// NewSignup function
func NewSignup() Signup {
	return Signup{Title: "Lemonade Stand Supply", Active: "home"}
}
