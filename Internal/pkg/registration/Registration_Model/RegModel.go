package registrationmodel

type RegModel struct {
	RegId         string `json:"reg_id,omitempty" gorm:"primaryKey; autoIncrement:true"`
	EventId       string `json:"event_id"`
	UserName      string `json:"user_name"`
	RegisteredOn  string `json:"registered_on"`
	ReservedSeats int    `json:"reserved_seats"`
}

type RegModeleUserInput struct {
	EventId       string `json:"event_id" binding:"required"`
	ReservedSeats int    `json:"reserved_seats" binding:"required"`
}
