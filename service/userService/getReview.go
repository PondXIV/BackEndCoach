package userservice

import (
	"backEndGo/models"
	"backEndGo/repository"
)

type ShowReviewService interface {
	ServiceGetReviewByCoId(CoID int) (*[]models.Review, error)
}
type ReviewData struct {
}

func (ReviewData) ServiceGetReviewByCoId(CoID int) (*[]models.Review, error) {
	repo := repository.NewReviewRepository()
	review, err := repo.GetReviewByIDCourse(CoID)
	if err != nil {
		panic(err)
	}
	return review, nil
}
func NewReviewDataService() ShowReviewService {
	return ReviewData{}
}
