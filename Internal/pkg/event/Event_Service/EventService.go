package eventservice

import (
	"fmt"

	eventmodel "github.com/DavG20/Tikus_Event_Api/Internal/pkg/event/EventModel"
	eventrepo "github.com/DavG20/Tikus_Event_Api/Internal/pkg/event/Event_Repo"
)

type IEventService interface {
	CreateEvent(eventmodel.EventUserInput) (eventmodel.EventModel, bool)
}

type EventService struct {
	EventRepo eventrepo.IEventRepo
}

func NewEventService(evnetRepo eventrepo.IEventRepo) EventService {
	return EventService{
		EventRepo: evnetRepo,
	}
}

func (eventService *EventService) CreateEvent(userInput *eventmodel.EventModel) (*eventmodel.EventModel, bool) {
	user, err := eventService.EventRepo.CreateEvent(userInput)
	if err != nil {
		fmt.Println("error in event service")
		return nil, false
	}
	return user, true
}
