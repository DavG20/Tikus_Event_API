package authrepo

import (
	"fmt"
	"time"

	constants "github.com/DavG20/Tikus_Event_Api/Internal/pkg/Entity/Constants"
	authmodel "github.com/DavG20/Tikus_Event_Api/Internal/pkg/auth/Auth_Model"

	"gorm.io/gorm"
)

type AuthRepo struct {
	DB *gorm.DB
}

func NewAuth(db *gorm.DB) AuthRepo {
	return AuthRepo{DB: db}
}

func (authRepo *AuthRepo) CreateUser(user *authmodel.AuthModel) (dbResponse *authmodel.DBResponse) {
	res := authRepo.DB.Table(constants.UserTableName).Create(&user)
	if res.RowsAffected == 0 {
		return nil
	}
	user.CreatedOn = time.Now()
	dbResponse = &authmodel.DBResponse{
		UserId:     user.UserId,
		UserName:   user.UserName,
		Email:      user.Email,
		CreatedOn:  user.CreatedOn,
		ProfileUrl: user.ProfileUrl,
		Admin:      user.Admin,
	}
	if dbResponse == nil {
		fmt.Println("error while decoding to dbresponse from input user")
		return nil
	}
	return dbResponse
}
