package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joaomarcuslf/color-generator/handlers"
)

type Server struct {
	Port string
}

func NewServer(port string) *Server {
	return &Server{
		Port: port,
	}
}

func (a *Server) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/{color}", handlers.GenerateColor).Methods("GET")

	fmt.Println("Server running on port: " + a.Port)

	log.Fatal(http.ListenAndServe(":"+a.Port, router))
}
