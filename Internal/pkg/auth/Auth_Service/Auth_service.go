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
func (authService *AuthService) CreateUser(user *authmodel.AuthModel) *authmodel.DBResponse {
	return authService.AuthRepo.CreateUser(user)
}
