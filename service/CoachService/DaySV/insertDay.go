package daysv

import (
	"backEndGo/models"
	"backEndGo/repository"
)

type InsertDayDataService interface {
	ServiceInsertDay(courseID int, Day *models.DayOfCouse) (int64, error)
}
type DayDataInsert struct {
}

// ServiceInsertDay implements InsertDayDataService.
func (DayDataInsert) ServiceInsertDay(courseID int, Day *models.DayOfCouse) (int64, error) {
	repoDayOfCourse := repository.NewDayOfCourseRepository()
	rowsAffecteds, err := repoDayOfCourse.InsertDay(courseID, Day)
	if err != nil {
		panic(err)
	}
	return rowsAffecteds, nil
}

func NewInsertDayDataService() InsertDayDataService {
	return DayDataInsert{}
}
