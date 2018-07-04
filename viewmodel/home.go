package viewmodel

// Home struct
type Home struct {
	Title        string
	Active       string
	Alert        string
	AlertMessage string
	AlertDanger  string
	AlertSuccess string
}

// NewHome function
func NewHome() Home {
	result := Home{
		Active: "home",
		Title:  "Lemonade Stand Supply",
		Alert:  "invisible",
	}
	return result
}
