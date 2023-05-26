package repository

import (
	"backEndGo/models"
	"fmt"

	"gorm.io/gorm"
)

type ListFoodRepository interface {
	GetListFoodAll() (*[]models.ListFood, error)

	GetListFood(Ifid int, Cid int, Name string) (*[]models.ListFood, error)

	InsertListFood(Cid int, food *models.ListFood) (int64, error)
	UpdateListFood(Ifid int, food *models.ListFood) (int64, error)
	DeleteListFood(Ifid int) (int64, error)
}
type LisFoodDB struct {
	db *gorm.DB
}

// DeleteListFood implements ListFoodRepository
func (l LisFoodDB) DeleteListFood(Ifid int) (int64, error) {
	foodID := &models.ListFood{Ifid: uint(Ifid)}
	result := l.db.Delete(foodID)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

// GetFood implements ListFoodRepository
func (l LisFoodDB) GetListFood(Ifid int, Cid int, Name string) (*[]models.ListFood, error) {
	foods := []models.ListFood{}
	result := l.db.Where("ifid IS NOT NULL")
	if Ifid != 0 {
		result.Where("ifid=?", Ifid)
	}
	if Cid != 0 {
		result.Where("cid=?", Cid)
	}
	if Name != "" {
		result.Where("name like ?", "%"+Name+"%")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	result.Find(&foods)

	return &foods, nil
}

// UpdateListFood implements ListFoodRepository
func (l LisFoodDB) UpdateListFood(Ifid int, food *models.ListFood) (int64, error) {
	result := l.db.Model(models.ListFood{}).Where("ifid = ?", Ifid).Updates(
		models.ListFood{Name: food.Name, Image: food.Image, Details: food.Details, Calories: food.Calories})

	if result.Error != nil {
		return -1, result.Error
	}
	if result.RowsAffected > 0 {
		fmt.Println("Update completed")
	}
	return result.RowsAffected, nil
}

// InsertListFood implements ListFoodRepository
func (l LisFoodDB) InsertListFood(Cid int, food *models.ListFood) (int64, error) {
	food.Ifid = 0
	result := l.db.Create(&models.ListFood{
		Ifid:     food.Ifid,
		CoachID:  Cid,
		Name:     food.Name,
		Image:    food.Image,
		Details:  food.Details,
		Calories: food.Calories,
	})
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

func NewListFoodRepository() ListFoodRepository {
	db, err := NewDatabaseConnection()
	if err != nil {
		return nil
	}

	return LisFoodDB{db}
}
