package main

import (
	"net/http"

	"github.com/tspn/tic-tac-toe-AI/controllers"
)

func main() {
	http.HandleFunc("/play", controllers.Play)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	http.Handle("/", http.FileServer(http.Dir("./views")))
	http.ListenAndServe(":8080", nil)
}
