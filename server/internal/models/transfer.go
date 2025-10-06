package models

import "time"

type Transfer struct {
	ID              uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID          uint      `gorm:"not null" json:"user_id"`
	FromAccountID   uint      `gorm:"not null" json:"from_account_id" validate:"required"`
	ToAccountID     uint      `gorm:"not null" json:"to_account_id" validate:"required"`
	Amount          float64   `gorm:"type:numeric;not null" json:"amount" validate:"required,gt=0"`
	Description     string    `gorm:"type:varchar(255)" json:"description" validate:"max=255"`
	Currency        string    `gorm:"type:varchar(10);default:'USD'" json:"currency" validate:"currency"`
	ExchangeRate    float64   `gorm:"type:numeric;default:1" json:"exchange_rate" validate:"gt=0"`
	ConvertedAmount float64   `gorm:"type:numeric" json:"converted_amount"`
	Date            time.Time `gorm:"not null" json:"date" validate:"required"`
	CreatedAt       time.Time `json:"created_at"`

	FromAccount Account `gorm:"foreignKey:FromAccountID" json:"from_account"`
	ToAccount   Account `gorm:"foreignKey:ToAccountID" json:"to_account"`
}
