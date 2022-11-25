package db

import (
	"fmt"

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
	PsqlInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", host, user, password, dbname, port)
	DB, err = gorm.Open(postgres.Open(PsqlInfo), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return DB, nil

}
