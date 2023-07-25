package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewServer(s *Service) {
	r := mux.NewRouter()

	r.HandleFunc("/users", s.GetUsers).Methods("GET")
	r.HandleFunc("/transfer", s.MakeTransfer).Methods("POST")

	http.ListenAndServe(":8000", r)
}
