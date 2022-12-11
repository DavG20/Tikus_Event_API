package apihandler

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

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

		var UserInput authmodel.UserRegisterInput
		ResponseMessage := authmodel.ResponseMessage{}
		session := sessionjwt.Session{}

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
		//set cookie for the next access
		// the helerp function set the cookies
		helper.SetCookie(tokenString, context)

		context.JSON(http.StatusOK, user)

	}
}

// login handler using username and password as an input
func (authHandler *AuthHandler) LoginHandler(context *gin.Context) {
	userInput := authmodel.UserLoginInput{}
	responseMessage := authmodel.ResponseMessage{}
	session := sessionjwt.Session{}

	// check is the user is already logged in
	_, isValid := authHandler.CookieHandler.ValidateCookie(context)
	if isValid {
		responseMessage.Message = "you are logged in already "
		responseMessage.Success = false
		context.JSON(http.StatusNotAcceptable, responseMessage)
		return
	}
	if err := context.ShouldBindJSON(&userInput); err != nil {
		var valErr validator.ValidationErrors
		if errors.As(err, &valErr) {
			errMessage := make([]authmodel.ResponseMessage, len(valErr))
			for i, fieldErr := range valErr {
				errMessage[i] = authmodel.ResponseMessage{Message: helper.GetErrorMessage(fieldErr), Success: false, InvalidField: fieldErr.Field()}
			}
			context.JSON(400, errMessage)
			return
		}
		context.JSON(http.StatusBadRequest, gin.H{"errors": "bad input"})
		return
	}
	// check if the client exist in our db
	// if the user not registered then tell the user to register
	user, isUserNameExist := authHandler.authService.FindUserByUserName(userInput.UserName)
	if !isUserNameExist {
		responseMessage.Message = "username doesn't exist , please use the correct username, or register first"
		responseMessage.Success = false
		context.JSON(401, responseMessage)
		return
	}
	// if the user is registered and the cookie is valid the check the password
	isPasswordCorrect := helper.ComparePassword(user.Password, userInput.Password)
	if !isPasswordCorrect {
		responseMessage.Message = " password is not correct"
		responseMessage.Success = false
		context.JSON(401, responseMessage)
		return
	}
	// the user is legal to get the service
	// set user session
	// set username for the session for the next use in the cookie
	session.UserName = userInput.UserName
	// generate JWT token for the user for authentication purpose
	tokenString, err := authHandler.CookieHandler.CreateCookie(&session)
	// check is tokenstring is valid
	if err != nil {
		responseMessage.Message = "failed to login due to internal server error"
		responseMessage.Success = false
		context.JSON(http.StatusInternalServerError, responseMessage)
		return
	}
	// get secured User response
	UserRes, Success := authHandler.authService.GetDbResponse(user)
	if !Success {
		responseMessage.Message = "can't get user information , internal error"
		responseMessage.Success = false
		context.JSON(http.StatusInternalServerError, responseMessage)
		return
	}

	//set cookie
	helper.SetCookie(tokenString, context)
	context.JSON(http.StatusOK, UserRes)

}

// logout Hander
func (authHandler *AuthHandler) LogoutHandler(context *gin.Context) {

	responseMessage := authmodel.ResponseMessage{}

	helper.RemoveCookie(context)
	responseMessage.Message = "logout Succesfuly"
	responseMessage.Success = true
	context.JSON(http.StatusOK, responseMessage)

}

func (authHandler *AuthHandler) SearchUser(context *gin.Context) {

	context.JSON(200, gin.H{"message": "try"})
}

func (authHandler *AuthHandler) AuthRequired() gin.HandlerFunc {
	return func(context *gin.Context) {
		session, isValid := authHandler.CookieHandler.ValidateCookie(context)
		responseMessage := &authmodel.ResponseMessage{Message: "unathorized User , access denied, logged in first :)", Success: false}
		fmt.Println(isValid)
		if !isValid {
			context.JSON(http.StatusUnauthorized, responseMessage)
			context.Abort()
			return
		}
		if session.UserName == "" {
			fmt.Println("inher")
			context.JSON(http.StatusUnauthorized, responseMessage)
			context.Abort()
			return
		}
		_, isUserNameExist := authHandler.authService.FindUserByUserName(session.UserName)
		if !isUserNameExist {
			responseMessage.Message = "you are not the right user , or username not found"
			context.JSON(http.StatusUnauthorized, responseMessage)
			context.Abort()
			return
		}
		context.Next()
	}
}

func (authHandler *AuthHandler) DeleteAccount(context *gin.Context) {
	responseMessage := authmodel.ResponseMessage{}
	UserInput := &struct {
		Password string `json:"password" binding:"required,min=8"`
	}{}
	if err := context.ShouldBindJSON(&UserInput); err != nil {
		responseMessage.Message = "The password is not correct or empty"
		responseMessage.Success = false
		context.JSON(http.StatusUnauthorized, responseMessage)
		return
	}
	// I can check is the user has a valid cookie or not but the authHandler func already check this so i don't need the validity check of the cookie
	// get the session from the context and get the user username and then check if the user is valid,

	session, isValid := authHandler.CookieHandler.ValidateCookie(context)
	if !isValid {
		// i can igonore this part , because the authhandler function already check the cookie validty check
		responseMessage.Message = "unauthorized user"
		responseMessage.Success = false
		context.JSON(http.StatusUnauthorized, responseMessage)
		return
	}

	// check if  user is found in our db
	user, isUserNameExist := authHandler.authService.FindUserByUserName(session.UserName)

	// if user not found, means the user has no right to delete the account
	if !isUserNameExist {
		responseMessage.Message = "can't delete account , "
		responseMessage.Success = false
		context.JSON(http.StatusUnauthorized, responseMessage)
		return
	}

	// check is the password is correct , if the password is not correct abort the process
	isPasswordCorrect := helper.ComparePassword(user.Password, UserInput.Password)
	if !isPasswordCorrect {
		responseMessage.Message = "the password is incorrect"
		responseMessage.Success = false
		context.JSON(http.StatusUnauthorized, responseMessage)
		return
	}

	// everything went well , so delete the account
	isUserDeleted := authHandler.authService.DeleteAccount(user.UserName)
	if !isUserDeleted {
		responseMessage.Message = "internale server error, can 't delete account "
		responseMessage.Success = false
		context.JSON(http.StatusInternalServerError, responseMessage)
		return
	}
	// the response or let the user know the account is deleted
	// and delete the cookie
	helper.RemoveCookie(context)
	responseMessage.Message = "account deleted succesfuly"
	responseMessage.Success = true
	context.JSON(http.StatusOK, responseMessage)

}

func (authHandler *AuthHandler) GetUserInfo(context *gin.Context) {
	responseMessage := authmodel.ResponseMessage{}

	// the auth func check whether or not the user is valid
	// since the user not provideing any input i have to get the session to get user's userName
	// and check is the username is correct and found in our db
	session, isValid := authHandler.CookieHandler.ValidateCookie(context)
	// incase check validty
	if !isValid {
		responseMessage.Message = "unathorized user , access denied"
		responseMessage.Success = false
		context.JSON(http.StatusUnauthorized, responseMessage)
		return
	}
	// get user's info using user's username
	user, isUserNameExist := authHandler.authService.FindUserByUserName(session.UserName)
	if !isUserNameExist {
		responseMessage.Message = "not authorized to get this service , "
		responseMessage.Success = false
		context.JSON(http.StatusUnauthorized, responseMessage)
		return
	}
	// the user is exist response the user info
	UserInfo, Succes := authHandler.authService.GetDbResponse(user)
	// check incase some error happens
	if !Succes {
		responseMessage.Message = "internal server error , can't get user info"
		responseMessage.Success = false
		context.JSON(http.StatusInternalServerError, responseMessage)

	}

	context.JSON(http.StatusOK, UserInfo)
}

// change password handler

func (authHandler *AuthHandler) ChangePasswordHandler(context *gin.Context) {
	responseMessage := authmodel.ResponseMessage{}
	UserPasswordInput := &struct {
		OldPassword string `json:"password" binding:"required,min=8"`
		NewPassword string `json:"newpassword" binding:"required,min=8"`
	}{}
	// check  input validation
	if err := context.ShouldBindJSON(&UserPasswordInput); err != nil {
		var valErr validator.ValidationErrors
		if errors.As(err, &valErr) {
			errMessage := make([]authmodel.ResponseMessage, len(valErr))
			for i, fieldErr := range valErr {
				errMessage[i] = authmodel.ResponseMessage{Message: helper.GetErrorMessage(fieldErr), Success: false, InvalidField: fieldErr.Field()}
			}
			context.JSON(http.StatusBadRequest, errMessage)
			return
		}
		responseMessage.Message = "invalid input , please try to fill acording to the requirment"
		responseMessage.Success = false
		context.JSON(http.StatusBadRequest, responseMessage)
		return

	}

	// get the username from user's session or cookie
	session, isValidToken := authHandler.CookieHandler.ValidateCookie(context)
	// just incase check if the cookie is  valid , even if it is checked by the authrequired func
	if !isValidToken {
		responseMessage.Message = "unauthorized user to get this service"
		responseMessage.Success = false
		context.JSON(http.StatusUnauthorized, responseMessage)
		return
	}
	// check if the user is in our db , if the user is in the system then processed ...
	user, isUserExist := authHandler.authService.FindUserByUserName(session.UserName)
	if !isUserExist {
		// the username in the cookie is not from the one we sent during the user loggedin or register
		responseMessage.Message = "hacker huh :), nah not today "
		responseMessage.Success = false
		context.JSON(http.StatusUnauthorized, responseMessage)
		return
	}
	// let's check the old password  is correct
	isPasswordCorrect := helper.ComparePassword(user.Password, UserPasswordInput.OldPassword)
	if !isPasswordCorrect {
		responseMessage.Message = "The old password is not correct , please try again"
		responseMessage.Success = false
		context.JSON(http.StatusUnauthorized, responseMessage)
		return
	}
	hashNewPassword, err := helper.HashPassword(UserPasswordInput.NewPassword)
	if err != nil {
		responseMessage.Message = "internal problem , can't change password "
		responseMessage.Success = false
		context.JSON(http.StatusInternalServerError, responseMessage)
		return
	}

	// the final step let's change the password
	passwordChangedSuccess := authHandler.authService.ChangePassword(user.UserName, hashNewPassword)
	if !passwordChangedSuccess {
		// incase some problem happens and can't change password
		responseMessage.Message = "can't change password , internal problem occur"
		responseMessage.Success = false
		context.JSON(http.StatusInternalServerError, responseMessage)
		return
	}
	fmt.Println(hashNewPassword)

	responseMessage.Message = "password changed succesfuly"
	responseMessage.Success = true
	context.JSON(http.StatusOK, responseMessage)

}

// let's store profile pic in filesystem
func (authHandler *AuthHandler) UploadProfileHandler(context *gin.Context) {

	responseMessage := authmodel.ResponseMessage{}
	//   get user's username from the context  for the purpose of profile pic filename
	session, _ := authHandler.CookieHandler.ValidateCookie(context)
	file, header, err := context.Request.FormFile("profile")
	if err != nil {
		responseMessage.Message = "please choose the right picture for profile"
		responseMessage.Success = false
		context.JSON(http.StatusBadRequest, responseMessage)
		return
	}
	var filename []string
	if filename = strings.Split(header.Filename, "."); len(filename) <= 1 {
		responseMessage.Message = "the input file is not valid , please try picture types only"
		responseMessage.Success = false
		context.JSON(http.StatusBadRequest, responseMessage)
		return

	}
	extension := filename[1]
	// check if the input is the correct format
	isExtensionRight := helper.CheckExstension(extension)
	if !isExtensionRight {
		responseMessage.Message = "please use images only"
		responseMessage.Success = false
		context.JSON(http.StatusBadRequest, responseMessage)
		return
	}
	// create a custom file name using username and exstension from the input image
	// let's use png extension for filtering mechanism , so there won't be a multiple file for a single username
	// it will overwrite every time a user update profile
	fileName := session.UserName + "." + "png"
	// create the image file in the static file folder
	profileUri := "../../pkg/Entity/Static/profile/" + fileName
	profileNamePath, err := os.Create(profileUri)
	if err != nil {
		responseMessage.Message = "fialed to save profile, "
		responseMessage.Success = false
		context.JSON(http.StatusInternalServerError, responseMessage)
		return
	}

	defer profileNamePath.Close()
	_, err = io.Copy(profileNamePath, file)
	if err != nil {
		return
	}

	// get user info and update the profileurl to the new one
	user, _ := authHandler.authService.FindUserByUserName(session.UserName)
	user.ProfileUrl = profileUri

	// chech is the profile is updated
	isProfileUrlUpdated := authHandler.authService.UpdateUserInfo(user)
	if !isProfileUrlUpdated {
		responseMessage.Message = "internal error , can't update profile please try again"
		responseMessage.Success = false
		context.JSON(http.StatusOK, responseMessage)
		return
	}
	// the profile is uploaded succesfuly
	responseMessage.Message = "profile uploaded succesfuly"
	responseMessage.Success = true
	context.JSON(http.StatusOK, responseMessage)
}

func (authHandler *AuthHandler) DownloadProfile(context *gin.Context) {
	responseMessage := authmodel.ResponseMessage{}
	// Session is used to get user's username
	// username is used for tracking user's profile
	session, _ := authHandler.CookieHandler.ValidateCookie(context)

	// lets get the profile from user's profile_url
	user, _ := authHandler.authService.FindUserByUserName(session.UserName)
	profilePath := user.ProfileUrl

	profile, err := os.Open(profilePath)
	// let's check if somting went wrong when reading profile pic
	if err != nil {
		responseMessage.Message = "you don't have profile picture"
		responseMessage.Success = false
		context.JSON(http.StatusInternalServerError, responseMessage)
		return
	}
	defer profile.Close()
	fmt.Println(profile)
	// context.JSON(http.StatusOK, gin.H{"message": "succes"})
	context.FileAttachment(profilePath, session.UserName+".png")

}

func (authHandler *AuthHandler) DeleteProfilePic(context *gin.Context) {
	responseMessage := authmodel.ResponseMessage{}
	session, _ := authHandler.CookieHandler.ValidateCookie(context)
	User, _ := authHandler.authService.FindUserByUserName(session.UserName)

	err := os.Remove(User.ProfileUrl)
	if err != nil {
		responseMessage.Message = "can't delete profile pic , try again"
		responseMessage.Success = false
		context.JSON(http.StatusInternalServerError, responseMessage)
		return
	}
	User.ProfileUrl = ""
	isUserInfoUpdated := authHandler.authService.UpdateUserInfo(User)
	if !isUserInfoUpdated {
		responseMessage.Message = "can't update profile picture , please try again"
		responseMessage.Success = false
		context.JSON(http.StatusInternalServerError, responseMessage)
		return
	}
	responseMessage.Success = true
	responseMessage.Message = "profile deleted succesfuly"
	context.JSON(http.StatusOK, responseMessage)

}

func (authHandler *AuthHandler) ForgotPasswordHandler(context *gin.Context) {
	responseMessage := authmodel.ResponseMessage{}
	userEmail := &struct {
		UserEmail string `json:"email" binding:"required,email"`
	}{}
	if err := context.ShouldBindJSON(&userEmail); err != nil {
		responseMessage.Message = "please use the correct email address"
		responseMessage.Success = false
		context.JSON(http.StatusBadRequest, responseMessage)
		return

	}
	// if the validation is ok then check the email
	user, isEmailExist := authHandler.authService.FindUserByEmail(userEmail.UserEmail)
	if !isEmailExist {
		responseMessage.Message = "the email address is not found, check if it is the the right email"
		responseMessage.Success = false
		context.JSON(http.StatusNotFound, responseMessage)
		return
	}
	fmt.Println(user)

}
