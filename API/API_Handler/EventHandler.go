package apihandler

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	sessionjwt "github.com/DavG20/Tikus_Event_Api/Internal/pkg/Session_JWT"
	authmodel "github.com/DavG20/Tikus_Event_Api/Internal/pkg/auth/Auth_Model"
	eventmodel "github.com/DavG20/Tikus_Event_Api/Internal/pkg/event/EventModel"
	eventservice "github.com/DavG20/Tikus_Event_Api/Internal/pkg/event/Event_Service"
	helper "github.com/DavG20/Tikus_Event_Api/pkg/Utils/Helper"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

	eventId := context.Request.FormValue("eventid")
	if eventId == "" {
		responseMessage.Message = "invalid event id , please provide event id"
		responseMessage.Success = false
		context.JSON(http.StatusBadRequest, responseMessage)
		return
	}

	// lets get event profile
	file, header, err := context.Request.FormFile("eventpic")
	if err != nil {
		responseMessage.Message = "can't get event profile , please try again"
		responseMessage.Success = false
		context.JSON(http.StatusBadRequest, responseMessage)
		return
	}
	// lets check if requestfile is file
	var fileName []string
	if fileName = strings.Split(header.Filename, "."); len(fileName) <= 1 {
		responseMessage.Message = "invalid input , please upload image formats only"
		responseMessage.Success = false
		context.JSON(http.StatusBadRequest, responseMessage)
		return
	}
	// lets check if the input is image
	extension := helper.CheckExstension(fileName[1])
	if !extension {
		responseMessage.Message = "invalid input , please upload image formats only"
		responseMessage.Success = false
		context.JSON(http.StatusBadRequest, responseMessage)
		return
	}
	// Username used for profile name so lets get username from session
	session, _ := eventHandler.cookieHandler.ValidateCookie(context)

	event, isEventExist := eventHandler.eventService.FindEventByEventId(eventId)
	if !isEventExist {
		responseMessage.Message = "event doesn't exist by this id"
		responseMessage.Success = false
		context.JSON(http.StatusOK, responseMessage)
		return
	}
	// lets check if the event is created by this user
	if session.UserName != event.UserName {
		responseMessage.Message = "You don't have event by this EventID "
		responseMessage.Success = false
		context.JSON(http.StatusUnauthorized, responseMessage)
		return
	}

	// then lets upload event profile
	// the profile name looks like username_eventid.png
	// the name format helps us to find events profile easly
	profilePath := helper.SaveProfileInFileSystem(file, session.UserName, eventId)
	fmt.Println("finennnnnn", profilePath)

	// Update event table
	event.EventPicture = profilePath
	fmt.Println(event, "event")
	updatedEvent, isEventSaved := eventHandler.eventService.SaveEvent(event)
	if !isEventSaved {
		responseMessage.Message = "Failed to update event, please try again"
		responseMessage.Success = false
		context.JSON(http.StatusInternalServerError, responseMessage)
		return
	}
	context.JSON(http.StatusOK, updatedEvent)
}

func (eventHandeler *EventHandler) UpdateEventHandler(context *gin.Context) {
	EventUpdateInput := eventmodel.UpdateEventInput{}
	responseMessage := authmodel.ResponseMessage{}

	if err := context.ShouldBindJSON(&EventUpdateInput); err != nil {
		var valErr validator.ValidationErrors
		if errors.As(err, &valErr) {
			errMessage := make([]authmodel.ResponseMessage, len(valErr))
			for i, fieldErr := range valErr {
				errMessage[i] = authmodel.ResponseMessage{Message: helper.GetErrorMessage(fieldErr), Success: false, InvalidField: fieldErr.Field()}
			}
			context.JSON(http.StatusBadRequest, errMessage)
			return
		}
		responseMessage.Message = "unknown error please check ur input"
		responseMessage.Success = false
		context.JSON(http.StatusBadRequest, responseMessage)
		return
	}
	fmt.Println("fine here")
}
