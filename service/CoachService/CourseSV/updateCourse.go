package coursesv

import (

	//"time"

	"backEndGo/models"
	"backEndGo/repository"
)

type UpdateCourseDataService interface {
	ServiceUpdateStatusCourse(Id int, Status string) int64
	ServiceUpdateCourse(course *models.Course) int64
}
type CourseDataUpdate struct {
}

// ServiceUpdateCourse implements UpdateCourseDataService
func (CourseDataUpdate) ServiceUpdateCourse(course *models.Course) int64 {
	repo := repository.NewCourseRepository()
	RowsAffected := repo.UpdateCourse(course)
	if RowsAffected > 0 {
		return 1
	} else if RowsAffected == 0 {
		return 0
	} else {
		return -1
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
