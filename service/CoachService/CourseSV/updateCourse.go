package coursesv

import (

	//"time"

	"backEndGo/repository"
)

type UpdateCourseDataService interface {
	ServiceUpdateStatusCourse(Id int, Status string) int64
}
type CourseDataUpdate struct {
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
