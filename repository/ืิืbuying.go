package repository

import (
	"backEndGo/models"

	"gorm.io/gorm"
)

type BuyingRepository interface {
	GetBuyingrAll() (*[]models.Buying, error)
}
type buyingDB struct {
	db *gorm.DB
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

func NewBuyingRepository(gormdb *gorm.DB) BuyingRepository {
	return buyingDB{db: gormdb}
}
