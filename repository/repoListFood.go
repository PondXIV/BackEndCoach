package repository

import (
	"backEndGo/models"

	"gorm.io/gorm"
)

type ListFoodRepository interface {
	GetListFoodAll() (*[]models.ListFood, error)
	GetListFoodByIDCoach(Cid int) (*[]models.ListFood, error)
	InsertListFood(food *models.ListFood) int64
}
type LisFoodDB struct {
	db *gorm.DB
}

// InsertListFood implements ListFoodRepository
func (l LisFoodDB) InsertListFood(food *models.ListFood) int64 {
	food.Ifid = 0
	result := l.db.Create(&food)
	if result.Error != nil {
		panic(result.Error)
	}

	return result.RowsAffected
}

// GetFoodAll implements ListFoodRepository
func (l LisFoodDB) GetListFoodAll() (*[]models.ListFood, error) {
	foods := []models.ListFood{}
	result := l.db.Preload("listFood").Find(&foods)

	if result.Error != nil {
		return nil, result.Error
	}
	return &foods, nil
}

// GetFood implements ListFoodRepository
func (l LisFoodDB) GetListFoodByIDCoach(Cid int) (*[]models.ListFood, error) {
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
