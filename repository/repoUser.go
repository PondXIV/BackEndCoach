package repository

import (
	"backEndGo/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Login(Email string, Password string, Type int) (*models.Coach, *models.Customer, error)
	LoginNotType(Email string, Password string) (*models.Coach, *models.Customer, error)
	LoginFB(fackbookID string) (*models.Coach, *models.Customer, error)
	RegisterCus(cus *models.Customer) int64
	RegisterCoach(coach *models.Coach) int64
}
type userDB struct {
	db *gorm.DB
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

// LoginTwo implements UserRepository
func (u userDB) LoginNotType(Email string, Password string) (*models.Coach, *models.Customer, error) {
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

// registerCoach implements UserRepository
func (u userDB) RegisterCoach(coach *models.Coach) int64 {
	coach.Cid = 0
	result := u.db.Create(&coach)
	if result.Error != nil {
		panic(result.Error)
	}
	return result.RowsAffected
}

// register implements UserRepository
func (u userDB) RegisterCus(cus *models.Customer) int64 {

	cus.Uid = 0
	result := u.db.Create(&cus)
	if result.Error != nil {
		panic(result.Error)
	}

	return result.RowsAffected
}

// Login implements LoginRepository
func (l userDB) Login(Email string, Password string, Type int) (*models.Coach, *models.Customer, error) {
	coachs := models.Coach{}
	customers := models.Customer{}
	if Type == 0 {
		resultCoa := l.db.Where("email = ?", Email).Where("password = ?", Password).Find(&coachs)
		if resultCoa.Error != nil {
			return nil, nil, resultCoa.Error
		}
	} else if Type == 1 {
		resultCus := l.db.Where("email = ?", Email).Where("password = ?", Password).Find(&customers)
		if resultCus.Error != nil {
			return nil, nil, resultCus.Error
		}
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
