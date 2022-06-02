package web

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type responseSuccess struct {
	Data interface{} `json:"data,omitempty"`
}

type responseError struct {
	Status  string `json:"status"`  // ex: "404 - Not Found"
	Message string `json:"message"` // ex: "could not found user"
}

func response(c *gin.Context, status int, data interface{}) {
	c.JSON(status, data)
}

func Success(c *gin.Context, status int, data interface{}) {
	response(c, status, data)
}

func Failure(c *gin.Context, statusCode int, message string) {
	code := fmt.Sprintf("%d - %s", statusCode, http.StatusText(statusCode))
	err := responseError{
		Status:  code,
		Message: message,
	}
	response(c, statusCode, err)
}
