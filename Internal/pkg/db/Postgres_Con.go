package db

import (
	"fmt"

	authmodel "github.com/DavG20/Tikus_Event_Api/Internal/pkg/auth/Auth_Model"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "dav627438"
	dbname   = "tikus_event"
)

var DB *gorm.DB
var err error

func CreatePostgresConnection() (*gorm.DB, error) {
	PsqlInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable timezone=Asia/Shanghai ", host, user, password, dbname, port)
	DB, err = gorm.Open(postgres.Open(PsqlInfo), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	DB.AutoMigrate(&authmodel.AuthModel{})
	return DB, nil

}
