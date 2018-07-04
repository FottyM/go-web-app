package viewmodel

import (
	"fmt"

	"github.com/kenigbolo/go-web-app/model"
)

// Shop struct
type Shop struct {
	Title        string
	Active       string
	Alert        string
	AlertMessage string
	AlertDanger  string
	AlertSuccess string
	Categories   []Category
}

// Category struct
type Category struct {
	URL           string
	ImageURL      string
	Title         string
	Alt           string
	Description   string
	IsOrientRight bool
}

// NewShop func
func NewShop(categories []model.Category) Shop {
	result := Shop{
		Title:  "Lemonade Stand Supply - Shop",
		Active: "shop",
		Alert:  "invisible",
	}
	result.Categories = make([]Category, len(categories))
	for i := 0; i < len(categories); i++ {
		vm := categorytoVM(categories[i])
		vm.IsOrientRight = i%2 == 1
		result.Categories[i] = vm
	}
	return result
}

func categorytoVM(c model.Category) Category {
	return Category{
		URL:         fmt.Sprintf("/shop/%v", c.ID),
		ImageURL:    c.ImageURL,
		Title:       c.Title,
		Description: c.Description,
	}
}
