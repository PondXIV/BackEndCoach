package daysv

import (
	"backEndGo/repository"
)

type DeleteDayDataService interface {
	SeviceDeleteDay(Did int) (int64, error)
}

type DeleteDayData struct {
}

// SeviceDeleteDay implements DeleteDayDataService.
func (DeleteDayData) SeviceDeleteDay(Did int) (int64, error) {
	repo := repository.NewDayOfCourseRepository()
	RowsAffected, err := repo.DeleteDay(Did)
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

func NewDeleteDayDataService() DeleteDayDataService {
	return DeleteDayData{}
}
