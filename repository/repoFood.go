package repository

import (
	"backEndGo/models"

	"gorm.io/gorm"
)

type FoodRepository interface {
	GetFoodByIDDid(Did int) (*[]models.Food, error)
	GetFoodByIDCourse(CoID int) (*[]models.Food, error)
}
type FoodDB struct {
	db *gorm.DB
}

// GetFoodByIDCourse implements FoodRepository
func (f FoodDB) GetFoodByIDCourse(CoID int) (*[]models.Food, error) {
	foods := []models.Food{}
	//days := []models.DayOfCouse{}
	// result := f.db.Preload("DayOfCouse").Joins("JOIN  listFood ON  listFood.ifid = Food.ifid ").Preload("ListFood").Joins("JOIN  DayOfCouse ON  DayOfCouse.did = Food.did  JOIN Course ON  DayOfCouse.coID = Course.coID  AND  Course.coID = ?", CoID).Find(&foods)
	result := f.db.Preload("DayOfCouse").Preload("ListFood").Find(&foods)
	if result.Error != nil {
		return nil, result.Error
	}

	return &foods, nil
}

// GetListClipByIDDid implements FoodRepository
func (f FoodDB) GetFoodByIDDid(Did int) (*[]models.Food, error) {
	foods := []models.Food{}
	//days := []models.DayOfCouse{}
	result := f.db.Preload("DayOfCouse").Preload("ListFood").Where("did = ?", Did).Find(&foods)

	if result.Error != nil {
		return nil, result.Error
	}

	return &foods, nil
}

func NewFoodRepository() FoodRepository {
	db, err := NewDatabaseConnection()
	if err != nil {
		return nil
	}

	return FoodDB{db}
}
