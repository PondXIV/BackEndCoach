package foodsv

import (
	"backEndGo/models"
	"backEndGo/repository"
)

type ShowListFoodDataService interface {
	SeviceGetFoodByIDCoach(Cid int) (*[]models.ListFood, error)
}
type ListFoodData struct {
}

// GetFoodByIDCoach implements ShowListFoodDataService
func (ListFoodData) SeviceGetFoodByIDCoach(Cid int) (*[]models.ListFood, error) {
	repo := repository.NewListFoodRepository()
	course, err := repo.GetListFoodByIDCoach(Cid)
	if err != nil {
		panic(err)
	}
	return course, nil
}

func NewListFoodDataService() ShowListFoodDataService {
	return ListFoodData{}
}
