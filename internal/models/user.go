package models

type User struct {
	Id      int  `gorm:"column:id;primary_key" json:"Id"`
	Balance uint `gorm:"column:balance" json:"Balance"`
}

func (u User) GetId() int {
	return u.Id
}

func (u User) GetBalance() uint {
	return u.Balance
}
