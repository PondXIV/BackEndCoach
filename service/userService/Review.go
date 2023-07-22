package userservice

import (
	"backEndGo/models"
	"backEndGo/repository"
	"fmt"
)

var repoReview = repository.NewReviewRepository()

type ShowReviewService interface {
	ServiceGetReviewByCoId(CoID int) (*[]models.Review, error)
	ServiceInsertReview(Bid int, review *models.Review) (int64, error)
}
type ReviewData struct {
}

// ServiceInsertReview implements ShowReviewService.
func (ReviewData) ServiceInsertReview(Bid int, review *models.Review) (int64, error) {
	repobill := repository.NewBuyingRepository()
	courseId, err := repobill.GetBuyingrAll(0, 0, Bid, 0)
	for _, valuecourseId := range *courseId {

		rowsAffecteds, _ := repoReview.InsertReview(int(valuecourseId.OriginalID), review)
		fmt.Printf("rowsAffecteds", rowsAffecteds)
	}
	//4

	if err != nil {
		panic(err)
	}
	return int64(Bid), nil
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
