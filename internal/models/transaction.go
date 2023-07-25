package models

import "time"

type Transaction struct {
	Id        int        `gorm:"column:id;primary_key" json:"Id"`
	UserFrom  int        `gorm:"column:user_from;foreignKey" json:"UserFrom"`
	UserTo    int        `gorm:"column:user_to;foreignKey" json:"UserTo"`
	Amount    uint       `gorm:"column:amount" json:"amount"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"CreatedAt"`
}
