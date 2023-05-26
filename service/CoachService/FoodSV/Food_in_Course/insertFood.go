package foodSV

import (
	"backEndGo/models"
	"backEndGo/repository"
)

type InsertFoodDataService interface {
	SeviceInsertFood(Did int, food *models.Food) (int64, error)
}
type InsertFoodData struct {
}

// InsertListFoodByID implements InsertListFoodDataService
func (InsertFoodData) SeviceInsertFood(Did int, food *models.Food) (int64, error) {
	repoFood := repository.NewFoodRepository()

	getAllFood, err := repoFood.GetFood(0, 0, Did)

	if err != nil {
		panic(err)
	}
	for _, f := range *getAllFood {
		if uint(f.ListFoodID) == uint(food.ListFoodID) {
			return 14, nil
		}
	}
	RowsAffected, err := repoFood.InsertFood(Did, food)
	if err != nil {
		return -1, err
	}
	if RowsAffected > 0 {
		return 1, nil
	} else if RowsAffected == 0 {
		return 0, nil
	} else {
		return -1, nil
	}
}

func NewInsertFoodDataService() InsertFoodDataService {
	return InsertFoodData{}
}
