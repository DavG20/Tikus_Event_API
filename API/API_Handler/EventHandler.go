package apihandler

import (
	"fmt"
	"net/http"
	"time"

	sessionjwt "github.com/DavG20/Tikus_Event_Api/Internal/pkg/Session_JWT"
	eventmodel "github.com/DavG20/Tikus_Event_Api/Internal/pkg/event/EventModel"
	eventservice "github.com/DavG20/Tikus_Event_Api/Internal/pkg/event/Event_Service"
	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	eventService  eventservice.EventService
	cookieHandler sessionjwt.CookieHandler
}

func NewEventHandler(eventService eventservice.EventService) *EventHandler {
	return &EventHandler{
		eventService: eventService,
	}
}

func (eventHandler *EventHandler) CreateEventHendler(context *gin.Context) {
	userInput := eventmodel.EventUserInput{}
	if err := context.ShouldBindJSON(&userInput); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "can't bind json"})
		return
	}
	fmt.Println("fine", userInput)
	begin, err := (time.Parse("2006-01-02T15:04:05.000Z", userInput.EventDeadline))
	fmt.Println(begin, err)
	event := &eventmodel.EventModel{
		UserName:       userInput.UserName,
		Description:    userInput.Description,
		EventCreatedOn: time.Now(),
		EventEndsOn:    begin,
		EventDeadline:  begin,
		EventBeginsOn:  begin,
		AllSeats:       userInput.AllSeats,
		ReservedSeats:  userInput.ReservedSeats,
	}
	savedEvent, issaved := eventHandler.eventService.CreateEvent(event)
	if !issaved {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "error"})
		return
	}
	fmt.Println("fine", savedEvent)

}
