package foodSV

import (
	"backEndGo/models"
	"backEndGo/repository"
)

type UpdateFoodDataService interface {
	ServiceUpdateFood(Fid int, food *models.Food) (int64, error)
}
type UpdateFoodData struct {
}

// ServiceUpdateFood implements UpdateFoodDataService
func (UpdateFoodData) ServiceUpdateFood(Fid int, food *models.Food) (int64, error) {
	repoLisFood := repository.NewFoodRepository()
	getAllFood, err := repoLisFood.GetFood(0, food.ListFoodID, 0)

	if err != nil {
		panic(err)
	}
	for _, f := range *getAllFood {
		if f.ListFoodID == food.ListFoodID {
			if f.Time == food.Time {
				if f.Fid != uint(Fid) {
					return 14, nil
				}
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
