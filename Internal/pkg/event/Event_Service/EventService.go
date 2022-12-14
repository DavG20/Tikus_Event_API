package eventservice

import (
	eventmodel "github.com/DavG20/Tikus_Event_Api/Internal/pkg/event/EventModel"
	eventrepo "github.com/DavG20/Tikus_Event_Api/Internal/pkg/event/Event_Repo"
)

type IEventService interface {
	CreateEvent(eventmodel.EventUserInput) bool
	FindEventByEventId(string) (eventmodel.EventModel, bool)
	SaveEvent(*eventmodel.EventModel) *eventmodel.EventModel
	EventEncoder(*eventmodel.EventModel, eventmodel.UpdateEventInput) (*eventmodel.EventModel, bool)
	DeleteEvent(*eventmodel.EventModel) bool
}

type EventService struct {
	EventRepo eventrepo.IEventRepo
}

func NewEventService(evnetRepo eventrepo.IEventRepo) EventService {
	return EventService{
		EventRepo: evnetRepo,
	}
}

func (eventService *EventService) CreateEvent(userInput *eventmodel.EventUserInput) (string, bool) {
	eventId, err := eventService.EventRepo.CreateEvent(userInput)
	if err != nil {
		return "", false
	}
	return eventId, true
}

func (eventService *EventService) FindEventByEventId(eventID string) (*eventmodel.EventModel, bool) {
	event, err := eventService.EventRepo.FindEventByEventId(eventID)
	if err != nil {
		return nil, false
	}
	return event, true
}

func (eventService *EventService) SaveEvent(eventInput *eventmodel.EventModel) (*eventmodel.EventModel, bool) {
	return eventService.EventRepo.SaveEvent(eventInput)
}

func (eventService *EventService) EventEncoder(event *eventmodel.EventModel, EventUpdateInput eventmodel.UpdateEventInput) (*eventmodel.EventModel, bool) {
	return eventService.EventRepo.EventEncoder(event, EventUpdateInput)
}

func (eventService *EventService) DeleteEvent(event *eventmodel.EventModel) bool {
	return eventService.EventRepo.DeleteEvent(event)
}
