package authmodel

import "time"

type AuthModel struct {
	UserId     string    `json:"user_id,omitempty" gorm:"primaryKey ; autoIncrement:true"`
	UserName   string    `json:"user_name,omitempty" gorm:"primaryKey"`
	Email      string    `json:"email,omitempty" gorm:"primaryKey"`
	Password   string    `json:"password,omitempty"`
	CreatedOn  time.Time `json:"created_on,omitempty"`
	ProfileUrl string    `json:"profile_url,omitempty"`
	Admin      bool      `json:"admin,omitempty"`
}

type DBResponse struct {
	UserId   string `json:"user_id,omitempty" gorm:"primaryKey ; autoIncrement:true"`
	UserName string `json:"user_name,omitempty" gorm:"primaryKey"`
	Email    string `json:"email,omitempty" gorm:"primaryKey"`
	// Password   string    `json:"password,omitempty"`
	CreatedOn  time.Time `json:"created_on,omitempty"`
	ProfileUrl string    `json:"profile_url,omitempty"`
	Admin      bool      `json:"admin,omitempty"`
}
