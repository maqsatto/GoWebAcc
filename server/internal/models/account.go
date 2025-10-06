package models

import "time"

type Account struct {
	ID           uint          `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID       uint          `gorm:"not null" json:"user_id"`
	Name         string        `gorm:"type:varchar(100);not null" json:"name" validate:"required,min=2,max=100"`
	Balance      float64       `gorm:"type:numeric;default:0" json:"balance" validate:"gte=0"`
	Currency     string        `gorm:"type:varchar(10);default:'USD'" json:"currency" validate:"currency"`
	CreatedAt    time.Time     `json:"created_at"`
	Transactions []Transaction `gorm:"foreignKey:AccountID" json:"transactions"`
}
