package apihandler

import (
	"errors"
	"fmt"
	"net/http"

	sessionjwt "github.com/DavG20/Tikus_Event_Api/Internal/pkg/Session_JWT"
	authmodel "github.com/DavG20/Tikus_Event_Api/Internal/pkg/auth/Auth_Model"
	authservice "github.com/DavG20/Tikus_Event_Api/Internal/pkg/auth/Auth_Service"
	helper "github.com/DavG20/Tikus_Event_Api/pkg/Utils/Helper"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	authService   authservice.AuthService
	CookieHandler sessionjwt.CookieHandler
}

func NewauthHandler(authService authservice.AuthService) AuthHandler {
	return AuthHandler{
		authService: authService,
	}
}

func (authHandler *AuthHandler) RegisterHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		// session:=sessions.Default(context)
		// userName=session.Get(constants.UserKey)

		var UserInput authmodel.UserRegisterInput
		ResponseMessage := authmodel.ResponseMessage{}
		session := sessionjwt.Session{}
		// session, isvalid := authHandler.CookieHandler.ValidateCookie(context)
		// if isvalid {
		// 	ResponseMessage.Message = "u r already logged in "
		// 	ResponseMessage.Success = false
		// 	context.JSON(http.StatusUnauthorized, ResponseMessage)
		// 	return
		// }
		if err := context.ShouldBindJSON(&UserInput); err != nil {
			var valErr validator.ValidationErrors
			if errors.As(err, &valErr) {
				fmt.Println("error")
				ErrMessage := make([]authmodel.ResponseMessage, len(valErr))
				for i, fieldErr := range valErr {
					ErrMessage[i] = authmodel.ResponseMessage{Message: helper.GetErrorMessage(fieldErr), Success: false, InvalidField: fieldErr.Field()}
				}
				fmt.Println("eroro in binding json")
				context.JSON(http.StatusBadRequest, gin.H{"errors": ErrMessage})
				return
			}
			context.JSON(http.StatusBadRequest, gin.H{"error": "unknown input error"})
			return
		}

		_, isUserNameExist := authHandler.authService.FindUserByUserName(UserInput.UserName)
		if isUserNameExist {
			ResponseMessage.Message = "username taken please use aonther one"
			ResponseMessage.Success = false
			ResponseMessage.InvalidField = "username"
			context.JSON(http.StatusBadRequest, ResponseMessage)
			return
		}

		_, isEmailExist := authHandler.authService.FindUserByEmail(UserInput.Email)
		if isEmailExist {
			ResponseMessage.Message = "Email is already registerd, please try another one"
			ResponseMessage.Success = false
			ResponseMessage.InvalidField = "email"
			context.JSON(http.StatusBadRequest, ResponseMessage)
			return
		}

		user, err := authHandler.authService.CreateUser(&UserInput)
		if err != nil {
			ResponseMessage.Message = "internal server error , please try again"
			ResponseMessage.Success = false
			context.JSON(http.StatusInternalServerError, ResponseMessage)
			return
		}
		fmt.Println(session, "session")
		session.UserName = user.UserName
		tokenString, err := authHandler.CookieHandler.CreateCookie(&session)
		if err != nil {
			fmt.Println("error while creating token ")
			ResponseMessage.Message = "failed to save cookies"
			ResponseMessage.Success = false
			context.JSON(http.StatusInternalServerError, ResponseMessage)
			return
		}
		fmt.Println(tokenString)
		// context.Writer.Header().Set("token", tokenString)
		context.SetCookie(
			"token",
			tokenString,
			3600,
			"/user",
			"localhost",
			false,
			true,
		)
		context.JSON(http.StatusOK, user)

	}
}
func (authHandler *AuthHandler) SearchUser(context *gin.Context) {

	context.JSON(200, gin.H{"message": "try"})
}

func (authHandler *AuthHandler) AuthRequired() gin.HandlerFunc {
	return func(context *gin.Context) {
		session, isValid := authHandler.CookieHandler.ValidateCookie(context)
		fmt.Println(isValid)
		if !isValid {
			context.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized user"})
			context.Abort()
			return
		}
		if session.UserName == "" {
			context.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized user"})
			context.Abort()
			return
		}
		context.Next()
	}
}
