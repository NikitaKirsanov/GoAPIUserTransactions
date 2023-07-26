package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func NewServer(s *Service) {
	r := mux.NewRouter()

	r.HandleFunc("/users", s.GetUsers).Methods("GET")
	r.HandleFunc("/transfer", s.MakeTransfer).Methods("POST")
	fmt.Println("Server started")
	http.ListenAndServe(":8080", r)
}
