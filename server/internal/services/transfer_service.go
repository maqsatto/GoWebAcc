package services

import (
	"accountantapp/go-service/internal/database"
	"accountantapp/go-service/internal/models"
	"errors"
	"time"
)

func CreateTransfer(transfer *models.Transfer) (*models.Transfer, error) {
	// Проверяем, что аккаунты существуют
	var fromAccount, toAccount models.Account
	if err := database.DB.First(&fromAccount, transfer.FromAccountID).Error; err != nil {
		return nil, errors.New("from account not found")
	}
	if err := database.DB.First(&toAccount, transfer.ToAccountID).Error; err != nil {
		return nil, errors.New("to account not found")
	}

	// Проверяем, что это не перевод на тот же аккаунт
	if transfer.FromAccountID == transfer.ToAccountID {
		return nil, errors.New("cannot transfer to the same account")
	}

	// Проверяем достаточность средств
	if fromAccount.Balance < transfer.Amount {
		return nil, errors.New("insufficient funds")
	}

	// Устанавливаем дату по умолчанию
	if transfer.Date.IsZero() {
		transfer.Date = time.Now()
	}

	// Рассчитываем конвертированную сумму (если валюты разные)
	transfer.ConvertedAmount = transfer.Amount * transfer.ExchangeRate

	// Начинаем транзакцию
	tx := database.DB.Begin()

	// Создаем запись о переводе
	if err := tx.Create(transfer).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Обновляем балансы
	fromAccount.Balance -= transfer.Amount
	toAccount.Balance += transfer.ConvertedAmount

	if err := tx.Save(&fromAccount).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Save(&toAccount).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Создаем транзакции для истории
	fromTransaction := models.Transaction{
		UserID:      transfer.UserID,
		AccountID:   transfer.FromAccountID,
		Amount:      -transfer.Amount,
		Description: "Transfer to " + toAccount.Name + ": " + transfer.Description,
		CategoryID:  0, // Можно создать специальную категорию для трансферов
		Date:        transfer.Date,
	}

	toTransaction := models.Transaction{
		UserID:      transfer.UserID,
		AccountID:   transfer.ToAccountID,
		Amount:      transfer.ConvertedAmount,
		Description: "Transfer from " + fromAccount.Name + ": " + transfer.Description,
		CategoryID:  0,
		Date:        transfer.Date,
	}

	if err := tx.Create(&fromTransaction).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Create(&toTransaction).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return transfer, nil
}

func GetAllTransfersService() ([]models.Transfer, error) {
	var transfers []models.Transfer
	if err := database.DB.Preload("FromAccount").Preload("ToAccount").Find(&transfers).Error; err != nil {
		return nil, err
	}
	return transfers, nil
}

func GetTransferByIDService(id int) (*models.Transfer, error) {
	var transfer models.Transfer
	if err := database.DB.Preload("FromAccount").Preload("ToAccount").First(&transfer, id).Error; err != nil {
		return nil, errors.New("transfer not found")
	}
	return &transfer, nil
}

func GetTransfersByAccountID(accountID int) ([]models.Transfer, error) {
	var transfers []models.Transfer
	if err := database.DB.Preload("FromAccount").Preload("ToAccount").
		Where("from_account_id = ? OR to_account_id = ?", accountID, accountID).
		Find(&transfers).Error; err != nil {
		return nil, err
	}
	return transfers, nil
}

func DeleteTransfer(id int) error {
	var transfer models.Transfer
	if err := database.DB.First(&transfer, id).Error; err != nil {
		return errors.New("transfer not found")
	}

	tx := database.DB.Begin()

	// Возвращаем средства
	var fromAccount, toAccount models.Account
	database.DB.First(&fromAccount, transfer.FromAccountID)
	database.DB.First(&toAccount, transfer.ToAccountID)

	fromAccount.Balance += transfer.Amount
	toAccount.Balance -= transfer.ConvertedAmount

	if err := tx.Save(&fromAccount).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Save(&toAccount).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Удаляем связанные транзакции
	if err := tx.Where("description LIKE ?", "%Transfer%").Delete(&models.Transaction{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Delete(&transfer).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
