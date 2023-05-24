package buycourse

import (
	"backEndGo/models"
	"backEndGo/repository"
)

type BuyCourseDataService interface {
	ServiceBuyCourse(CoID int, Buying *models.Buying) (int64, error)
}
type BuyCourseData struct {
}

// ServiceBuyCourse implements BuyCourseDataService
func (BuyCourseData) ServiceBuyCourse(CoID int, Buying *models.Buying) (int64, error) {

	repo := repository.NewBuyingRepository()

	bid, err := repo.BuyCourse(Buying)
	if err != nil {
		return -1, err
	}

	repoCourse := repository.NewCourseRepository()

	Price, err := repoCourse.InsertCourseByID(CoID, int(bid))

	if err != nil {
		return -1, err
	}
	if Price > 0 {
		return 1, nil
	} else if Price == 0 {
		return 0, nil
	} else {
		return -1, nil
	}
}

func NewBuyCourseDataService() BuyCourseDataService {
	return BuyCourseData{}
}
