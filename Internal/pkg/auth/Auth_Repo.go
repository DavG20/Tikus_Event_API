package auth

import authmodel "github.com/DavG20/Tikus_Event_Api/Internal/pkg/auth/Auth_Model"

type AuthRepoInter interface {
	CreateUser(*authmodel.AuthModel) *authmodel.DBResponse
}
