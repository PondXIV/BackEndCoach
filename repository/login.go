package repository

import (
	"backEndGo/models"

	"gorm.io/gorm"
)

type LoginRepository interface {
	Login(Email string, Password string, Type int) (*[]models.Coach, *[]models.Customer, error)
}
type loginDB struct {
	db *gorm.DB
}

// Login implements LoginRepository
func (l loginDB) Login(Email string, Password string, Type int) (*[]models.Coach, *[]models.Customer, error) {
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

func NewLoginRepository(gormdb *gorm.DB) LoginRepository {
	return loginDB{db: gormdb}
}
