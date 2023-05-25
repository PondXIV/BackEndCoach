package foodsv

import (
	"backEndGo/models"
	"backEndGo/repository"
)

type ShowListFoodDataService interface {
	SeviceGetFood(Ifid int, Cid int, Name string) (*[]models.ListFood, error)
}
type ListFoodData struct {
}

// GetFoodByIDCoach implements ShowListFoodDataService
func (ListFoodData) SeviceGetFood(Ifid int, Cid int, Name string) (*[]models.ListFood, error) {
	repo := repository.NewListFoodRepository()
	listFood, err := repo.GetListFood(Ifid, Cid, Name)
	if err != nil {
		return nil, err
	}
	return listFood, nil
}

func NewListFoodDataService() ShowListFoodDataService {
	return ListFoodData{}
}
