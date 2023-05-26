package repository

import (
	"backEndGo/models"

	"gorm.io/gorm"
)

type BuyingRepository interface {
	GetBuyingrAll() (*[]models.Buying, error)
	BuyCourse(Buying *models.Buying) (int, error)
}
type buyingDB struct {
	db *gorm.DB
}

// BuyCourse implements BuyingRepository
func (b buyingDB) BuyCourse(Buying *models.Buying) (int, error) {

	result := b.db.Create(&models.Buying{
		Bid:         0,
		CustomerID:  Buying.CustomerID,
		BuyDateTime: Buying.BuyDateTime,
		Image:       Buying.Image,
		//Customer:    models.Customer{},
	})
	if result.Error != nil {
		return -1, result.Error
	}
	buyingLast := models.Buying{}
	getBuyLast := b.db.Order("bid DESC").Find(&buyingLast)
	if getBuyLast.Error != nil {
		panic(getBuyLast.Error)
	}
	return int(buyingLast.Bid), nil
}

// GetBuyingrAll implements BuyingRepository
func (b buyingDB) GetBuyingrAll() (*[]models.Buying, error) {
	buying := []models.Buying{}
	result := b.db.Preload("Customer").Find(&buying)
	if result.Error != nil {
		return nil, result.Error
	}
	return &buying, nil
}

func NewBuyingRepository() BuyingRepository {
	db, err := NewDatabaseConnection()
	if err != nil {
		return nil
	}

	return buyingDB{db}
}
