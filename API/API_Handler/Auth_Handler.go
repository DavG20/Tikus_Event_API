package apihandler

import (
	"errors"
	"fmt"
	"net/http"

	authmodel "github.com/DavG20/Tikus_Event_Api/Internal/pkg/auth/Auth_Model"
	authservice "github.com/DavG20/Tikus_Event_Api/Internal/pkg/auth/Auth_Service"
	"github.com/gin-gonic/gin"
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
	userInput := authmodel.UserInput{}
	if err := context.ShouldBind(&userInput); err != nil {
		fmt.Println("eror in binding")
		context.JSON(http.StatusBadRequest, gin.H{"error": "some inputs are already exist in our db"})
		return

	}
	fmt.Println(userInput)
	context.BindJSON(&userInput)

	_, isExist := authHandler.authService.FindUserByUserName(userInput.UserName)
	if isExist {
		context.JSON(http.StatusBadRequest, gin.H{"error": "User already registred"})
		return
	}

	dbresponse, err := authHandler.authService.CreateUser(&userInput)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": errors.New("internal server error")})
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
