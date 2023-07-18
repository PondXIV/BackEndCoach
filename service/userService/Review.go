package userservice

import (
	"backEndGo/models"
	"backEndGo/repository"
	"fmt"
)

var repoReview = repository.NewReviewRepository()

type ShowReviewService interface {
	ServiceGetReviewByCoId(CoID int) (*[]models.Review, error)
	ServiceInsertReview(CoID int, review *models.Review) (int64, error)
}
type ReviewData struct {
}

// ServiceInsertReview implements ShowReviewService.
func (ReviewData) ServiceInsertReview(CoID int, review *models.Review) (int64, error) {
	//1
	repo := repository.NewCourseRepository()
	courseDup, err := repo.GetCouseByCoID(CoID)
	fmt.Println(courseDup.CoID)

	//2
	billID, err := repo.GetCouseByCoID(int(courseDup.CoID))
	fmt.Println(billID.BuyingID)
	//3
	repobill := repository.NewBuyingRepository()
	courseId, err := repobill.GetBuyingrAll(0, 0, int(billID.BuyingID))
	for _, valuecourseId := range *courseId {

		rowsAffecteds, _ := repoReview.InsertReview(int(valuecourseId.CourseID), review)
		fmt.Printf("rowsAffecteds", rowsAffecteds)
	}
	//4

	if err != nil {
		panic(err)
	}
	return int64(billID.BuyingID), nil
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
