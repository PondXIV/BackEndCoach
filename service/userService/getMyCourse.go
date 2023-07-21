package userservice

import (
	"backEndGo/models"
	"backEndGo/repository"
)

var repo = repository.NewCourseRepository()

type GetMycourseService interface {
	ServiceGetMycourse(Uid int) (*[]models.Buying, error)
	ServiceGetCourseEX(Uid int) (*[]models.Buying, error)
}
type MyCourseData struct {
}

// ServiceGetCourseEX implements GetMycourseService.
func (MyCourseData) ServiceGetCourseEX(Uid int) (*[]models.Buying, error) {
	course, err := repository.NewBuyingRepository().GetCourseByIDCusEX(Uid)
	if err != nil {
		panic(err)
	}
	return course, nil
}

// ServiceGetMycourse implements GetMycourseService
func (MyCourseData) ServiceGetMycourse(Uid int) (*[]models.Buying, error) {

	course, err := repository.NewBuyingRepository().GetCourseByIDCus(Uid)
	if err != nil {
		panic(err)
	}
	return course, nil
}

func NewMyCourseDataService() GetMycourseService {
	return MyCourseData{}
}
