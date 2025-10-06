package controllers

import (
	"accountantapp/go-service/internal/models"
	"accountantapp/go-service/internal/services"
	"accountantapp/go-service/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func Register(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid input", err)
		return
	}

	user, err := services.CreateUserService(&input)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Registration failed", err)
		return
	}

	utils.Success(c, http.StatusCreated, "User registered successfully", user)
}

func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid input", err)
		return
	}

	token, user, err := services.LoginUserService(input.Email, input.Password)
	if err != nil {
		utils.Error(c, http.StatusUnauthorized, "Login failed", err)
		return
	}

	utils.Success(c, http.StatusOK, "Login successful", gin.H{
		"token": token,
		"user":  user,
	})
}

func GetProfile(c *gin.Context) {
	userID, _ := c.Get("userID")

	user, err := services.GetUserByIDService(int(userID.(uint)))
	if err != nil {
		utils.Error(c, http.StatusNotFound, "User not found", err)
		return
	}

	utils.Success(c, http.StatusOK, "Profile retrieved successfully", user)
}
