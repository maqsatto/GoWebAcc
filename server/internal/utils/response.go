package utils

import "github.com/gin-gonic/gin"

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error,omitempty"`
}

func Success(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func Error(c *gin.Context, status int, message string, err error) {
	errorMsg := ""
	if err != nil {
		errorMsg = err.Error()
	}
	c.JSON(status, Response{
		Success: false,
		Message: message,
		Error:   errorMsg,
	})
}
