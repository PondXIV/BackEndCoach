package userservice

import (
	"backEndGo/models"
	"backEndGo/repository"
)

type ShowUserService interface {
	ServiceGetUserByUid(Uid int, Email string) (*[]models.Customer, error)
	ServiceUpdateCustomer(Uid int, customer *models.Customer) (int64, error)
}
type UserData struct {
}

// ServiceUpdateCustomer implements ShowUserService
func (UserData) ServiceUpdateCustomer(Uid int, customer *models.Customer) (int64, error) {
	repo := repository.NewCustomerRepository()
	RowsAffected, err := repo.UpdateUser(Uid, customer)
	if err != nil {
		return -1, err
	}
	if RowsAffected > 0 {
		return 1, nil
	} else if RowsAffected == 0 {
		return 0, nil
	} else {
		return -1, nil
	}
}

// ServiceGetUserByUid implements ShowUserService
func (UserData) ServiceGetUserByUid(Uid int, Email string) (*[]models.Customer, error) {
	repo := repository.NewUserRepository()
	course := repo.GetUserID(Uid, Email)

	return course, nil
}

func NewUserDataService() ShowUserService {
	return UserData{}
}
