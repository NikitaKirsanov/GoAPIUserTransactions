package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func NewServer(s *Service) {
	r := mux.NewRouter()

	r.HandleFunc(
		"/ping",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode("PONG")
		}).Methods("GET")
	r.HandleFunc("/users", s.GetUsers).Methods("GET")
	r.HandleFunc("/transfer", s.MakeTransfer).Methods("POST")
	fmt.Println("Server started")
	http.ListenAndServe(":8080", r)
}
