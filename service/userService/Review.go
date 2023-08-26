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
	ServiceGetReviewByCid(Cid int) (float64, error)
}
type ReviewData struct {
}

// ServiceGetReviewByCid implements ShowReviewService.
func (ReviewData) ServiceGetReviewByCid(Cid int) (float64, error) {
	repo := repository.NewCourseRepository()
	courseId, err := repo.GetCourseSell(0, Cid, "")
	sum := 0.0
	avg := 0.0
	count := 0
	fmt.Println("333 =", Cid)
	for _, valuecourseId := range *courseId {
		fmt.Println("1111 =", valuecourseId.CoID)

		review, _ := repoReview.GetReviewByIDCourse(int(valuecourseId.CoID))
		for _, valuereview := range *review {

			count++
			sum = sum + float64(valuereview.Score)

		}
	}
	if sum > 0 {
		avg = sum / float64(count)
		fmt.Println("count =", count)
		fmt.Println("sum =", sum)
		fmt.Println("avg =", avg)
	} else {
		avg = 0.0
		fmt.Println("count =", count)
		fmt.Println("sum =", sum)
		fmt.Println("avg =", avg)
	}

	if err != nil {
		panic(err)
	}

	return avg, nil
}

// ServiceInsertReview implements ShowReviewService.
func (ReviewData) ServiceInsertReview(Bid int, review *models.Review) (int64, error) {
	repobill := repository.NewBuyingRepository()
	courseId, err := repobill.GetBuyingrAll(0, 0, Bid, 0, 0)
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
