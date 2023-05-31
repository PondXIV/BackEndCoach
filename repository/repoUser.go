package repository

import (
	"backEndGo/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Login(Email string, Password string) (*models.Coach, *models.Customer, error)
	LoginFB(fackbookID string) (*models.Coach, *models.Customer, error)
	RegisterCus(cus *models.Customer) (int64, error)
	RegisterCoach(coach *models.Coach) (int64, error)
	GetUserID(Uid int) *models.Customer
}
type userDB struct {
	db *gorm.DB
}

// GetUserID implements UserRepository
func (u userDB) GetUserID(Uid int) *models.Customer {
	customer := models.Customer{}
	result := u.db.Where("uid = ?", Uid).Find(&customer)
	if result.Error != nil {
		panic(result.Error)
	}
	return &customer
}

// UserByUid implements UserRepository

// LoginFB implements UserRepository
func (u userDB) LoginFB(fackbookID string) (*models.Coach, *models.Customer, error) {
	coachs := models.Coach{}
	customers := models.Customer{}

	resultCoa := u.db.Where("facebookId = ?", fackbookID).Find(&coachs)
	if resultCoa.Error != nil {
		return nil, nil, resultCoa.Error
	}

	resultCus := u.db.Where("facebookId = ?", fackbookID).Find(&customers)
	if resultCus.Error != nil {
		return nil, nil, resultCus.Error
	}

	return &coachs, &customers, nil
}

// registerCoach implements UserRepository
func (u userDB) RegisterCoach(coach *models.Coach) (int64, error) {
	coach.Cid = 0
	result := u.db.Create(&coach)
	if result.Error != nil {
		return -1, result.Error
	}
	if result.Error != nil {
		panic(result.Error)
	}

	return result.RowsAffected, nil
}

// register implements UserRepository
func (u userDB) RegisterCus(cus *models.Customer) (int64, error) {

	cus.Uid = 0
	result := u.db.Create(&cus)
	if result.Error != nil {
		return -1, result.Error
	}
	if result.Error != nil {
		panic(result.Error)
	}

	return result.RowsAffected, nil
}

// LoginTwo implements UserRepository
func (u userDB) Login(Email string, Password string) (*models.Coach, *models.Customer, error) {
	coachs := models.Coach{}
	customers := models.Customer{}

	resultCoa := u.db.Where("email = ?", Email).Where("password = ?", Password).Find(&coachs)
	if resultCoa.Error != nil {
		return nil, nil, resultCoa.Error
	}

	resultCus := u.db.Where("email = ?", Email).Where("password = ?", Password).Find(&customers)
	if resultCus.Error != nil {
		return nil, nil, resultCus.Error
	}

	return &coachs, &customers, nil
}
func NewUserRepository() UserRepository {
	db, err := NewDatabaseConnection()
	if err != nil {
		return nil
	}

	return userDB{db}
}
