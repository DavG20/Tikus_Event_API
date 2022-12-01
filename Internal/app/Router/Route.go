package router

import (
	"log"
	"os"
	"sync"

	apihandler "github.com/DavG20/Tikus_Event_Api/API/API_Handler"
	authrepo "github.com/DavG20/Tikus_Event_Api/Internal/pkg/auth/Auth_Repo"
	authservice "github.com/DavG20/Tikus_Event_Api/Internal/pkg/auth/Auth_Service"
	DB "github.com/DavG20/Tikus_Event_Api/Internal/pkg/db"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

var Db *gorm.DB
var err error
var once sync.Once
var router *gin.Engine

var authHandler apihandler.AuthHandler
var auth_repo authrepo.AuthRepo

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
	auth_repo = authrepo.NewAuth(Db)
	authService := authservice.NewAuthService(auth_repo)
	authHandler = apihandler.NewauthHandler(authService)
	StartUp()
}

func PrivateRoute(g *gin.RouterGroup) {
	// g.GET("/userinfo")

}
func PublicRoute(g *gin.RouterGroup) {
	g.POST("/register", authHandler.RegisterHandler())
	// g.POST("/login")
}
