package buycourse

import (
	"backEndGo/models"
	"backEndGo/repository"
	"fmt"
)

var repoBuying = repository.NewBuyingRepository()
var repoUser = repository.NewUserRepository()
var repoCourse = repository.NewCourseRepository()
var repoCustomer = repository.NewCustomerRepository()
var repoDay = repository.NewDayOfCourseRepository()

type BuyCourseDataService interface {
	ServiceBuyCourse(CoID int, Buying *models.Buying) (int64, error)
}
type BuyCourseData struct {
}

// ServiceBuyCourse implements BuyCourseDataService
func (BuyCourseData) ServiceBuyCourse(CoID int, Buying *models.Buying) (int64, error) {

	bid, err := repoBuying.BuyCourse(CoID, 0, Buying)
	if err != nil {
		return -1, err
	}

	user := repoUser.GetUserID(int(Buying.CustomerID), "")

	Price, courseID, Days, err := repoCourse.InsertCourseByID(CoID, int(bid))
	userPrice := 0
	for _, value := range *user {
		userPrice = int(value.Price)
	}
	sum := int64(Price) - int64(userPrice)

	rowsAffected := repoDay.InsertDayOfCourse(uint(courseID), Days)
	if rowsAffected != 0 {
		fmt.Print(sum)
	}
	// updateMoney := repoCustomer.EditPrice(CoID,CoID)
	// if updateMoney != 0 {

	// }
	return int64(courseID), nil
}

func NewBuyCourseDataService() BuyCourseDataService {
	return BuyCourseData{}
}
