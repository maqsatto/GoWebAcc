package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           uint          `gorm:"primaryKey;autoIncrement" json:"id"`
	Name         string        `gorm:"type:varchar(100);not null" json:"name" validate:"required,min=2,max=100"`
	Email        string        `gorm:"uniqueIndex;type:varchar(100);not null" json:"email" validate:"required,email"`
	Password     string        `gorm:"type:varchar(255);not null" json:"password" validate:"required,min=6"`
	CreatedAt    time.Time     `json:"created_at"`
	Accounts     []Account     `gorm:"foreignKey:UserID" json:"accounts"`
	Transactions []Transaction `gorm:"foreignKey:UserID" json:"transactions"`
	Categories   []Category    `gorm:"foreignKey:UserID" json:"categories"`
	Templates    []Template    `gorm:"foreignKey:UserID" json:"templates"`
	Settings     []AppSettings `gorm:"foreignKey:UserID" json:"settings"`
}

// HashPassword хеширует пароль
func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// CheckPassword проверяет пароль
func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
