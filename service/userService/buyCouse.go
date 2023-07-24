package userservice

import (
	"backEndGo/models"
	"backEndGo/repository"
	"fmt"
	"math"
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
	repoCus := repository.NewCustomerRepository()
	user := repoUser.GetUserID(int(BuyCourse.CustomerID))
	sum := 0.0
	Pricecourse, _ := repoCourse.GetCourse(CoID, 0, "")
	for _, value := range *Pricecourse {
		sum = float64(user.Price) - float64(value.Price)
	}
	fmt.Println("DayID ", sum)
	res := math.Signbit(sum)

	if res != true {
		r, _ := repoCus.EditPrice(int(BuyCourse.CustomerID), sum)
		fmt.Println("DayID ", r)
		bid, err := repoBuycourse.BuyCourse(CoID, BuyCourse)
		//fid, err := repoFood.InsertFood(foods,&foodMD)

		if err != nil {
			return -1, err
		}

		_, courseID, Days, err := repoCourse.InsertCourseByID(CoID, int(bid))
		err = repoBuycourse.UpdateCoIDBuying(int(bid), courseID)

		dayID, rowsAffected := repoDay.BuyInsertDayOfCourse(uint(courseID), Days)

		didNew := make([]int, 0, len(*dayID))
		for _, value := range *dayID {
			didNew = append(didNew, int(value.Did))
		}
		fmt.Println("DayID ", didNew)

		if rowsAffected != 0 {
			//fmt.Print("dayID = \t", dayID, "\n")
			fmt.Println("SUM", sum)
		}
		/////Insert Food ต้องใช้ coIDเดิมเพื่อที่ food จะได้มี่ค่าในmodel
		getOriginalDid, _ := repoDay.DayOfCourse(0, CoID, 0)
		//getNewDid, _ := repoDay.DayOfCourse(0, courseID, 0)

		var getCilp *[]models.Clip
		var getFood *[]models.Food

		for i, value := range *getOriginalDid {
			getFood, _ = repoFood.GetFood(0, 0, int(value.Did), "")
			getCilp, _ = repoClip.GetClip(0, 0, int(value.Did))
			fmt.Println("MyGetFood", dayID)
			for _, valuefood := range *getFood {
				food, _ := repoFood.InsertBuyFood(didNew[i], valuefood.ListFoodID, valuefood.Time)
				fmt.Print("MyFood", food)
			}
			for _, valueClip := range *getCilp {
				clip, _ := repoClip.InsertBuyClip(didNew[i], int(valueClip.ListClipID), valueClip.Status)
				fmt.Print("Myclip", clip)
			}
			fmt.Println("MyGetFood", getFood)
			fmt.Println("MyGetFood", getCilp)
		}

		return 1, nil
	} else {
		return 0, nil
	}

}

func NewBuyingCourseDataService() BuyCourseService {
	return BuyingCourseData{}
}
