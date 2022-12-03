package authservice

import (
	"fmt"

	authmodel "github.com/DavG20/Tikus_Event_Api/Internal/pkg/auth/Auth_Model"
	authrepo "github.com/DavG20/Tikus_Event_Api/Internal/pkg/auth/Auth_Repo"
)

type AuthService struct {
	AuthRepo authrepo.AuthRepo
}

func NewAuthService(authRepo authrepo.AuthRepo) AuthService {
	return AuthService{
		AuthRepo: authRepo,
	}
}
func (authService *AuthService) CreateUser(user *authmodel.UserRegisterInput) (*authmodel.DBResponse, error) {
	return authService.AuthRepo.CreateUser(user)
}

func (authService *AuthService) FindUserByUserName(userName string) (user *authmodel.AuthModel, state bool) {
	user, err := authService.AuthRepo.FindUserByUserName(userName)
	if err != nil {
		fmt.Println("error")
		return nil, false
	}

	// if user == nil {
	// 	return nil, false
	// }
	return user, true

}

func (authService *AuthService) FindUserByEmail(email string) (*authmodel.AuthModel, bool) {
	user, err := authService.AuthRepo.FindUserByEmail(email)
	if err != nil {
		return nil, false
	}
	return user, true
}

func (authService *AuthService) GetDbResponse(user *authmodel.AuthModel) (*authmodel.DBResponse, bool) {
	dbResponse, err := authService.AuthRepo.GetDbResponse(user)
	if err != nil {
		return nil, false
	}
	return dbResponse, true
}

func (authService *AuthService) DeleteAccount(userName string) bool {
	err := authService.AuthRepo.DeleteAccount(userName)
	if err != nil {
		return false
	}
	return true
}

func (authService *AuthService) ChangePassword(userName, newPassword string) bool {
	return authService.AuthRepo.ChangePassword(userName, newPassword)
}
