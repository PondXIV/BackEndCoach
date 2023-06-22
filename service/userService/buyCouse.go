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
	repoClip := repository.NewClipRepository()
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
	/////Insert Food ต้องใช้ coIDเดิมเพื่อที่ food จะได้มี่ค่าในmodel
	getDid, _ := repoDay.DayOfCourse(0, CoID, 0)
	for _, value := range *getDid {
		getFood, _ := repoFood.GetFood(0, 0, int(value.Did))
		getCilp, _ := repoClip.GetClip(0, 0, int(value.Did))

		fmt.Println("MyGetFood", getFood)
		for _, valuefood := range *getFood {
			food, _ := repoFood.InsertBuyFood(int(valuefood.DayOfCouseID), valuefood.ListFoodID, valuefood.Time)
			fmt.Printf("MyFood", food)
		}
		for _, valueClip := range *getCilp {
			clip, _ := repoClip.InsertBuyClip(int(valueClip.DayOfCouseID), int(valueClip.ListClipID), valueClip.Status)
			fmt.Printf("Myclip", clip)
		}
	}
	/////Insert Clip ต้องใช้ coIDเดิมเพื่อที่ clip จะได้มี่ค่าในmodel

	//rowsAffectedFood, err := repoFood.InsertBuyFood(dayID, getFood)

	// days := models.DayOfCouse{}
	// getday := repoDay.InsertDayOfCourse(uint(CoID))
	// fmt.Println("520", days)
	// food := models.Food{}
	// fid, err := repoFood.InsertFood(int(days.Did), &food)
	// rowsAffectedFood := &fid
	// if rowsAffectedFood > 0 {

	// 	panic("Success")

	// }
	//row := repoFood.InsertFood(int(days.Did),&food)
	//fmt.Printf("fid", fid)

	return int64(courseID), nil
}

func NewBuyingCourseDataService() BuyCourseService {
	return BuyingCourseData{}
}
