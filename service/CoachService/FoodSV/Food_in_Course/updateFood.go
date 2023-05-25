package foodsv

import (
	"backEndGo/models"
	"backEndGo/repository"
)

type UpdateFoodDataService interface {
	ServiceUpdateFood(Fid int, food *models.Food) (int64, error)
}
type UpdateFoodData struct {
}

// ServiceUpdateListFood implements UpdateListFoodDataService
func (UpdateFoodData) ServiceUpdateFood(Fid int, food *models.Food) (int64, error) {
	repoFood := repository.NewFoodRepository()
	getAllFood, err := repoFood.GetFood(Fid, 0, 0)

	if err != nil {
		panic(err)
	}
	for _, f := range *getAllFood {
		if uint(f.ListFoodID) == uint(food.ListFoodID) {
			if f.fid != uint(Fid) {
				return 3, nil
			}

		}
	}

	RowsAffected, err := repoLisFood.UpdateFood(Fid, food)
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

func NewUpdateFoodDataService() UpdateFoodDataService {
	return UpdateFoodData{}
}
