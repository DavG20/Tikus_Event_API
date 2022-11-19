package main

import (
	"log"
	"os"
	"sync"

	api_handler "github.com/DavG20/Tikus_Event_Api/API/API_Handler"
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

	router.GET("/", authHandler.CreateUserHandler)
	router.Run(":9090")
}
