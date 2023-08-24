package buycourse

import (
	"backEndGo/models"
	"backEndGo/repository"
	"fmt"
	"time"
)

type ShowBuyingDataService interface {
	GetBuying(uid int, coID int, bid int, cid int, ocoID int) (*[]models.Buying, error)
	SeviceGetCourseByUser(Cid int) (*[]models.Buying, error)
	SeviceGetCourseCount(OcoID int) int
}
type BuyingData struct {
}

// SeviceGetCourseCount implements ShowBuyingDataService.
func (BuyingData) SeviceGetCourseCount(OcoID int) int {
	sum := 0
	count := 0
	dt := time.Now()
	day := dt.AddDate(0, 0, -1)
	repobuy := repository.NewBuyingRepository()
	repocourse := repository.NewCourseRepository()
	courseOriginal, err := repobuy.GetBuyingrAll(0, 0, 0, 0, OcoID)
	for _, valuecourseId := range *courseOriginal {
		sum++
		fmt.Println("Sum++", sum)
		course, _ := repocourse.GetCourseNotNull(int(valuecourseId.CourseID))
		for _, valuecourse := range *course {

			if valuecourse.ExpirationDate.After(dt) || valuecourse.ExpirationDate.IsZero() {
				fmt.Printf("CoID1", "%v", valuecourse.CoID)
				fmt.Printf("time1", "%v", valuecourse.ExpirationDate)
				fmt.Println("count1", "%v", count)
				//count = count + 1
			} else if valuecourse.ExpirationDate.Before(day) {
				fmt.Printf("CoID3", "%v", valuecourse.CoID)
				fmt.Printf("time3", "%v", valuecourse.ExpirationDate)
				//fmt.Println("count3", count)
				fmt.Println("day", day)
				count = count - 1
			}
			fmt.Println("count", count)
		}
		sum = sum + count
		fmt.Println("count", count)
		fmt.Println("Sum", sum)
	}
	if err != nil {
		panic(err)
	}
	return sum
}

// SeviceGetCourseByUser implements ShowCourseDataService.
func (b BuyingData) SeviceGetCourseByUser(Cid int) (*[]models.Buying, error) {
	repo := repository.NewBuyingRepository()
	course, err := repo.GetBuyingrAll(0, 0, 0, Cid, 0)
	keys := make(map[int]bool)
	buying := []models.Buying{}
	for _, entry := range *course {
		if _, value := keys[int(entry.CustomerID)]; !value {
			keys[int(entry.CustomerID)] = true
			buying = append(buying, entry)
		}
	}
	if err != nil {
		panic(err)
	}
	return &buying, nil
}

// SeviceGetListClipByIDCoach implements ShowListClipDataService
func (c BuyingData) GetBuying(uid int, coID int, bid int, cid int, ocoID int) (*[]models.Buying, error) {
	repo := repository.NewBuyingRepository()
	Buys, err := repo.GetBuyingrAll(uid, coID, bid, cid, ocoID)
	if err != nil {
		return nil, err
	}
	return Buys, nil
}

func NewBuyingDataService() ShowBuyingDataService {
	return BuyingData{}
}
