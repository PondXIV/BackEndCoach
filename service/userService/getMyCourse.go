package userservice

import (
	"backEndGo/models"
	"backEndGo/repository"
)

type GetMycourseService interface {
	ServiceGetMycourse(Uid int, Bid int) (*[]models.Course, error)
}
type MyCourseData struct {
}

// ServiceGetMycourse implements GetMycourseService
func (MyCourseData) ServiceGetMycourse(Uid int, Bid int) (*[]models.Course, error) {
	repo := repository.NewCourseRepository()
	course, err := repo.GetCourseByIDCus(Uid, Bid)
	if err != nil {
		panic(err)
	}
	return course, nil
}

func NewMyCourseDataService() GetMycourseService {
	return MyCourseData{}
}
