package registrationrepo

import (
	registrationmodel "github.com/DavG20/Tikus_Event_Api/Internal/pkg/registration/Registration_Model"
	"gorm.io/gorm"
)

type IRegRepo interface {
	CreateRegi(registrationmodel.RegModel) (*registrationmodel.RegModel, bool)
}

type RegRepo struct {
	Db *gorm.DB
}

func NewRegRepo(db *gorm.DB) *RegRepo {
	return &RegRepo{
		Db: db,
	}
}

func (regRepo *RegRepo) CreateRegi(regInput registrationmodel.RegModel) (*registrationmodel.RegModel, bool) {
	err := regRepo.Db.Create(regInput).Error
	if err != nil {
		return nil, false
	}
	return nil, true

}
