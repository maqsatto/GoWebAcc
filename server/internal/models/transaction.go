package models

import "time"

type Transaction struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID      uint      `gorm:"not null" json:"user_id"`
	AccountID   uint      `gorm:"not null" json:"account_id" validate:"required"`
	Amount      float64   `gorm:"type:numeric;not null" json:"amount" validate:"required,gt=0"`
	Description string    `gorm:"type:varchar(255)" json:"description" validate:"max=255"`
	CategoryID  uint      `json:"category_id" validate:"required"`
	Type        string    `gorm:"type:varchar(20);not null" json:"type" validate:"required,transaction_type"`
	Date        time.Time `gorm:"not null" json:"date" validate:"required"`
	CreatedAt   time.Time `json:"created_at"`
}
