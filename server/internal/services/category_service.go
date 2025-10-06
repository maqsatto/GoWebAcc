package services

import (
	"accountantapp/go-service/internal/database"
	"accountantapp/go-service/internal/models"
	"accountantapp/go-service/internal/validators"
	"errors"
	"fmt"
)

func CreateCategoryService(userID uint, cat *models.Category) (*models.Category, error) {
	if errs := validators.ValidateStruct(cat); len(errs) > 0 {
		return nil, fmt.Errorf("validation failed: %v", errs)
	}

	cat.UserID = userID

	if err := database.DB.Create(cat).Error; err != nil {
		return nil, fmt.Errorf("failed to create category: %v", err)
	}

	return cat, nil
}

func GetAllCategoriesService(userID uint) ([]models.Category, error) {
	var cats []models.Category
	if err := database.DB.Where("user_id = ?", userID).Find(&cats).Error; err != nil {
		return nil, errors.New("failed to fetch categories")
	}
	return cats, nil
}

func GetCategoryByIDService(userID uint, id int) (*models.Category, error) {
	var cat models.Category
	if err := database.DB.Where("user_id = ? AND id = ?", userID, id).First(&cat).Error; err != nil {
		return nil, errors.New("category not found")
	}
	return &cat, nil
}

func UpdateCategoryService(userID uint, id int, input *models.Category) (*models.Category, error) {
	var cat models.Category
	if err := database.DB.Where("user_id = ? AND id = ?", userID, id).First(&cat).Error; err != nil {
		return nil, errors.New("category not found")
	}

	if errs := validators.ValidateStruct(input); len(errs) > 0 {
		return nil, fmt.Errorf("validation failed: %v", errs)
	}

	cat.Name = input.Name
	cat.Type = input.Type
	cat.Color = input.Color
	cat.Icon = input.Icon

	if err := database.DB.Save(&cat).Error; err != nil {
		return nil, errors.New("failed to update category")
	}

	return &cat, nil
}

func DeleteCategoryService(userID uint, id int) error {
	if err := database.DB.Where("user_id = ? AND id = ?", userID, id).Delete(&models.Category{}).Error; err != nil {
		return errors.New("failed to delete category")
	}
	return nil
}
