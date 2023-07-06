package coursesv

import (

	//"time"

	"backEndGo/models"
	"backEndGo/repository"
)

type UpdateCourseDataService interface {
	ServiceUpdateStatusCourse(Id int, Status string) int64
	ServiceUpdateCourse(CoID int, course *models.Course) (int64, error)
	ServiceUpdateExpiration(CoID int, Days int) (int64, error)
	ServiceUpdateDay(CoID int, Days int) (int64, error)
}

var repo = repository.NewCourseRepository()

type CourseDataUpdate struct {
}

// ServiceUpdateDay implements UpdateCourseDataService.
func (CourseDataUpdate) ServiceUpdateDay(CoID int, Days int) (int64, error) {
	RowsAffected, err := repo.UpdateDay(CoID, Days)
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

// ServiceUpdateExpiration implements UpdateCourseDataService.
func (c CourseDataUpdate) ServiceUpdateExpiration(CoID int, Days int) (int64, error) {
	RowsAffected, err := repo.UpdateExpiration(CoID, Days)
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

// ServiceUpdateCourse implements UpdateCourseDataService
func (CourseDataUpdate) ServiceUpdateCourse(CoID int, course *models.Course) (int64, error) {

	RowsAffected, err := repo.UpdateCourse(CoID, course)
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

// ServiceUpdateStatusCourse implements UpdateCourseDataService
func (CourseDataUpdate) ServiceUpdateStatusCourse(Id int, Status string) int64 {
	repo := repository.NewCourseRepository()
	RowsAffected := repo.UpdateStatusCourse(Id, Status)
	if RowsAffected > 0 {
		return 1
	} else if RowsAffected == 0 {
		return 0
	} else {
		return -1
	}
}

func NewUpdateCourseDataService() UpdateCourseDataService {
	return CourseDataUpdate{}
}
