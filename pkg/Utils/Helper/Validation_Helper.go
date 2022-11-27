package helper

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func GetErrorMessage(fieldErr validator.FieldError) string {
	switch fieldErr.Tag() {
	case "required_without":
		return fmt.Sprintf("The Field %s is required if %s is not sapplied", fieldErr.Field(), fieldErr.Param())
	case "unique_without":
		return " must be unique, please try another one"
	case "min":
		return "The length must be greater than or equal to " + fieldErr.Param()
	case "email":
		return "ur email is not valid, please try to use the right email address"

	}
	return "Unknown error , "
}
