package userservice

import (
	"backEndGo/models"
	"backEndGo/repository"
)

type ShowUserService interface {
	ServiceGetUserByUid(Uid int) (*models.Customer, error)
}
type UserData struct {
}

// ServiceGetUserByUid implements ShowUserService
func (UserData) ServiceGetUserByUid(Uid int) (*models.Customer, error) {
	repo := repository.NewCustomerRepository()
	course, err := repo.GetCustomerByID(Uid)
	if err != nil {
		panic(err)
	}
	return course, nil
}

func NewUserDataService() ShowUserService {
	return UserData{}
}
