package server

import (
	"KirsanovStavkaTV/internal/contracts"
	"KirsanovStavkaTV/internal/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Service struct {
	dbProvider contracts.DatabaseProvider
}

func NewService(provider contracts.DatabaseProvider) *Service {
	return &Service{
		dbProvider: provider,
	}
}

func (s Service) GetUsers(w http.ResponseWriter, r *http.Request) {
	users := s.dbProvider.GetUsers()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (s Service) MakeTransfer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	if params["UserFrom"] == "" || params["UserTo"] == "" || params["Amount"] == "" {
		json.NewEncoder(w).Encode("Invalid Argument")
	}

	userFromId, err := strconv.Atoi(params["UserFrom"])
	if err != nil {
		json.NewEncoder(w).Encode("Invalid Argument UserFrom")
	}

	userToId, err := strconv.Atoi(params["UserTo"])
	if err != nil {
		json.NewEncoder(w).Encode("Invalid Argument userFrom")
	}

	amount64, err := strconv.ParseUint(params["Amount"], 10, 32)
	if err != nil {
		json.NewEncoder(w).Encode("Invalid Argument Amount")
	}
	amount := uint(amount64)

	transaction := new(models.Transaction)
	transaction.UserFrom = userFromId
	transaction.UserTo = userToId
	transaction.Amount = amount

	err = s.dbProvider.MakeTransfer(transaction)
	if err != nil {
		json.NewEncoder(w).Encode("Couldn't save transaction contact N.Kirsanov")
	}
}
