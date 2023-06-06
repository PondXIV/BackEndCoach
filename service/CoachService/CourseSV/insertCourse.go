package coursesv

import (
	"backEndGo/models"
	"backEndGo/repository"
)

type InsertCourseDataService interface {
	ServiceInsertCourse(Cid int, course *models.Course) (int64, error)
}
type CourseDataInsert struct {
}

// ServiceInsertCourse implements InsertCourseDataService
func (CourseDataInsert) ServiceInsertCourse(Cid int, course *models.Course) (int64, error) {

	repoCourse := repository.NewCourseRepository()
	repoDayOfCourse := repository.NewDayOfCourseRepository()
	coID, err := repoCourse.InsertCourse(Cid, course)
	if err != nil {
		panic(err)
	}
	rowsAffecteds := repoDayOfCourse.InsertDayOfCourse(uint(coID), course.Days)

	return int64(rowsAffecteds), nil
}

func NewInsertCourseDataService() InsertCourseDataService {
	return CourseDataInsert{}
}
