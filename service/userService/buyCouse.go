package userservice

import (
	"backEndGo/models"
	"backEndGo/repository"
	"fmt"
)

type BuyCourseService interface {
	ServiceBuyCourse(CoID int, BuyCourse *models.Buying) (int64, error)
}
type BuyingCourseData struct {
}

// ServiceBuyCourse implements BuyCourseService.
func (BuyingCourseData) ServiceBuyCourse(CoID int, BuyCourse *models.Buying) (int64, error) {
	repoBuycourse := repository.NewBuyingRepository()
	repoUser := repository.NewUserRepository()
	repoCourse := repository.NewCourseRepository()
	repoDay := repository.NewDayOfCourseRepository()
	// repoClip := repository.NewClipRepository()
	repoFood := repository.NewFoodRepository()
	bid, err := repoBuycourse.BuyCourse(BuyCourse)
	//fid, err := repoFood.InsertFood(foods,&foodMD)
	if err != nil {
		return -1, err
	}
	user := repoUser.GetUserID(int(BuyCourse.CustomerID))

	Price, courseID, Days, err := repoCourse.InsertCourseByID(CoID, int(bid))
	sum := int64(Price) - int64(user.Price)

	dayID, rowsAffected := repoDay.BuyInsertDayOfCourse(uint(courseID), Days)

	if rowsAffected != 0 {
		fmt.Printf("dayID = %d \t", dayID, "\n")
		fmt.Println("SUM", sum)
	}
	//dayID,err := repoDay.DayOfCourseByCoid(courseID)
	//did  := dayID

	rowsAffectedFood, err := repoFood.InsertBuyFood(dayID, &models.Food{})
	if rowsAffectedFood != 0 {

		fmt.Println(rowsAffectedFood)

	}

	// days := models.DayOfCouse{}
	// getday := repoDay.InsertDayOfCourse(uint(CoID))
	// fmt.Println("520", days)
	// food := models.Food{}
	// fid, err := repoFood.InsertFood(int(days.Did), &food)
	// rowsAffectedFood := &fid
	// if rowsAffectedFood != nil {

	// 	fmt.Println(fid)

	// }
	//row := repoFood.InsertFood(int(days.Did),&food)
	//fmt.Printf("fid", fid)

	return int64(courseID), nil
}

func NewBuyingCourseDataService() BuyCourseService {
	return BuyingCourseData{}
}
