package services

import (
	"accountantapp/go-service/internal/auth"
	"accountantapp/go-service/internal/database"
	"accountantapp/go-service/internal/models"
	"accountantapp/go-service/internal/validators"
	"errors"
	"fmt"
)

func CreateUserService(input *models.User) (*models.User, error) {
	// Валидация
	if errs := validators.ValidateStruct(input); len(errs) > 0 {
		return nil, fmt.Errorf("validation failed: %v", errs)
	}

	// Проверка email
	if !validators.IsValidEmail(input.Email) {
		return nil, errors.New("invalid email format")
	}

	// Проверка существования пользователя
	var existing models.User
	if err := database.DB.Where("email = ?", input.Email).First(&existing).Error; err == nil {
		return nil, errors.New("user with this email already exists")
	}

	// Хеширование пароля
	if err := input.HashPassword(); err != nil {
		return nil, errors.New("failed to hash password")
	}

	// Создание пользователя
	if err := database.DB.Create(input).Error; err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	// Убираем пароль из ответа
	input.Password = ""
	return input, nil
}

func LoginUserService(email, password string) (string, *models.User, error) {
	var user models.User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return "", nil, errors.New("invalid credentials")
	}

	if err := user.CheckPassword(password); err != nil {
		return "", nil, errors.New("invalid credentials")
	}

	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		return "", nil, errors.New("failed to generate token")
	}

	user.Password = ""
	return token, &user, nil
}

func GetAllUsersService() ([]models.User, error) {
	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
		return nil, errors.New("failed to fetch users")
	}

	// Убираем пароли
	for i := range users {
		users[i].Password = ""
	}
	return users, nil
}

func GetUserByIDService(id int) (*models.User, error) {
	var user models.User
	if err := database.DB.Preload("Accounts").Preload("Categories").First(&user, id).Error; err != nil {
		return nil, errors.New("user not found")
	}

	user.Password = ""
	return &user, nil
}

func UpdateUserService(id int, input *models.User) (*models.User, error) {
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return nil, errors.New("user not found")
	}

	// Валидация
	if errs := validators.ValidateStruct(input); len(errs) > 0 {
		return nil, fmt.Errorf("validation failed: %v", errs)
	}

	// Если обновляется пароль, хешируем его
	if input.Password != "" {
		if err := input.HashPassword(); err != nil {
			return nil, errors.New("failed to hash password")
		}
	}

	if err := database.DB.Model(&user).Updates(input).Error; err != nil {
		return nil, errors.New("failed to update user")
	}

	user.Password = ""
	return &user, nil
}

func DeleteUserService(id int) error {
	if err := database.DB.Delete(&models.User{}, id).Error; err != nil {
		return errors.New("failed to delete user")
	}
	return nil
}
