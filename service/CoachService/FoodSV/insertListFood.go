package foodsv

import (
	"backEndGo/models"
	"backEndGo/repository"
)

type InsertListFoodDataService interface {
	SeviceInsertListFoodByID(Cid int, food *models.ListFood) (int64, error)
}
type InsertListFoodData struct {
}

// InsertListFoodByID implements InsertListFoodDataService
func (InsertListFoodData) SeviceInsertListFoodByID(Cid int, food *models.ListFood) (int64, error) {
	repoLisFood := repository.NewListFoodRepository()

	getAllFood, err := repoLisFood.GetListFood(0, Cid, "")

	if err != nil {
		panic(err)
	}
	for _, f := range *getAllFood {
		if f.Name == food.Name {
			return 0, nil
		}
	}
	RowsAffected, err := repoLisFood.InsertListFood(Cid, food)
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

func NewInsertListFoodDataService() InsertListFoodDataService {
	return InsertListFoodData{}
}
