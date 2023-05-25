package repository

import (
	"backEndGo/models"
	"fmt"

	"gorm.io/gorm"
)

type FoodRepository interface {
	GetFood(Fid int, Ifid int, Did int) (*[]models.Food, error)
	InsertFood(Did int, Food *models.Food) (int64, error)
	UpdateFood(Fid int, food *models.Food) (int64, error)
	DeleteFood(Fid int) (int64, error)
}
type FoodDB struct {
	db *gorm.DB
}

// DeleteFood implements FoodRepository
func (f FoodDB) DeleteFood(Fid int) (int64, error) {
	foodID := &models.Food{
		Fid: uint(Fid),
	}
	result := f.db.Delete(foodID)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

// UpdateFood implements FoodRepository
func (f FoodDB) UpdateFood(Fid int, food *models.Food) (int64, error) {
	result := f.db.Model(models.Food{}).Where("fid = ?", Fid).Updates(
		models.Food{
			ListFoodID: food.ListFoodID,
			Time:       food.Time,
		})

	if result.Error != nil {
		return -1, result.Error
	}
	if result.RowsAffected > 0 {
		fmt.Println("Update completed")
	}
	return result.RowsAffected, nil
}

// InsertFood implements FoodRepository
func (f FoodDB) InsertFood(Did int, Food *models.Food) (int64, error) {
	result := f.db.Create(&models.Food{
		Fid:          0,
		ListFoodID:   int(Food.Fid),
		DayOfCouseID: uint(Did),
		Time:         Food.Time,
	})
	if result.Error != nil {
		return -1, result.Error
	}
	if result.Error != nil {
		panic(result.Error)
	}

	return result.RowsAffected, nil
}

// GetListClipByIDDid implements FoodRepository
func (f FoodDB) GetFood(Fid int, Ifid int, Did int) (*[]models.Food, error) {
	foods := []models.Food{}
	result := f.db.Where("ifid IS NOT NULL")
	if Fid != 0 {
		result.Where("fid=?", Fid)
	}
	if Ifid != 0 {
		result.Where("ifid=?", Ifid)
	}
	if Did != 0 {
		result.Where("did=?", Did)
	}
	if result.Error != nil {
		return nil, result.Error
	}
	result.Find(&foods)

	return &foods, nil
}

func NewFoodRepository() FoodRepository {
	db, err := NewDatabaseConnection()
	if err != nil {
		return nil
	}

	return FoodDB{db}
}
