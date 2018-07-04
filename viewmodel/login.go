package viewmodel

// Login struct
type Login struct {
	Title        string
	Active       string
	Email        string
	Password     string
	AlertMessage string
	AlertDanger  string
	AlertSuccess string
	Alert        string
}

// NewLogin function
func NewLogin() Login {
	return Login{Active: "home", Title: "Lemonade Stand Supply", Alert: "invisible"}
}
