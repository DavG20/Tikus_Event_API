package eventrepo

import (
	"fmt"
	"time"

	eventmodel "github.com/DavG20/Tikus_Event_Api/Internal/pkg/event/EventModel"
	constants "github.com/DavG20/Tikus_Event_Api/pkg/Utils/Constants"
	helper "github.com/DavG20/Tikus_Event_Api/pkg/Utils/Helper"
	"gorm.io/gorm"
)

type IEventRepo interface {
	CreateEvent(*eventmodel.EventUserInput) (string, error)
	FindEventByEventId(string) (*eventmodel.EventModel, error)
	SaveEvent(*eventmodel.EventModel) (*eventmodel.EventModel, bool)
	EventEncoder(*eventmodel.EventModel, eventmodel.UpdateEventInput) (*eventmodel.EventModel, bool)
	DeleteEvent(*eventmodel.EventModel) bool
}

type EventRepo struct {
	DB *gorm.DB
}

func NewEventRepo(db *gorm.DB) *EventRepo {
	return &EventRepo{
		DB: db,
	}
}

func (eventRepo *EventRepo) CreateEvent(userInput *eventmodel.EventUserInput) (string, error) {

	eventbeginOn, eventBeginsOn, eventDeadline, err := helper.ParseDateInput(userInput.EventBeginsOn, userInput.EventEndsOn, userInput.EventDeadline)
	if err != nil {

		return "", err
	}

	event := &eventmodel.EventModel{
		EventTitle:     userInput.EventTitle,
		UserName:       userInput.UserName,
		Description:    userInput.Description,
		EventCreatedOn: time.Now().Format("2006-01-02 15:04:05.12"),
		EventEndsOn:    eventbeginOn.Format("2006-01-02 15:04:05.12"),
		EventDeadline:  eventDeadline.Format("2006-01-02 15:04:05.12"),
		EventBeginsOn:  eventBeginsOn.Format("2006-01-02 15:04:05.12"),
		AllSeats:       userInput.AllSeats,
		ReservedSeats:  userInput.ReservedSeats,
	}

	ressult := eventRepo.DB.Table(constants.EventTableName).Create(event)

	if ressult.Error != nil {
		return "", ressult.Error
	}
	return event.EventID, nil
}

func (eventRepo *EventRepo) FindEventByEventId(eventID string) (event *eventmodel.EventModel, err error) {
	err = eventRepo.DB.Table(constants.EventTableName).Where("event_id=?", eventID).First(&event).Error

	if err != nil {
		return nil, err
	}
	return event, nil

}

func (eventRepo *EventRepo) SaveEvent(eventInput *eventmodel.EventModel) (event *eventmodel.EventModel, isSaved bool) {
	err := eventRepo.DB.Table(constants.EventTableName).Save(eventInput).Error
	if err != nil {
		fmt.Println("error here")
		return nil, false
	}
	event, _ = eventRepo.FindEventByEventId(eventInput.EventID)
	return event, true
}

func (eventRepo *EventRepo) EventEncoder(event *eventmodel.EventModel, EventUpdateInput eventmodel.UpdateEventInput) (*eventmodel.EventModel, bool) {

	if EventUpdateInput.Description != "" {
		event.Description = EventUpdateInput.Description
	}
	if EventUpdateInput.EventTitle != "" {
		event.EventTitle = EventUpdateInput.EventTitle
	}
	if EventUpdateInput.EventEndsOn != "" {
		eventEnds, err := helper.SingleDateHelper(EventUpdateInput.EventEndsOn)
		if err == nil {
			event.EventEndsOn = eventEnds.Format("2006-01-02 15:04:05.12")
		}

	}
	if EventUpdateInput.EventBeginsOn != "" {
		eventBegins, err := helper.SingleDateHelper(EventUpdateInput.EventBeginsOn)
		if err == nil {
			event.EventBeginsOn = eventBegins.Format("2006-01-02 15:04:05.12")
		}

	}
	if EventUpdateInput.EventDeadline != "" {
		eventDeadline, err := helper.SingleDateHelper(EventUpdateInput.EventDeadline)
		if err == nil {
			event.EventBeginsOn = eventDeadline.Format("2006-01-02 15:04:05.12")
		}

	}

	// updatedEvent, isUpdated := eventRepo.SaveEvent(event)
	// if !isUpdated {
	// 	fmt.Println("can't update event")
	// 	return nil, false
	// }
	return event, true

}

func (eventRepo *EventRepo) DeleteEvent(event *eventmodel.EventModel) bool {
	err := eventRepo.DB.Table(constants.EventTableName).Delete(event).Error
	if err != nil {
		fmt.Println("can't delete event in repo")
		return false
	}
	return true
}
