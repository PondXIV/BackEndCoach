package repository

import (
	"backEndGo/models"
	"fmt"

	"gorm.io/gorm"
)

type ListFoodRepository interface {
	GetListFoodAll() (*[]models.ListFood, error)
	GetListFoodAllByIDCoach(Cid int) (*[]models.ListFood, error)
	GetListFoodByIDCoach(Cid int) (*[]models.ListFood, error)
	GetListFoodByID(Ifid int) (*models.ListFood, error)
	InsertListFood(food *models.ListFood) (int64, error)
	UpdateListFood(food *models.ListFood) (int64, error)
}
type LisFoodDB struct {
	db *gorm.DB
}

// GetListFoodByID implements ListFoodRepository
func (l LisFoodDB) GetListFoodByID(Ifid int) (*models.ListFood, error) {
	listFood := models.ListFood{}
	result := l.db.Where("ifid = ?", Ifid).Find(&listFood)
	if result.Error != nil {
		return nil, result.Error
	}
	return &listFood, nil
}

// UpdateListFood implements ListFoodRepository
func (l LisFoodDB) UpdateListFood(food *models.ListFood) (int64, error) {
	result := l.db.Model(models.ListFood{}).Where("ifid = ?", food.Ifid).Updates(
		models.ListFood{Name: food.Name, Image: food.Image, Details: food.Details, Calories: food.Calories})

	if result.Error != nil {
		return -1, result.Error
	}
	if result.RowsAffected > 0 {
		fmt.Println("Update completed")
	}
	return result.RowsAffected, nil
}

// GetListFoodAllByIDCoach implements ListFoodRepository
func (l LisFoodDB) GetListFoodAllByIDCoach(Cid int) (*[]models.ListFood, error) {
	foods := []models.ListFood{}
	result := l.db.Preload("listFood").Where("cid = ?", Cid).Find(&foods)

	if result.Error != nil {
		return nil, result.Error
	}
	return &foods, nil
}

// InsertListFood implements ListFoodRepository
func (l LisFoodDB) InsertListFood(food *models.ListFood) (int64, error) {
	food.Ifid = 0
	result := l.db.Create(&food)
	if result.Error != nil {
		return -1, result.Error
	}
	if result.Error != nil {
		panic(result.Error)
	}

	return result.RowsAffected, nil
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
