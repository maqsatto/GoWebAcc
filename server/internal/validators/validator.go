package validators

import (
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func Init() {
	validate = validator.New()
	validate.RegisterValidation("currency", validateCurrency)
	validate.RegisterValidation("transaction_type", validateTransactionType)
}

func ValidateStruct(s interface{}) map[string]string {
	errs := make(map[string]string)

	if err := validate.Struct(s); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field := strings.ToLower(err.Field())
			switch err.Tag() {
			case "required":
				errs[field] = "This field is required"
			case "email":
				errs[field] = "Invalid email format"
			case "min":
				errs[field] = "Value is too short"
			case "max":
				errs[field] = "Value is too long"
			case "currency":
				errs[field] = "Invalid currency code"
			case "transaction_type":
				errs[field] = "Type must be 'income' or 'expense'"
			default:
				errs[field] = "Invalid value"
			}
		}
	}
	return errs
}

func validateCurrency(fl validator.FieldLevel) bool {
	currency := fl.Field().String()
	validCurrencies := map[string]bool{
		"USD": true, "EUR": true, "GBP": true, "RUB": true,
		"JPY": true, "CNY": true, "CAD": true, "AUD": true,
	}
	return validCurrencies[currency]
}

func validateTransactionType(fl validator.FieldLevel) bool {
	ttype := fl.Field().String()
	return ttype == "income" || ttype == "expense"
}

func IsValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(pattern, email)
	return matched
}
