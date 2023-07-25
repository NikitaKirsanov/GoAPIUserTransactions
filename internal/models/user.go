package models

import "time"

type User struct {
	Id        int        `gorm:"column:id;primary_key" json:"Id"`
	Balance   uint       `gorm:"column:balance" json:"Balance"`
	CreatedAt *time.Time `gorm:"column:created_at;" json:"CreatedAt"`
}

func (u User) GetId() int {
	return u.Id
}

func (u User) GetBalance() uint {
	return u.Balance
}

func (u User) SetBalance(newBalace uint) {
	u.Balance = newBalace
}
