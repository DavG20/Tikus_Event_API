package apihandler

import (
	"errors"
	"fmt"
	"net/http"

	authmodel "github.com/DavG20/Tikus_Event_Api/Internal/pkg/auth/Auth_Model"
	authservice "github.com/DavG20/Tikus_Event_Api/Internal/pkg/auth/Auth_Service"
	helper "github.com/DavG20/Tikus_Event_Api/pkg/Utils/Helper"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	authService authservice.AuthService
}

func NewauthHandler(authService authservice.AuthService) AuthHandler {
	return AuthHandler{
		authService: authService,
	}
}

func (authHandler *AuthHandler) CreateUserHandler(context *gin.Context) {
	var userInput authmodel.UserInput
	RessponseMessage := authmodel.ResponseMessage{}
	if err := context.ShouldBind(&userInput); err != nil {

		var valError validator.ValidationErrors
		if errors.As(err, &valError) {
			errMesssage := make([]authmodel.ResponseMessage, len(valError))
			for i, fieldErr := range valError {
				errMesssage[i] = authmodel.ResponseMessage{Message: helper.GetErrorMessage(fieldErr), Success: false, InvalidField: fieldErr.Field()}
			}
			context.JSON(http.StatusBadRequest, gin.H{"error": errMesssage})
			return
		}

		RessponseMessage.Message = "something went wrong"
		RessponseMessage.Success = false

		context.JSON(http.StatusBadRequest, RessponseMessage)
		return

	}

	context.BindJSON(&userInput)

	_, isUserNameExist := authHandler.authService.FindUserByUserName(userInput.UserName)
	fmt.Println(isUserNameExist, "chcked")
	if isUserNameExist {
		RessponseMessage.Message = "username aleardy taken , please provide another one"
		RessponseMessage.Success = false
		RessponseMessage.InvalidField = "UserName"
		context.JSON(http.StatusBadRequest, RessponseMessage)
		return
	}

	_, isEmailExist := authHandler.authService.FindUserByEmail(userInput.Email)
	if isEmailExist {
		RessponseMessage.Message = "email already registerd , use another one"
		RessponseMessage.Success = false
		RessponseMessage.InvalidField = "email"
		context.JSON(http.StatusBadRequest, RessponseMessage)
		return
	}

	dbresponse, err := authHandler.authService.CreateUser(&userInput)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": errors.New("internal server error")})
		return
	}
	fmt.Println(dbresponse)
	context.JSON(http.StatusOK, dbresponse)

}

func (authHandler *AuthHandler) Checkuser(cxt *gin.Context) {
	input := struct {
		Username string `json:"user_name"`
	}{}
	cxt.BindJSON(&input)
	user, stat := authHandler.authService.FindUserByUserName(input.Username)
	if !stat {
		cxt.JSON(http.StatusExpectationFailed, "failed")
	}
	cxt.JSON(http.StatusOK, user)
}
