package foodsv

import (
	"backEndGo/repository"
)

type DeleteListFoodDataService interface {
	SeviceDeleteListFood(Ifid int) (int64, error)
}

type DeleteListFoodData struct {
}

// SeviceGetListClipByIDCoach implements ShowListClipDataService
func (DeleteListFoodData) SeviceDeleteListFood(Ifid int) (int64, error) {
	repo := repository.NewListFoodRepository()
	RowsAffected, err := repo.DeleteListFood(Ifid)
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

func NewDeleteListFoodDataService() DeleteListFoodDataService {
	return DeleteListFoodData{}
}
