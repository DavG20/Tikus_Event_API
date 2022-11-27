package authservice

import (
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
func (authService *AuthService) CreateUser(user *authmodel.UserInput) (*authmodel.DBResponse, error) {
	return authService.AuthRepo.CreateUser(user)
}

func (authService *AuthService) FindUserByUserName(userName string) (user *authmodel.AuthModel, state bool) {
	user, err := authService.AuthRepo.FindUserByUserName(userName)
	if err != nil {
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
