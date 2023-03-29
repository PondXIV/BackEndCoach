package repository

import (
	"backEndGo/models"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	GetCustomerAll() (*[]models.Customer, error)
	GetCustomerByID(Id int) (*models.Customer, error)
	UserByUid(Uid int) (*models.Customer, error)
}
type custimerDB struct {
	db *gorm.DB
}

// UserByUid implements CustomerRepository
func (c custimerDB) UserByUid(Uid int) (*models.Customer, error) {
	users := models.Customer{}
	result := c.db.Where("uid = ?", Uid).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return &users, nil
}

// GetCustomerByID implements CustomerRepository
func (c custimerDB) GetCustomerByID(Id int) (*models.Customer, error) {
	cus := models.Customer{}
	resultCus := c.db.Where("uid = ?", Id).Find(&cus)
	if resultCus.Error != nil {
		return nil, resultCus.Error
	}
	return &cus, nil
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

func NewCustomerRepository() CustomerRepository {
	db, err := NewDatabaseConnection()
	if err != nil {
		return nil
	}

	return custimerDB{db}
}
