package daysv

import (
	"backEndGo/models"
	"backEndGo/repository"
)

type UpdateDayDataService interface {
	ServiceUpdateDay(Did int, Day *models.DayOfCouse) (int64, error)
}
type UpdateDayData struct {
}

// ServiceUpdateDay implements UpdateListFoodDataService
func (UpdateDayData) ServiceUpdateDay(Did int, Day *models.DayOfCouse) (int64, error) {
	repoDay := repository.NewDayOfCourseRepository()

	RowsAffected, err := repoDay.UpdateDay(Did, Day)
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

func NewUpdateDayDataService() UpdateDayDataService {
	return UpdateDayData{}
}
