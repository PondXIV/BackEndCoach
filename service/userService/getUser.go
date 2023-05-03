package userservice

import (
	"backEndGo/models"
	"backEndGo/repository"
)

type ShowUserService interface {
	ServiceGetUserByUid(Uid int) (*models.Customer, error)
	ServiceUpdateCustomer(Uid int, customer *models.Customer) int64
}
type UserData struct {
}

// ServiceUpdateCustomer implements ShowUserService
func (UserData) ServiceUpdateCustomer(Uid int, customer *models.Customer) int64 {
	repo := repository.NewCustomerRepository()
	RowsAffected := repo.UpdateUser(Uid, customer)
	if RowsAffected > 0 {
		return 1
	} else if RowsAffected == 0 {
		return 0
	} else {
		return -1
	}
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
