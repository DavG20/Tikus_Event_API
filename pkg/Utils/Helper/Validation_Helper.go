package helper

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func GetErrorMessage(fieldErr validator.FieldError) string {
	switch fieldErr.Tag() {
	case "required":
		return fmt.Sprintf("The  %s is required , please write ur %s correctly ", fieldErr.Field(), fieldErr.Field())
	case "unique":
		return fmt.Sprintf("%s must be unique, please try another one", fieldErr.Field())
	case "min":
		return "The length must be greater than or equal to " + fieldErr.Param()
	case "email":
		return "ur email is not valid, please try to use the right email address"

	}
	return "Unknown error , "
}
