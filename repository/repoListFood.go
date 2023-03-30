package repository

import (
	"backEndGo/models"

	"gorm.io/gorm"
)

type ListFoodRepository interface {
	GetFoodByIDCoach(Cid int) (*[]models.ListFood, error)
}
type LisFoodDB struct {
	db *gorm.DB
}

// GetFood implements ListFoodRepository
func (l LisFoodDB) GetFoodByIDCoach(Cid int) (*[]models.ListFood, error) {
	foods := []models.ListFood{}
	result := l.db.Where("cid = ?", Cid).Find(&foods)
	if result.Error != nil {
		return nil, result.Error
	}
	return &foods, nil
}

func NewListFoodRepository() ListFoodRepository {
	db, err := NewDatabaseConnection()
	if err != nil {
		return nil
	}

	return LisFoodDB{db}
}
