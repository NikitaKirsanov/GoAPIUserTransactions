package contracts

import "KirsanovStavkaTV/internal/models"

type DatabaseProvider interface {
	Provide() DatabaseProvider
	FindUser(int) models.User
	GetUsers() []models.User
	MakeTransfer(*models.Transaction) error
}
