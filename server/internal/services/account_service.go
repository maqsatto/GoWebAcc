package services

import (
	"accountantapp/go-service/internal/database"
	"accountantapp/go-service/internal/models"
	"accountantapp/go-service/internal/validators"
	"errors"
	"fmt"
)

func CreateAccountService(userID uint, acc *models.Account) (*models.Account, error) {
	// Валидация
	if errs := validators.ValidateStruct(acc); len(errs) > 0 {
		return nil, fmt.Errorf("validation failed: %v", errs)
	}

	// Устанавливаем владельца
	acc.UserID = userID

	if err := database.DB.Create(acc).Error; err != nil {
		return nil, fmt.Errorf("failed to create account: %v", err)
	}

	return acc, nil
}

func GetAllAccountsService(userID uint) ([]models.Account, error) {
	var accounts []models.Account
	if err := database.DB.Where("user_id = ?", userID).Find(&accounts).Error; err != nil {
		return nil, errors.New("failed to fetch accounts")
	}
	return accounts, nil
}

func GetAccountByIDService(userID uint, id int) (*models.Account, error) {
	var acc models.Account
	if err := database.DB.Where("user_id = ? AND id = ?", userID, id).First(&acc).Error; err != nil {
		return nil, errors.New("account not found")
	}
	return &acc, nil
}

func UpdateAccountService(userID uint, id int, input *models.Account) (*models.Account, error) {
	var acc models.Account
	if err := database.DB.Where("user_id = ? AND id = ?", userID, id).First(&acc).Error; err != nil {
		return nil, errors.New("account not found")
	}

	// Валидация
	if errs := validators.ValidateStruct(input); len(errs) > 0 {
		return nil, fmt.Errorf("validation failed: %v", errs)
	}

	acc.Name = input.Name
	acc.Balance = input.Balance
	acc.Currency = input.Currency

	if err := database.DB.Save(&acc).Error; err != nil {
		return nil, errors.New("failed to update account")
	}

	return &acc, nil
}

func DeleteAccountService(userID uint, id int) error {
	if err := database.DB.Where("user_id = ? AND id = ?", userID, id).Delete(&models.Account{}).Error; err != nil {
		return errors.New("failed to delete account")
	}
	return nil
}
