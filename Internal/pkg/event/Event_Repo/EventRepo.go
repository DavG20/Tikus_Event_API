package eventrepo

import (
	eventmodel "github.com/DavG20/Tikus_Event_Api/Internal/pkg/event/EventModel"
	constants "github.com/DavG20/Tikus_Event_Api/pkg/Utils/Constants"
	"gorm.io/gorm"
)

type IEventRepo interface {
	CreateEvent(*eventmodel.EventModel) (*eventmodel.EventModel, error)
}

type EventRepo struct {
	DB *gorm.DB
}

func NewEventRepo(db *gorm.DB) *EventRepo {
	return &EventRepo{
		DB: db,
	}
}

func (eventRepo *EventRepo) CreateEvent(userInput *eventmodel.EventModel) (*eventmodel.EventModel, error) {

	err := eventRepo.DB.Table(constants.EventTableName).Create(userInput).Error
	if err != nil {
		return nil, err
	}

	return nil, nil
}
