package viewmodel

// StandLocator struct
type StandLocator struct {
	Title        string
	Active       string
	Alert        string
	AlertMessage string
	AlertDanger  string
	AlertSuccess string
}

// StandCoordinate struct
type StandCoordinate struct {
	Title     string  `json:"title"`
	Latitude  float32 `json:"lat"`
	Longitude float32 `json:"lng"`
}

// NewStandLocator FUNCTION
func NewStandLocator() StandLocator {
	return StandLocator{
		Active: "standlocator",
		Title:  "Lemonade Stand Supply - Stand Locator",
		Alert:  "invisible",
	}
}
