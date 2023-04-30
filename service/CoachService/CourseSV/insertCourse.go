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
	repoCourse := repository.NewCourseRepository()
	repoDayOfCourse := repository.NewDayOfCourseRepository()
	RowsAffected := repoCourse.InsertCourse(course)
	RowsAffecteds := repoDayOfCourse.InsertDayOfCourse(course.CoachID, course.Days)
	if RowsAffected > 0 || RowsAffecteds > 0 {
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
