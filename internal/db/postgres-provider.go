package db

import (
	"KirsanovStavkaTV/internal/contracts"
	"KirsanovStavkaTV/internal/models"
	"errors"
	"fmt"
	"os"
	"time"

	postgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresProvider struct {
	DB *gorm.DB
}

func (p PostgresProvider) Provide() contracts.DatabaseProvider {
	dbHost := os.Getenv("POSTGRES_HOST")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbPort := os.Getenv("POSTGRES_PORT")
	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName)

	DB, err := p.connect(url)
	if err != nil {
		panic(err)
	}

	p.DB = DB

	return p
}

func (p PostgresProvider) FindUser(id int) models.User {
	userFrom := &models.User{}
	p.DB.Find(userFrom, id)

	return *userFrom
}

func (p PostgresProvider) GetUsers() []models.User {
	result := []models.User{}
	p.DB.Find(&result)
	return result
}

func (p PostgresProvider) MakeTransfer(t *models.Transaction) error {
	userFrom := p.FindUser(t.UserFrom)
	if userFrom.Id == 0 {
		return errors.New("Couldnt find UserTo transaction")
	}

	userTo := p.FindUser(t.UserTo)
	if userTo.Id == 0 {
		return errors.New("Couldnt find UserTo transaction")
	}

	userFromBalance := userFrom.GetBalance() - t.Amount
	userFrom.Balance = userFromBalance

	userToBalance := userTo.GetBalance() + t.Amount
	userTo.Balance = userToBalance

	err := p.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(t).Error; err != nil {
			return err
		}

		if err := tx.Save(&userFrom).Error; err != nil {
			return err
		}

		if err := tx.Save(&userTo).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (p PostgresProvider) connect(url string) (*gorm.DB, error) {
	for i := 0; i < 5; i++ {
		DB, err := gorm.Open(postgres.Open(url), &gorm.Config{})
		if err == nil {
			return DB, nil
		}
		time.Sleep(time.Second * 2)
	}
	return nil, errors.New("couldn't connect")
}
