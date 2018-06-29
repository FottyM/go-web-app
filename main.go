package main

import (
	"net/http"
)

func main() {
	filePath := "../src/github.com/kenigbolo/go-web-app/"
	http.ListenAndServe(":8000", http.FileServer(http.Dir(filePath+"public")))
}
