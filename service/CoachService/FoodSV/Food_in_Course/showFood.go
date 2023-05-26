package foodSV

import (
	"backEndGo/models"
	"backEndGo/repository"
)

type ShowFoodDataService interface {
	SeviceGetFood(Fid int, Ifid int, Did int) (*[]models.Food, error)
}
type FoodData struct {
}

// SeviceGetFoodByDid implements ShowFoodDataService
func (FoodData) SeviceGetFood(Fid int, Ifid int, Did int) (*[]models.Food, error) {
	repo := repository.NewFoodRepository()
	foods, err := repo.GetFood(Fid, Ifid, Did)

	//food, err := repository.NewDayOfCourseRepository().DayOfCourseByDid(foods.did)
	if err != nil {
		panic(err)
	}
	return foods, nil
}

func NewFoodDataService() ShowFoodDataService {
	return FoodData{}
}
