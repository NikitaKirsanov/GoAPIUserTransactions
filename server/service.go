package server

import (
	"KirsanovStavkaTV/internal/contracts"
	"KirsanovStavkaTV/internal/models"
	"encoding/json"
	"net/http"
	"strconv"
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
	w.Header().Set("Content-Type", "application/json")

	reqUserFrom := r.FormValue("UserFrom")
	reqUserTo := r.FormValue("UserTo")
	reqAmount := r.FormValue("Amount")

	if reqUserFrom == "" || reqUserTo == "" || reqAmount == "" {
		json.NewEncoder(w).Encode("Invalid Argument")
	}

	userFromId, err := strconv.Atoi(reqUserFrom)
	if err != nil {
		json.NewEncoder(w).Encode("Invalid Argument UserFrom")
	}

	userToId, err := strconv.Atoi(reqUserTo)
	if err != nil {
		json.NewEncoder(w).Encode("Invalid Argument userFrom")
	}

	amount64, err := strconv.ParseUint(reqAmount, 10, 32)
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
	} else {
		json.NewEncoder(w).Encode("transaction sucsessful")
	}
}
