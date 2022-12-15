package apihandler

import (
	"fmt"
	"net/http"
	"time"

	sessionjwt "github.com/DavG20/Tikus_Event_Api/Internal/pkg/Session_JWT"
	authmodel "github.com/DavG20/Tikus_Event_Api/Internal/pkg/auth/Auth_Model"
	eventservice "github.com/DavG20/Tikus_Event_Api/Internal/pkg/event/Event_Service"
	registrationmodel "github.com/DavG20/Tikus_Event_Api/Internal/pkg/registration/Registration_Model"
	registrationservice "github.com/DavG20/Tikus_Event_Api/Internal/pkg/registration/Registration_Service"
	helper "github.com/DavG20/Tikus_Event_Api/pkg/Utils/Helper"
	"github.com/gin-gonic/gin"
)

type RegiHandler struct {
	cookieHandler sessionjwt.CookieHandler
	regiService   registrationservice.RegiService
	eventService  eventservice.EventService
}

func NewRegiHandler(regService registrationservice.RegiService, eventService eventservice.EventService) RegiHandler {
	return RegiHandler{
		regiService:  regService,
		eventService: eventService,
	}
}

func (regiHandler *RegiHandler) CreateRegiHandler(context *gin.Context) {
	responseMessage := authmodel.ResponseMessage{}
	regiInput := registrationmodel.RegModeleUserInput{}
	regiModel := registrationmodel.RegModel{}
	if err := context.ShouldBindJSON(&regiInput); err != nil {
		responseMessage.Message = "invalid input"
		responseMessage.Success = false
		context.JSON(http.StatusBadRequest, responseMessage)
		return
	}
	isEventIdValid := helper.CheckEventId(regiInput.EventId)
	if !isEventIdValid {
		responseMessage.Message = "invalid event id"
		responseMessage.Success = false
		context.JSON(http.StatusBadRequest, responseMessage)
		return
	}
	fmt.Println(regiInput.EventId)
	event, isEventExist := regiHandler.eventService.FindEventByEventId(regiInput.EventId)
	if !isEventExist {
		responseMessage.Message = "event id doesn't exist"
		responseMessage.Success = false
		context.JSON(http.StatusNotFound, responseMessage)
		return
	}
	// check if some tries to register his/her own event
	session, _ := regiHandler.cookieHandler.ValidateCookie(context)
	if event.UserName == session.UserName {
		responseMessage.Message = "you can't register your own event"
		responseMessage.Success = false
		context.JSON(http.StatusBadRequest, responseMessage)
		return
	}
	if event.ReservedSeats < regiInput.ReservedSeats {
		responseMessage.Message = "can't register this event unavailable seat"
		responseMessage.Success = false
		context.JSON(http.StatusInsufficientStorage, responseMessage)
		return
	}
	regiModel.EventId = regiInput.EventId
	regiModel.ReservedSeats = regiInput.ReservedSeats
	regiModel.RegisteredOn = time.Now().Format("2006-01-02T15:04:05.000Z")
	regiModel.UserName = session.UserName
	regiRespnseDb, isSaved := regiHandler.regiService.CreateRegi(regiModel)
	if !isSaved {
		responseMessage.Message = "failed to register"
		responseMessage.Success = false
		context.JSON(http.StatusInternalServerError, responseMessage)
		return
	}
	context.JSON(http.StatusOK, regiRespnseDb)

}
