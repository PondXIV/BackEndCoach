package repository

import (
	"backEndGo/models"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	GetCustomerAll() (*[]models.Customer, error)
	LoginCustomer(Email string, Password string) (*[]models.Customer, error)
}
type custimerDB struct {
	db *gorm.DB
}

// LoginCustomer implements CustomerRepository
func (c custimerDB) LoginCustomer(Email string, Password string) (*[]models.Customer, error) {
	customers := []models.Customer{}
	//AliasName = "%" + name + "%"
	result := c.db.Where("email = ?", Email).Where("password = ?", Password).Find(&customers)
	if result.Error != nil {
		return nil, result.Error
	}
	return &customers, nil
}

// GetAll implements CustomerRepository
func (c custimerDB) GetCustomerAll() (*[]models.Customer, error) {
	customers := []models.Customer{}
	result := c.db.Find(&customers)
	if result.Error != nil {
		return nil, result.Error
	}
	return &customers, nil
}

func NewCustomerRepository(gormdb *gorm.DB) CustomerRepository {
	return custimerDB{db: gormdb}
}
