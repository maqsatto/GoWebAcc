package controllers

import (
	"accountantapp/go-service/internal/models"
	"accountantapp/go-service/internal/services"
	"accountantapp/go-service/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {
	userID := c.GetUint("userID")

	var input models.Category
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid input data", err)
		return
	}

	category, err := services.CreateCategoryService(userID, &input)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Failed to create category", err)
		return
	}

	utils.Success(c, http.StatusCreated, "Category created successfully", category)
}

func GetAllCategories(c *gin.Context) {
	userID := c.GetUint("userID")

	categories, err := services.GetAllCategoriesService(userID)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to fetch categories", err)
		return
	}

	utils.Success(c, http.StatusOK, "Categories retrieved successfully", categories)
}

func GetCategoryByID(c *gin.Context) {
	userID := c.GetUint("userID")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid category ID", err)
		return
	}

	category, err := services.GetCategoryByIDService(userID, id)
	if err != nil {
		utils.Error(c, http.StatusNotFound, "Category not found", err)
		return
	}

	utils.Success(c, http.StatusOK, "Category retrieved successfully", category)
}

func UpdateCategory(c *gin.Context) {
	userID := c.GetUint("userID")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid category ID", err)
		return
	}

	var input models.Category
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid input data", err)
		return
	}

	category, err := services.UpdateCategoryService(userID, id, &input)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Failed to update category", err)
		return
	}

	utils.Success(c, http.StatusOK, "Category updated successfully", category)
}

func DeleteCategory(c *gin.Context) {
	userID := c.GetUint("userID")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid category ID", err)
		return
	}

	err = services.DeleteCategoryService(userID, id)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to delete category", err)
		return
	}

	utils.Success(c, http.StatusOK, "Category deleted successfully", nil)
}
