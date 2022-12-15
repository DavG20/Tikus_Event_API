package registrationservice

import (
	registrationmodel "github.com/DavG20/Tikus_Event_Api/Internal/pkg/registration/Registration_Model"
	registrationrepo "github.com/DavG20/Tikus_Event_Api/Internal/pkg/registration/Registration_Repo"
)

type iRegService interface {
	CreateRegi(registrationmodel.RegModel) (*registrationmodel.RegModel, bool)
}

type RegiService struct {
	regiRepo registrationrepo.IRegRepo
}

func NewRegiService(regiRepo registrationrepo.IRegRepo) RegiService {
	return RegiService{
		regiRepo: regiRepo,
	}
}

func (regiService *RegiService) CreateRegi(regiInput registrationmodel.RegModel) (*registrationmodel.RegModel, bool) {
	return regiService.regiRepo.CreateRegi(regiInput)

}
