package eventmodel

import "time"

// user model
type EventModel struct {
	EventID        string    `json:"event_id,omitempty"  gorm:"primaryKey; autoIncrement:true"`
	UserName       string    `json:"user_name,omitempty"`
	Description    string    `json:"description"`
	EventCreatedOn time.Time `json:"eventcreated_on,omitempty"`
	EventBeginsOn  time.Time `json:"eventbegins_on"`
	EventEndsOn    time.Time `json:"eventends_on"`
	EventDeadline  time.Time `json:"event_deadline"`
	EventPicture   string    `json:"event_picture"`
	AllSeats       int       `json:"allseats"`
	ReservedSeats  int       `json:"reserveredseats"`
}

// user input while creating event
type EventUserInput struct {
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
	UserName      string    `json:"user_name"`
	Description   string    `json:"description"`
	EventBeginsOn time.Time `json:"eventbegins_on"`
	EventEndsOn   time.Time `json:"eventends_on"`
	EventDeadline time.Time `json:"event_deadline"`
	EventPicture  string    `json:"event_picture"`
	AllSeats      int       `json:"allseats"`
	ReservedSeats int       `json:"reserveredseats"`
}
