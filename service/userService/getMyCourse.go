package userservice

import (
	"backEndGo/models"
	"backEndGo/repository"
)

type GetMycourseService interface {
	ServiceGetMycourse(Uid int, CoId int, Name string) (*[]models.Course, error)
}
type MyCourseData struct {
}

// ServiceGetMycourse implements GetMycourseService
func (MyCourseData) ServiceGetMycourse(Uid int, CoId int, Name string) (*[]models.Course, error) {
	repo := repository.NewCourseRepository()

	course, err := repo.GetCourse(Uid, CoId, Name)
	if err != nil {
		panic(err)
	}
	return course, nil

}

func NewMyCourseDataService() GetMycourseService {
	return MyCourseData{}
}
