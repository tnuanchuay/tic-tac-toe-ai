package main

import (
	"github.com/tspn/tic-tac-toe-AI/controllers"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

func main(){
	m_route := mux.NewRouter()
	route := controllers.Init()

	for k, v := range route.GET {
		fmt.Println("register", k, v)
		m_route.HandleFunc(k, v).Methods("GET")
	}

	for k, v := range route.POST {
		fmt.Println("register", k, v)
		m_route.HandleFunc(k, v).Methods("POST")
	}

	for k, v := range route.PUT {
		fmt.Println("register", k, v)
		m_route.HandleFunc(k, v).Methods("PUT")
	}

	for k, v := range route.DELETE {
		fmt.Println("register", k, v)
		m_route.HandleFunc(k, v).Methods("DELETE")
	}

	http.Handle("/", m_route)

	http.ListenAndServe(":8080", nil)
}
