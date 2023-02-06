package customer

import (
	"backEndGo/models"
	"backEndGo/repository"

	//"time"

	"gorm.io/gorm"
)

type ShowCourseDataService interface {
	ServiceGetCourseByIDCoach(Id int) (*models.Course, error)
}
type CourseData struct {
	db *gorm.DB
}

// ServiceGetCourseByIDCoach implements ShowCourseDataService
func (CourseData) ServiceGetCourseByIDCoach(Id int) (*models.Course, error) {
	repo := repository.NewUserRepository()
	coach, cus, err := repo.LoginFB(FackbookID)
	if err != nil {
		panic(err)
	}
	return coach, cus, nil
}

func NewCourseDataService() ShowCourseDataService {
	return CourseData{}
}
