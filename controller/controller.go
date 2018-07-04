package controller

import (
	"html/template"
	"net/http"
)

// Startup controller initializing function for templates
func Startup(filePath string, templates map[string]*template.Template) {
	// Set up Controllers
	homeController := home{
		homeTemplate:         templates["home.html"],
		standLocatorTemplate: templates["stand_locator.html"],
	}
	authController := authentication{
		loginTemplate:  templates["login.html"],
		signupTemplate: templates["signup.html"],
	}
	shopController := shop{
		shopTemplate:     templates["shop.html"],
		categoryTemplate: templates["shop_details.html"],
		productTemplate:  templates["shop_detail.html"],
	}
	standLocatorController := standLocator{standLocatorTemplate: templates["stand_locator.html"]}
	// Register Routes
	homeController.registerRoutes()
	shopController.registerRoutes()
	standLocatorController.registerRoutes()
	authController.registerRoutes()
	// Serve static assests
	http.Handle("/img/", http.FileServer(http.Dir(filePath+"public")))
	http.Handle("/css/", http.FileServer(http.Dir(filePath+"public")))
	http.Handle("/js/", http.FileServer(http.Dir(filePath+"public")))
}
