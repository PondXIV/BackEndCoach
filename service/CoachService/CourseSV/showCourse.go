package coursesv

import (
	"backEndGo/models"
	"backEndGo/repository"
	//"time"
)

type ShowCourseDataService interface {
	ServiceGetCourseByIDCoach(Id int) (*[]models.Course, error)
	SeviceGetCourseByName(Name string) (*[]models.Course, error)
	SeviceGetCourseByCoID(CoID int) (*models.Course, error)
	SeviceGetCourse(CoID int, Cid int, Name string) (*[]models.Course, error)
}
type CourseData struct {
}

// SeviceGetCourseAll implements ShowCourseDataService
func (CourseData) SeviceGetCourse(CoID int, Cid int, Name string) (*[]models.Course, error) {
	repo := repository.NewCourseRepository()
	course, err := repo.GetCourse(CoID, Cid, Name)

	if err != nil {
		panic(err)
	}
	return course, nil
}

// SeviceGetCourseByCoID implements ShowCourseDataService
func (CourseData) SeviceGetCourseByCoID(CoID int) (*models.Course, error) {
	repo := repository.NewCourseRepository()
	course, err := repo.GetCouseByCoID(CoID)

	if err != nil {
		panic(err)
	}
	return course, nil
}

// SeviceGetCourseByName implements ShowCourseDataService
func (CourseData) SeviceGetCourseByName(Name string) (*[]models.Course, error) {
	repo := repository.NewCourseRepository()
	course, err := repo.GetCouseByname(Name)
	if err != nil {
		panic(err)
	}
	return course, nil
}

// ServiceGetCourseByIDCoach implements ShowCourseDataService
func (CourseData) ServiceGetCourseByIDCoach(Id int) (*[]models.Course, error) {
	_, err := repository.NewDayOfCourseRepository().DayOfCourseAll()
	// course, err := repo.GetCourseByIDCoach(Id)
	repo := repository.NewCourseRepository()
	course, err := repo.GetCourseByIDCoach(Id)

	if err != nil {
		panic(err)
	}
	return course, nil
}

func NewCourseDataService() ShowCourseDataService {
	return CourseData{}
}
