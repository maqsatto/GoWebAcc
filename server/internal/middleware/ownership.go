package middleware

import (
	"accountantapp/go-service/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Простые middleware для проверки ID (опционально, можно удалить если не нужны)
func ValidateAccountID() gin.HandlerFunc {
	return validateID("account")
}

func ValidateCategoryID() gin.HandlerFunc {
	return validateID("category")
}

func ValidateTransactionID() gin.HandlerFunc {
	return validateID("transaction")
}

func ValidateTemplateID() gin.HandlerFunc {
	return validateID("template")
}

func ValidateSettingID() gin.HandlerFunc {
	return validateID("setting")
}

func ValidateTransferID() gin.HandlerFunc {
	return validateID("transfer")
}

func validateID(resourceType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			utils.Error(c, http.StatusBadRequest, "Invalid "+resourceType+" ID", err)
			c.Abort()
			return
		}
		c.Next()
	}
}
