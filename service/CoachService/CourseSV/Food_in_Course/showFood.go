package foodincourse

import (
	"backEndGo/models"
	"backEndGo/repository"
)

type ShowFoodDataService interface {
	SeviceGetFoodByDid(Did int) (*[]models.Food, error)
	SeviceGetFoodByIDCourse(CoID int) (*[]models.Food, error)
}
type FoodData struct {
}

// GetFoodByIDCourse implements ShowFoodDataService
func (FoodData) SeviceGetFoodByIDCourse(CoID int) (*[]models.Food, error) {
	repo := repository.NewFoodRepository()
	foods, err := repo.GetFoodByIDCourse(CoID)

	if err != nil {
		panic(err)
	}
	return foods, nil
}

// SeviceGetFoodByDid implements ShowFoodDataService
func (FoodData) SeviceGetFoodByDid(Did int) (*[]models.Food, error) {
	repo := repository.NewFoodRepository()
	foods, err := repo.GetFoodByIDDid(Did)

	//food, err := repository.NewDayOfCourseRepository().DayOfCourseByDid(foods.did)
	if err != nil {
		panic(err)
	}
	return foods, nil
}

func NewFoodDataService() ShowFoodDataService {
	return FoodData{}
}
