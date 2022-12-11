package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	api_handler "github.com/DavG20/Tikus_Event_Api/API/API_Handler"
	helper "github.com/DavG20/Tikus_Event_Api/pkg/Utils/Helper"

	authmodel "github.com/DavG20/Tikus_Event_Api/Internal/pkg/auth/Auth_Model"
	authrepo "github.com/DavG20/Tikus_Event_Api/Internal/pkg/auth/Auth_Repo"
	authservice "github.com/DavG20/Tikus_Event_Api/Internal/pkg/auth/Auth_Service"
	DB "github.com/DavG20/Tikus_Event_Api/Internal/pkg/db"
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

var Db *gorm.DB
var err error
var once sync.Once
var router *gin.Engine

var users []authmodel.AuthModel

func StartUp() {
	once.Do(
		func() {
			Db, err = DB.CreatePostgresConnection()
			if Db == nil || err != nil {
				log.Fatal("exiting...")
				os.Exit(1)
			}

			return
		},
	)
}

func init() {
	StartUp()
}

func main() {
	authRepo := authrepo.NewAuth(Db)
	authService := authservice.NewAuthService(authRepo)
	authHandler := api_handler.NewauthHandler(authService)

	router = gin.Default()

	public := router.Group("/user")

	rsetCode := helper.EncodeRestPassword("hey")
	correct, err := helper.DecodeRestPassword(rsetCode)
	fmt.Println(correct, err)

	// public routers which doesn't need authorization
	public.POST("/createuser", authHandler.RegisterHandler())
	public.POST("/login", authHandler.LoginHandler)

	// private router, which needs authorization
	private := router.Group("/user")
	private.Use(authHandler.AuthRequired())
	private.GET("/searchuser", authHandler.SearchUser)
	private.POST("/logout", authHandler.LogoutHandler)
	private.POST("/deleteaccount", authHandler.DeleteAccount)
	private.GET("/getuserinfo", authHandler.GetUserInfo)
	private.POST("/changepassword", authHandler.ChangePasswordHandler)
	private.POST("uploadprofile", authHandler.UploadProfileHandler)
	private.GET("downloadprofile", authHandler.DownloadProfile)
	private.POST("/deleteprofile", authHandler.DeleteProfilePic)
	private.GET("forgotpassword", authHandler.ForgotPasswordHandler)

	router.Run(":8080")
}
