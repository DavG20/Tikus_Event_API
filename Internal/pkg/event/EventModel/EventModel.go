package eventmodel

// user model
type EventModel struct {
	EventID        string `json:"event_id,omitempty"  gorm:"primaryKey; autoIncrement:true"`
	EventTitle     string `json:"event_title"`
	UserName       string `json:"user_name,omitempty"`
	Description    string `json:"description"`
	EventCreatedOn string `json:"eventcreated_on,omitempty" gorm:"time"`
	EventBeginsOn  string `json:"eventbegins_on"`
	EventEndsOn    string `json:"eventends_on"`
	EventDeadline  string `json:"event_deadline"`
	EventPicture   string `json:"event_picture"`
	AllSeats       int    `json:"allseats"`
	ReservedSeats  int    `json:"reserveredseats"`
}

// user input while creating event
type EventUserInput struct {
	EventTitle     string `json:"event_title"`
	UserName       string `json:"user_name"`
	Description    string `json:"description"`
	EventCreatedOn string `json:"eventcreated_on,omitempty"`
	EventBeginsOn  string `json:"eventbegins_on"`
	EventEndsOn    string `json:"eventends_on"`
	EventDeadline  string `json:"event_deadline"`
	EventPicture   string `json:"event_picture"`
	AllSeats       int    `json:"allseats"`
	ReservedSeats  int    `json:"reserveredseats"`
}

// filterd Response for user

type DBResponse struct {
	EventTitle    string `json:"event_title"`
	UserName      string `json:"user_name"`
	Description   string `json:"description"`
	EventBeginsOn string `json:"eventbegins_on"`
	EventEndsOn   string `json:"eventends_on"`
	EventDeadline string `json:"event_deadline"`
	EventPicture  string `json:"event_picture"`
	AllSeats      int    `json:"allseats"`
	ReservedSeats int    `json:"reserveredseats"`
}

type UpdateEventInput struct {
	EventID       string `json:"event_id" binding:"required"`
	EventTitle    string `json:"event_title,omitempty"`
	Description   string `json:"description,omitempty"`
	EventBeginsOn string `json:"eventbegins_on,omitempty"`
	EventEndsOn   string `json:"eventends_on,omitempty"`
	EventDeadline string `json:"event_deadline,omitempty"`
}
