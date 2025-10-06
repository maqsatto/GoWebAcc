package controllers

import (
	"accountantapp/go-service/internal/models"
	"accountantapp/go-service/internal/services"
	"accountantapp/go-service/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateAccount(c *gin.Context) {
	userID := c.GetUint("userID")

	var input models.Account
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid input data", err)
		return
	}

	account, err := services.CreateAccountService(userID, &input)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Failed to create account", err)
		return
	}

	utils.Success(c, http.StatusCreated, "Account created successfully", account)
}

func GetAllAccounts(c *gin.Context) {
	userID := c.GetUint("userID")

	accounts, err := services.GetAllAccountsService(userID)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to fetch accounts", err)
		return
	}

	utils.Success(c, http.StatusOK, "Accounts retrieved successfully", accounts)
}

func GetAccountByID(c *gin.Context) {
	userID := c.GetUint("userID")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid account ID", err)
		return
	}

	account, err := services.GetAccountByIDService(userID, id)
	if err != nil {
		utils.Error(c, http.StatusNotFound, "Account not found", err)
		return
	}

	utils.Success(c, http.StatusOK, "Account retrieved successfully", account)
}

func UpdateAccount(c *gin.Context) {
	userID := c.GetUint("userID")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid account ID", err)
		return
	}

	var input models.Account
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid input data", err)
		return
	}

	account, err := services.UpdateAccountService(userID, id, &input)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Failed to update account", err)
		return
	}

	utils.Success(c, http.StatusOK, "Account updated successfully", account)
}

func DeleteAccount(c *gin.Context) {
	userID := c.GetUint("userID")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid account ID", err)
		return
	}

	err = services.DeleteAccountService(userID, id)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to delete account", err)
		return
	}

	utils.Success(c, http.StatusOK, "Account deleted successfully", nil)
}
