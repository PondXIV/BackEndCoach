package userservice

import (
	"backEndGo/models"
	"backEndGo/repository"
)

var repoReview = repository.NewReviewRepository()

type ShowReviewService interface {
	ServiceGetReviewByCoId(CoID int) (*[]models.Review, error)
	ServiceInsertReview(CusID int, review *models.Review) (int64, error)
}
type ReviewData struct {
}

// ServiceInsertReview implements ShowReviewService.
func (ReviewData) ServiceInsertReview(CusID int, review *models.Review) (int64, error) {
	rowsAffecteds, err := repoReview.InsertReview(CusID, review)
	if err != nil {
		panic(err)
	}
	return int64(rowsAffecteds), nil
}

func (ReviewData) ServiceGetReviewByCoId(CoID int) (*[]models.Review, error) {

	review, err := repoReview.GetReviewByIDCourse(CoID)
	if err != nil {
		panic(err)
	}
	return review, nil
}
func NewReviewDataService() ShowReviewService {
	return ReviewData{}
}
