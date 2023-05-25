package foodsv

import (
	"backEndGo/repository"
)

type DeleteFoodDataService interface {
	SeviceDeleteFood(Fid int) (int64, error)
}

type DeleteFoodData struct {
}

// SeviceGetListClipByIDCoach implements ShowListClipDataService
func (DeleteFoodData) SeviceDeleteFood(Fid int) (int64, error) {
	repo := repository.NewFoodRepository()
	RowsAffected, err := repo.DeleteFood(Fid)
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

func NewDeleteFoodDataService() DeleteFoodDataService {
	return DeleteFoodData{}
}
