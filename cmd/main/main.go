package main

import (
	"log"
	"os"
	"sync"

	api_handler "github.com/DavG20/Tikus_Event_Api/API/API_Handler"
	apihandler "github.com/DavG20/Tikus_Event_Api/API/API_Handler"

	authmodel "github.com/DavG20/Tikus_Event_Api/Internal/pkg/auth/Auth_Model"
	authrepo "github.com/DavG20/Tikus_Event_Api/Internal/pkg/auth/Auth_Repo"
	authservice "github.com/DavG20/Tikus_Event_Api/Internal/pkg/auth/Auth_Service"
	"github.com/DavG20/Tikus_Event_Api/Internal/pkg/db"
	DB "github.com/DavG20/Tikus_Event_Api/Internal/pkg/db"
	registrationmodel "github.com/DavG20/Tikus_Event_Api/Internal/pkg/registration/Registration_Model"
	registrationrepo "github.com/DavG20/Tikus_Event_Api/Internal/pkg/registration/Registration_Repo"
	registrationservice "github.com/DavG20/Tikus_Event_Api/Internal/pkg/registration/Registration_Service"

	eventrepo "github.com/DavG20/Tikus_Event_Api/Internal/pkg/event/Event_Repo"
	eventService "github.com/DavG20/Tikus_Event_Api/Internal/pkg/event/Event_Service"
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

	// auth part
	authRepo := authrepo.NewAuth(Db)
	authService := authservice.NewAuthService(authRepo)
	authHandler := api_handler.NewauthHandler(authService)

	// event
	eventRepo := eventrepo.NewEventRepo(Db)
	eventService := eventService.NewEventService(eventRepo)
	eventHandler := api_handler.NewEventHandler(eventService)

	// regi
	regiRepo := registrationrepo.NewRegRepo(Db)
	regiService := registrationservice.NewRegiService(regiRepo)
	regiHandler := apihandler.NewRegiHandler(regiService, eventService)

	router = gin.Default()

	db.DB.AutoMigrate(registrationmodel.RegModel{})

	public := router.Group("/user")

	// public routers which doesn't need authorization
	public.POST("/createuser", authHandler.RegisterHandler())
	public.POST("/login", authHandler.LoginHandler)

	// private router, which needs authorization
	private := router.Group("/user")
	private.Use(authHandler.AuthRequired())
	private.GET("/searchuser", authHandler.SearchUser)
	private.POST("/logout", authHandler.LogoutHandler)
	private.DELETE("/deleteaccount", authHandler.DeleteAccount)
	private.GET("/getuserinfo", authHandler.GetUserInfo)
	private.PUT("/changepassword", authHandler.ChangePasswordHandler)
	private.PUT("/uploadprofile", authHandler.UploadProfileHandler)
	private.GET("/downloadprofile", authHandler.DownloadProfile)
	private.DELETE("/deleteprofile", authHandler.DeleteProfilePic)
	private.GET("/forgotpassword", authHandler.ForgotPasswordHandler)

	// event part
	private.POST("event/createevent", eventHandler.CreateEventHendler)
	private.PUT("event/uploadeeventpic", eventHandler.UploadEventProfilePic)
	private.PUT("event/update", eventHandler.UpdateEventHandler)
	private.DELETE("event/deleteevent", eventHandler.DeleteEventHandler)
	private.GET("event/geteventinfo", eventHandler.GetEventInfoHandler)

	// registration part
	private.POST("regi/createregi", regiHandler.CreateRegiHandler)

	router.Run(":8080")
}
