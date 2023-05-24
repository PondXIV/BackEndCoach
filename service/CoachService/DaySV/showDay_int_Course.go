package daysv

import (
	"backEndGo/models"
	"backEndGo/repository"
)

type ShowDayDataService interface {
	SeviceGetDay(Did int, CoID int, Sequence int) (*[]models.DayOfCouse, error)
}
type DayData struct {
}

// GetFoodByIDCoach implements ShowListFoodDataService
func (DayData) SeviceGetDay(Did int, CoID int, Sequence int) (*[]models.DayOfCouse, error) {
	repo := repository.NewDayOfCourseRepository()
	dyas, err := repo.DayOfCourse(Did, CoID, Sequence)
	if err != nil {
		return nil, err
	}
	return dyas, nil
}

func NewDayDataService() ShowDayDataService {
	return DayData{}
}
