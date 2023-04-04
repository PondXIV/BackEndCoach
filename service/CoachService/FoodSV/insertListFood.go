package foodsv

import (
	"backEndGo/models"
	"backEndGo/repository"
)

type InsertListFoodDataService interface {
	SeviceInsertListFoodByID(food *models.ListFood) int64
}
type InsertListFoodData struct {
}

// InsertListFoodByID implements InsertListFoodDataService
func (InsertListFoodData) SeviceInsertListFoodByID(food *models.ListFood) int64 {
	repoLisFood := repository.NewListFoodRepository()

	getAllFood, err := repoLisFood.GetListFoodByIDCoach(food.Cid)

	if err != nil {
		panic(err)
	}
	for _, f := range *getAllFood {
		if f.Name == food.Name {
			return 0
		}
	}
	RowsAffected := repoLisFood.InsertListFood(food)
	if RowsAffected > 0 {
		return 1
	} else if RowsAffected == 0 {
		return 0
	} else {
		return -1
	}
}

func NewInsertListFoodDataService() InsertListFoodDataService {
	return InsertListFoodData{}
}
