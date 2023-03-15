package coursesv

import (
	"backEndGo/models"
	"backEndGo/repository"

	//"time"

	"gorm.io/gorm"
)

type ShowCourseDataService interface {
	ServiceGetCourseByIDCoach(Id int) (*[]models.Course, error)
	SeviceGetCourseByName(Name string) (*[]models.Course, error)
}
type CourseData struct {
	db *gorm.DB
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
func NewCourseByNameService() ShowCourseDataService {
	return CourseData{}
}
