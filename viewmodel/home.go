package viewmodel

// Home struct
type Home struct {
	Title  string
	Active string
}

// NewHome function
func NewHome() Home {
	result := Home{
		Active: "home",
		Title:  "Lemonade Stand Supply",
	}
	return result
}
