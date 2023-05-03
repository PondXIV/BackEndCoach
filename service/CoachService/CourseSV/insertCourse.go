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
	//repoDayOfCourse := repository.NewDayOfCourseRepository()
	RowsAffected, err := repoCourse.InsertCourse(Cid, course)
	if err != nil {
		return -1, err
	}
	//rowsAffecteds := repoDayOfCourse.InsertDayOfCourse(course.CoID, course.Days)

	if RowsAffected > 0 {
		return 1, nil
	} else if RowsAffected == 0 {
		return 0, nil
	} else {
		return -1, nil
	}
}

func NewInsertCourseDataService() InsertCourseDataService {
	return CourseDataInsert{}
}
