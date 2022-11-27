package authmodel

type AuthModel struct {
	UserId     string `json:"user_id,omitempty" gorm:"primaryKey ; autoIncrement:true"`
	UserName   string `json:"user_name,omitempty" gorm:"unique" binding:"required,min=4"`
	Email      string `json:"email,omitempty" gorm:"unique" binding:"required,email"`
	Password   string `json:"password,omitempty" binding:"required"`
	CreatedOn  string `json:"created_on,omitempty" `
	ProfileUrl string `json:"profile_url,omitempty"`
	Admin      bool   `json:"admin,omitempty"`
}

type UserInput struct {
	UserName string `json:"user_name,omitempty" gorm:"unique" binding:"required,min=4"`
	Email    string `json:"email,omitempty" gorm:"unique" binding:"required,email"`
	Password string `json:"password,omitempty" binding:"required"`
	// CreatedOn  time.Time `json:"created_on,omitempty" gorm:"type:time"`
	ProfileUrl string `json:"profile_url,omitempty"`
}

// it helps to return the created on value with out changed it to time
type DBResponse struct {
	UserId   string `json:"user_id,omitempty" gorm:"primaryKey ; autoIncrement:true"`
	UserName string `json:"user_name,omitempty" gorm:"unique" binding:"required,min=4"`
	Email    string `json:"email,omitempty" gorm:"unique" binding:"required,email"`
	// Password   string    `json:"password,omitempty"`
	CreatedOn  string `json:"created_on,omitempty" gorm:"type:time"`
	ProfileUrl string `json:"profile_url,omitempty"`
	Admin      bool   `json:"admin,omitempty"`
}

type ErrMessage struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ResponseMessage struct {
	Message      string `json:"message"`
	Success      bool   `json:"success"`
	InvalidField string `json:"error"`
}
