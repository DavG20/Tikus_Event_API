package apihandler

import (
	"fmt"
	"net/http"

	sessionjwt "github.com/DavG20/Tikus_Event_Api/Internal/pkg/Session_JWT"
	authmodel "github.com/DavG20/Tikus_Event_Api/Internal/pkg/auth/Auth_Model"
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
	responseMessage := authmodel.ResponseMessage{}
	userInput := &eventmodel.EventUserInput{}
	if err := context.ShouldBindJSON(&userInput); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "can't bind json"})
		return
	}

	// lets save the event
	// createEvent method in repo will add the event in to the db and return the newly added event id
	eventId, issaved := eventHandler.eventService.CreateEvent(userInput)
	if !issaved {
		responseMessage.Message = "Internal server error,"
		responseMessage.Success = false
		context.JSON(http.StatusInternalServerError, responseMessage)
		return
	}

	eventFromDB, isEventExist := eventHandler.eventService.FindEventByEventId(eventId)
	if !isEventExist {
		responseMessage.Message = "can't get event , internal server error"
		responseMessage.Success = false
		context.JSON(http.StatusInternalServerError, responseMessage)
		return
	}
	context.JSON(http.StatusOK, eventFromDB)

}

func (eventHandler *EventHandler) UplaodEventProfilePic(context *gin.Context) {
	responseMessage := authmodel.ResponseMessage{}
	//  here the request will automatically checked by authrequired  func if the user is authenticated or not
	// but here we need the event's ID to update event's profile
	eventID := context.Request.FormValue("eventid")
	if eventID == "" {
		responseMessage.Message = "Invalid event id please use the correct event id to update event profile"
		responseMessage.Success = false
		context.JSON(http.StatusBadRequest, responseMessage)
		return

	}
	fmt.Println("fine")

}
