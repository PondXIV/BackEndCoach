package foodsv

import (
	"backEndGo/models"
	"backEndGo/repository"
)

type UpdateListFoodDataService interface {
	ServiceUpdateListFood(food *models.ListFood) (int64, error)
}
type UpdateListFoodData struct {
}

// ServiceUpdateListFood implements UpdateListFoodDataService
func (UpdateListFoodData) ServiceUpdateListFood(food *models.ListFood) (int64, error) {
	repoLisFood := repository.NewListFoodRepository()
	getAllFood, err := repoLisFood.GetListFoodByIDCoach(food.CoachID)

	if err != nil {
		panic(err)
	}
	for _, f := range *getAllFood {
		if f.Name == food.Name {
			return 0, nil
		}
	}

	RowsAffected, err := repoLisFood.UpdateListFood(food)
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

func NewUpdateListFoodDataService() UpdateListFoodDataService {
	return UpdateListFoodData{}
}