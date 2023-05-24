package repository

import (
	"backEndGo/models"
	"fmt"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	GetCustomerAll() (*[]models.Customer, error)
	GetCustomerByID(Id int) (*models.Customer, error)
	UserByUid(Uid int) (*models.Customer, error)
	UpdateUser(Uid int, customer *models.Customer) (int64, error)
	EditPrice(Uid int, Price int) (int64, error)
}
type custumerDB struct {
	db *gorm.DB
}

// editMoney implements CustomerRepository
func (c custumerDB) EditPrice(Uid int, Price int) (int64, error) {
	result := c.db.Model(models.Customer{}).Where("uid = ?", Uid).Updates(
		models.Customer{Price: uint(Price)})
	if result.Error != nil {
		return -1, result.Error
	}

	return result.RowsAffected, nil
}

// UpdateUser implements CustomerRepository
func (c custumerDB) UpdateUser(Uid int, customer *models.Customer) (int64, error) {
	result := c.db.Model(models.Customer{}).Where("uid = ?", Uid).Updates(
		models.Customer{Username: customer.Username, Password: customer.Username, FullName: customer.FullName, Birthday: customer.Birthday,
			Gender: customer.Gender, Phone: customer.Phone, Image: customer.Image, Weight: customer.Weight, Height: customer.Height})

	if result.Error != nil {
		return -1, result.Error
	}
	if result.RowsAffected > 0 {
		fmt.Println("Update completed")
	}
	return result.RowsAffected, nil
}

// UserByUid implements CustomerRepository
func (c custumerDB) UserByUid(Uid int) (*models.Customer, error) {
	users := models.Customer{}
	result := c.db.Find(&users)
	if Uid != 0 {
		result.Where("uid = ?", Uid).Find(&users)
	}
	result.Find(&users)
	return &users, nil
}

// GetCustomerByID implements CustomerRepository
func (c custumerDB) GetCustomerByID(Id int) (*models.Customer, error) {
	cus := models.Customer{}
	resultCus := c.db.Where("uid = ?", Id).Find(&cus)
	if resultCus.Error != nil {
		return nil, resultCus.Error
	}
	return &cus, nil
}

// GetAll implements CustomerRepository
func (c custumerDB) GetCustomerAll() (*[]models.Customer, error) {
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

	return custumerDB{db}
}
