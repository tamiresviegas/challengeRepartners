package main

import (
	"net/http"

	"github.com/tamiresviegas/challengeRepartners/handler"
)

func main() {
	http.HandleFunc("/calculate", handler.CalculatePacks)
	http.HandleFunc("/addpackage", handler.AddPackage)
	http.HandleFunc("/removepackage", handler.RemovePackage)
	http.Handle("/", http.FileServer(http.Dir("./ui")))
	http.ListenAndServe(":8080", nil)
}
