package app

import (
	"fmt"
	"net/http"
)

type Server struct {
}

func (server *Server) Run() {
	http.HandleFunc("/", server.getIndex)
	http.HandleFunc("/weather", server.getWeather)
	fmt.Printf("Server running on port 8080\n")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
