package main

import (
	"log"
	"os"
	"sync"

	api_handler "github.com/DavG20/Tikus_Event_Api/API/API_Handler"

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

	// public routers which doesn't need authorization
	public.POST("/createuser", authHandler.RegisterHandler())

	// private router, which needs authorization
	private := router.Group("/user")
	private.Use(authHandler.AuthRequired())
	private.GET("/searchuser", authHandler.SearchUser)

	router.Run(":8080")
}
