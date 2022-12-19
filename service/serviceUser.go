package service

import (
	"backEndGo/models"
	"backEndGo/repository"
	"time"

	"gorm.io/gorm"
)

type UserDataService interface {
	ServiceLogin(Email string, Password string, Type int) (*[]models.Coach, *[]models.Customer, error)
	ServiceLoginNotType(Email string, Password string) (*[]models.Coach, *[]models.Customer, error)
	ServiceRegisterCus(
		AliasName string,
		Password string,
		Email string,
		FullName string,
		Birthday time.Time,
		Gender string,
		Phone string,
		Image string,
		Weight int,
		Height int,
		Price int,
	) (*[]models.Customer, error)
}
type UserData struct {
	db *gorm.DB
}

// ServiceRegisterCus implements UserDataService
func (UserData) ServiceRegisterCus(AliasName string, Password string, Email string, FullName string, Birthday time.Time, Gender string, Phone string, Image string, Weight int, Height int, Price int) (*[]models.Customer, error) {
	repo := repository.NewUserRepository()
	cus, err := repo.RegisterCus(AliasName, Password,
		Email, FullName, Birthday,
		Gender, Phone, Image, Weight, Height, Price)
	if err != nil {
		panic(err)
	}
	return cus, nil
}

// LoginTwo implements ShowDataService
func (s UserData) ServiceLoginNotType(Email string, Password string) (*[]models.Coach, *[]models.Customer, error) {
	repo := repository.NewUserRepository()
	coach, cus, err := repo.LoginNotType(Email, Password)
	if err != nil {
		panic(err)
	}
	return coach, cus, nil
}

// Login implements ShowDataService
func (s UserData) ServiceLogin(Email string, Password string, Type int) (*[]models.Coach, *[]models.Customer, error) {
	repo := repository.NewUserRepository()
	coach, cus, err := repo.Login(Email, Password, Type)
	if err != nil {
		panic(err)
	}
	return coach, cus, nil
}

func NewUserDataService() UserDataService {
	return UserData{}
}
