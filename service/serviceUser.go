package service

import (
	"backEndGo/models"
	"backEndGo/repository"
	//"time"
)

type UserDataService interface {
	ServiceLogin(Email string, Password string) (*models.Coach, *models.Customer, error)
	ServiceLoginFB(FackbookID string) (*models.Coach, *models.Customer, error)
	ServiceRegisterCus(cus *models.Customer) (int64, error)
	ServiceRegisterCoach(coach *models.Coach) (int64, error)
	ServiceUpdatePasswordCus(uid int, cus *models.Customer) (int64, error)
	ServiceUpdatePasswordCoach(cid int, coach *models.Coach) (int64, error)
}
type UserData struct {
}

// ServiceUpdatePasswordCoach implements UserDataService.
func (UserData) ServiceUpdatePasswordCoach(cid int, coach *models.Coach) (int64, error) {
	repo := repository.NewUserRepository()
	result, err := repo.UpdatePasswordCoach(cid, coach.Password)
	if err != nil {
		panic(err)
	}
	return result, nil
}

// ServiceUpdatePasswordCus implements UserDataService.
func (UserData) ServiceUpdatePasswordCus(uid int, cus *models.Customer) (int64, error) {
	repo := repository.NewUserRepository()
	result, err := repo.UpdatePasswordCus(uid, cus.Password)
	if err != nil {
		panic(err)
	}
	return result, nil
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
func (UserData) ServiceRegisterCoach(coach *models.Coach) (int64, error) {
	repoCoach := repository.NewCoachRepository()
	repoRegister := repository.NewUserRepository()
	getAllCoach, err := repoCoach.GetCoachAll()

	if err != nil {
		panic(err)
	}
	for _, c := range *getAllCoach {
		if c.Email == coach.Email {
			if c.Phone == coach.Phone {
				return 0, nil
			}
		}
	}

	RowsAffected, err := repoRegister.RegisterCoach(coach)
	if err != nil {
		return 0, err
	}
	if RowsAffected > 0 {
		return 1, nil
	} else if RowsAffected == 0 {
		return 0, nil
	} else {
		return -1, nil
	}
}

// ServiceRegisterCus implements UserDataService
func (UserData) ServiceRegisterCus(cus *models.Customer) (int64, error) {
	repoRegister := repository.NewUserRepository()
	getAllCus, err := repoCus.GetCustomerAll()

	if err != nil {
		return 0, err
	}
	for _, c := range *getAllCus {
		if c.Email == cus.Email {
			if c.Phone == cus.Phone {
				return 0, nil
			}
		}
	}
	RowsAffected, err := repoRegister.RegisterCus(cus)
	if err != nil {
		return 0, err
	}
	if RowsAffected > 0 {
		return 1, nil
	} else if RowsAffected == 0 {
		return 0, nil
	} else {
		return -1, nil
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
