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

func (UserData) ServiceGetUserByUid(Uid int) (*models.Customer, error) {
	repo := repository.NewCustomerRepository()
	user, err := repo.GetCustomerByID(Uid)
	if err != nil {
		panic(err)
	}
	return user, nil
}
func NewUserDataService() ShowUserService {
	return UserData{}
}
