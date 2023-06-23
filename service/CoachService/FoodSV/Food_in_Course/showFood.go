package foodSV

import (
	"backEndGo/models"
	"backEndGo/repository"
)

type ShowFoodDataService interface {
	SeviceGetFood(Fid int, Ifid int, Did int, Name string) (*[]models.Food, error)
}
type FoodData struct {
}

// SeviceGetFoodByDid implements ShowFoodDataService
func (FoodData) SeviceGetFood(Fid int, Ifid int, Did int, Name string) (*[]models.Food, error) {
	repo := repository.NewFoodRepository()
	foods, err := repo.GetFood(Fid, Ifid, Did, Name)

	//food, err := repository.NewDayOfCourseRepository().DayOfCourseByDid(foods.did)
	if err != nil {
		panic(err)
	}
	return foods, nil
}

func NewFoodDataService() ShowFoodDataService {
	return FoodData{}
}
