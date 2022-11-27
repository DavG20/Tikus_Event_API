package auth

import authmodel "github.com/DavG20/Tikus_Event_Api/Internal/pkg/auth/Auth_Model"

type AuthServiceInter interface {
	CreateUser(*authmodel.AuthModel) *authmodel.DBResponse
	FindUserByUserName(userName string) (*authmodel.AuthModel, error)
	FindUserByEmail(email string )( *authmodel.AuthModel ,bool)
}
