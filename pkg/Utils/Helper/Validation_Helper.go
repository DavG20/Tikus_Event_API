package helper

import (
	"errors"
	"fmt"

	authmodel "github.com/DavG20/Tikus_Event_Api/Internal/pkg/auth/Auth_Model"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetErrorMessage(fieldErr validator.FieldError) string {
	switch fieldErr.Tag() {
	case "required":
		return fmt.Sprintf("The  %s is required , please write ur %s correctly ", fieldErr.Field(), fieldErr.Field())
	case "unique":
		return fmt.Sprintf("%s must be unique, please try another one", fieldErr.Field())
	case "min":
		return fmt.Sprintf("The length of %s must be greater than or equal to %s ", fieldErr.Field(), fieldErr.Param())
	case "email":
		return "ur email is not valid, please try to use the right email address"

	}
	return "Unknown error , "
}

func ShouldBindJSONHelper(context *gin.Context, jsonObject interface{}) []authmodel.ResponseMessage {
	errMess := []authmodel.ResponseMessage{}
	fmt.Println("jsonobject", jsonObject)
	if err := context.ShouldBindJSON(&jsonObject); err != nil {
		var valErr validator.ValidationErrors
		if errors.As(err, &valErr) {
			errMessage := make([]authmodel.ResponseMessage, len(valErr))
			for i, fieldErr := range valErr {
				errMessage[i] = authmodel.ResponseMessage{Message: GetErrorMessage(fieldErr), Success: false, InvalidField: fieldErr.Field()}
			}
			return errMessage
		}
		errMess = append(errMess, authmodel.ResponseMessage{Message: "invalid input ,please try to use the right input", Success: false})
		return errMess

	}
	fmt.Println(jsonObject, "after")
	return nil
}
