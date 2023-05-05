package service

import (
	"backEndGo/models"
	"backEndGo/repository"
	//"time"
)

type UserDataService interface {
	ServiceLogin(Email string, Password string) (*models.Coach, *models.Customer, error)
	ServiceLoginFB(FackbookID string) (*models.Coach, *models.Customer, error)
	ServiceRegisterCus(cus *models.Customer) int64
	ServiceRegisterCoach(coach *models.Coach) int64
}
type UserData struct {
}

// ServiceLoginFB implements UserDataService
func (UserData) ServiceLoginFB(FackbookID string) (*models.Coach, *models.Customer, error) {
	repo := repository.NewUserRepository()
	coach, cus, err := repo.LoginFB(FackbookID)
	if err != nil {
		panic(err)
	}
	return coach, cus, nil
}

var repoCus = repository.NewCustomerRepository()

// ServiceRegisterCoach implements UserDataService
func (UserData) ServiceRegisterCoach(coach *models.Coach) int64 {
	repoCoach := repository.NewCoachRepository()
	repoRegister := repository.NewUserRepository()
	getAllCoach, err := repoCoach.GetCoachAll()

	if err != nil {
		panic(err)
	}
	for _, c := range *getAllCoach {
		if c.Email == coach.Email {
			return 0
		} else if c.Phone == coach.Phone {
			return 0
		}
	}

	RowsAffected := repoRegister.RegisterCoach(coach)
	if RowsAffected > 0 {
		return 1
	} else if RowsAffected == 0 {
		return 0
	} else {
		return -1
	}
}

// ServiceRegisterCus implements UserDataService
func (UserData) ServiceRegisterCus(cus *models.Customer) int64 {
	repoRegister := repository.NewUserRepository()
	getAllCus, err := repoCus.GetCustomerAll()

	if err != nil {
		panic(err)
	}
	for _, c := range *getAllCus {
		if c.Email == cus.Email {
			return 0
		} else if c.Phone == cus.Phone {
			return 0
		}
	}
	RowsAffected := repoRegister.RegisterCus(cus)
	if RowsAffected > 0 {
		return 1
	} else if RowsAffected == 0 {
		return 0
	} else {
		return -1
	}
}

// LoginTwo implements ShowDataService
func (s UserData) ServiceLogin(Email string, Password string) (*models.Coach, *models.Customer, error) {
	repo := repository.NewUserRepository()
	coach, cus, err := repo.Login(Email, Password)
	if err != nil {
		panic(err)
	}
	return coach, cus, nil
}

func NewUserDataService() UserDataService {
	return UserData{}
}
