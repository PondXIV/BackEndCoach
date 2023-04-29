package foodsv

import (
	"backEndGo/models"
	"backEndGo/repository"
)

type ShowListFoodDataService interface {
	SeviceGetFoodByID(Ifid int) (*models.ListFood, error)
	SeviceGetFoodByIDCoach(Cid int) (*[]models.ListFood, error)
}
type ListFoodData struct {
}

// SeviceGetFoodByID implements ShowListFoodDataService
func (ListFoodData) SeviceGetFoodByID(Ifid int) (*models.ListFood, error) {
	repo := repository.NewListFoodRepository()
	listFood, err := repo.GetListFoodByID(Ifid)
	if err != nil {
		return nil, err
	}
	return listFood, nil
}

// GetFoodByIDCoach implements ShowListFoodDataService
func (ListFoodData) SeviceGetFoodByIDCoach(Cid int) (*[]models.ListFood, error) {
	repo := repository.NewListFoodRepository()
	listFood, err := repo.GetListFoodByIDCoach(Cid)
	if err != nil {
		return nil, err
	}
	return listFood, nil
}

func NewListFoodDataService() ShowListFoodDataService {
	return ListFoodData{}
}
