package models

import "time"

type Category struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name" validate:"required,min=2,max=100"`
	Type      string    `gorm:"type:varchar(50);not null" json:"type" validate:"required,transaction_type"`
	Color     string    `gorm:"type:varchar(7);default:'#007bff'" json:"color"`
	Icon      string    `gorm:"type:varchar(50)" json:"icon"`
	CreatedAt time.Time `json:"created_at"`
}
