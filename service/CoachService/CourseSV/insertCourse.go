package coursesv

import (
	"backEndGo/models"
	"backEndGo/repository"
)

type InsertCourseDataService interface {
	ServiceInsertCourse(course *models.Course) int64
}
type CourseDataInsert struct {
}

// ServiceInsertCourse implements InsertCourseDataService
func (CourseDataInsert) ServiceInsertCourse(course *models.Course) int64 {
	repo := repository.NewCourseRepository()
	RowsAffected := repo.InsertCourse(course)
	if RowsAffected > 0 {
		return 1
	} else if RowsAffected == 0 {
		return 0
	} else {
		return -1
	}
}

func NewInsertCourseDataService() InsertCourseDataService {
	return CourseDataInsert{}
}
