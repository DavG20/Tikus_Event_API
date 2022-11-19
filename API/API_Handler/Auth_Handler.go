package apihandler

import (
	"fmt"
	"net/http"

	authmodel "github.com/DavG20/Tikus_Event_Api/Internal/pkg/auth/Auth_Model"
	authservice "github.com/DavG20/Tikus_Event_Api/Internal/pkg/auth/Auth_Service"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	AuthService authservice.AuthService
}

func NewauthHandler(authService authservice.AuthService) AuthHandler {
	return AuthHandler{
		AuthService: authService,
	}
}

func (authHandler *AuthHandler) CreateUserHandler(context *gin.Context) {

	user := authmodel.AuthModel{}
	context.BindJSON(&user)
	dbresponse := authHandler.AuthService.CreateUser(&user)
	fmt.Println(dbresponse)
	context.JSON(http.StatusOK, dbresponse)

}
