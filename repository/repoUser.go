package repository

import (
	"backEndGo/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Login(Email string, Password string, Type int) (*[]models.Coach, *[]models.Customer, error)
	LoginNotType(Email string, Password string) (*[]models.Coach, *[]models.Customer, error)
	// RegisterCus(AliasName string,
	// 	Password string,
	// 	Email string,
	// 	FullName string,
	// 	Birthday time.Time,
	// 	Gender string,
	// 	Phone string,
	// 	Image string,
	// 	Weight int,
	// 	Height int,
	// 	Price int,) int64
	RegisterCus(cus *models.Customer) int64
	RegisterCoach(coach *models.Coach) int64
}
type userDB struct {
	db *gorm.DB
}

// LoginTwo implements UserRepository
func (l userDB) LoginNotType(Email string, Password string) (*[]models.Coach, *[]models.Customer, error) {
	coachs := []models.Coach{}
	customers := []models.Customer{}

	resultCus := l.db.Where("email = ?", Email).Where("password = ?", Password).Find(&coachs)
	if resultCus.Error != nil {
		return nil, nil, resultCus.Error
	}

	resultCoa := l.db.Where("email = ?", Email).Where("password = ?", Password).Find(&customers)
	if resultCoa.Error != nil {
		return nil, nil, resultCoa.Error
	}

	return &coachs, &customers, nil
}

// registerCoach implements UserRepository
func (u userDB) RegisterCoach(coach *models.Coach) int64 {

	result := u.db.Create(&coach)
	if result.Error != nil {
		panic(result.Error)
	}
	return result.RowsAffected
}

// register implements UserRepository
func (u userDB) RegisterCus(cus *models.Customer) int64 {
	//cus := models.Customer{}
	result := u.db.Create(&cus)
	if result.Error != nil {
		panic(result.Error)
	}

	return result.RowsAffected
}

// Login implements LoginRepository
func (l userDB) Login(Email string, Password string, Type int) (*[]models.Coach, *[]models.Customer, error) {
	coachs := []models.Coach{}
	customers := []models.Customer{}
	if Type == 0 {
		result := l.db.Where("email = ?", Email).Where("password = ?", Password).Find(&coachs)
		if result.Error != nil {
			return nil, nil, result.Error
		}
	}
	if Type == 1 {
		result := l.db.Where("email = ?", Email).Where("password = ?", Password).Find(&customers)
		if result.Error != nil {
			return nil, nil, result.Error
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
