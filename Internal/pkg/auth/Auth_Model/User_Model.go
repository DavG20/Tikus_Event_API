package authmodel

import "time"

type AuthModel struct {
	UserId     string    `json:"user_id,omitempty" gorm:"primaryKey ; autoIncrement:true"`
	UserName   string    `json:"user_name,omitempty" gorm:"unique" binding:"required"`
	Email      string    `json:"email,omitempty" gorm:"unique" binding:"required"`
	Password   string    `json:"password,omitempty"`
	CreatedOn  time.Time `json:"created_on,omitempty" gorm:"type:time"`
	ProfileUrl string    `json:"profile_url,omitempty"`
	Admin      bool      `json:"admin,omitempty"`
}

type UserInput struct {
	UserName string `json:"user_name,omitempty" gorm:"unique" binding:"required"`
	Email    string `json:"email,omitempty" gorm:"unique" binding:"required"`
	Password string `json:"password,omitempty"`
	// CreatedOn  time.Time `json:"created_on,omitempty" gorm:"type:time"`
	ProfileUrl string `json:"profile_url,omitempty"`
}

// it helps to return the created on value with out changed it to time
type DBResponse struct {
	UserId   string `json:"user_id,omitempty" gorm:"primaryKey ; autoIncrement:true"`
	UserName string `json:"user_name,omitempty" gorm:"unique" binding:"required"`
	Email    string `json:"email,omitempty" gorm:"unique" binding:"required"`
	// Password   string    `json:"password,omitempty"`
	CreatedOn  time.Time `json:"created_on,omitempty" gorm:"type:time"`
	ProfileUrl string    `json:"profile_url,omitempty"`
	Admin      bool      `json:"admin,omitempty"`
}
